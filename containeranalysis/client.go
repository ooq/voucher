package containeranalysis

import (
	"context"
	"errors"

	containeranalysisapi "cloud.google.com/go/containeranalysis/apiv1"
	grafeasv1 "cloud.google.com/go/grafeas/apiv1"

	"github.com/docker/distribution/reference"
	"google.golang.org/api/iterator"
	grafeas "google.golang.org/genproto/googleapis/grafeas/v1"

	"github.com/Shopify/voucher"
	"github.com/Shopify/voucher/attestation"
	"github.com/Shopify/voucher/repository"
)

var errCannotAttest = errors.New("cannot create attestations, keyring is empty")

// Client implements voucher.MetadataClient, connecting to containeranalysis Grafeas.
type Client struct {
	containeranalysis *grafeasv1.Client // The client reference.
	keyring           *voucher.KeyRing  // The keyring used for signing metadata.
	binauthProject    string            // The project that Binauth Notes and Occurrences are written to.
	imageProject      string            // The project that image information is stored.
}

// CanAttest returns true if the client can create and sign attestations.
func (g *Client) CanAttest() bool {
	return nil != g.keyring
}

// NewPayloadBody returns a payload body appropriate for this MetadataClient.
func (g *Client) NewPayloadBody(reference reference.Canonical) (string, error) {
	payload, err := attestation.NewPayload(reference).ToString()
	if err != nil {
		return "", err
	}

	return payload, err
}

// AddAttestationToImage adds a new attestation with the passed AttestationPayload
// to the image described by ImageData.
func (g *Client) AddAttestationToImage(ctx context.Context, reference reference.Canonical, payload voucher.AttestationPayload) (interface{}, error) {
	if !g.CanAttest() {
		return nil, errCannotAttest
	}

	signed, keyID, err := payload.Sign(g.keyring)
	if nil != err {
		return nil, err
	}

	attestation := newOccurrenceAttestation(payload.Body, signed, keyID)
	occurrenceRequest := g.getCreateOccurrenceRequest(reference, payload.CheckName, attestation)
	occ, err := g.containeranalysis.CreateOccurrence(ctx, occurrenceRequest)

	if isAttestionExistsErr(err) {
		err = nil
		occ = nil
	}

	return occ, err
}

func (g *Client) getCreateOccurrenceRequest(reference reference.Canonical, parentNoteID string, attestation *grafeas.Occurrence_Attestation) *grafeas.CreateOccurrenceRequest {
	binauthProjectPath := "projects/" + g.binauthProject
	noteName := binauthProjectPath + "/notes/" + parentNoteID

	occurrence := grafeas.Occurrence{
		NoteName:    noteName,
		ResourceUri: "https://" + reference.Name() + "@" + reference.Digest().String(),
		Details:     attestation,
	}

	req := &grafeas.CreateOccurrenceRequest{Parent: binauthProjectPath, Occurrence: &occurrence}

	return req
}

// GetVulnerabilities returns the detected vulnerabilities for the Image described by voucher.ImageData.
func (g *Client) GetVulnerabilities(ctx context.Context, reference reference.Canonical) (vulnerabilities []voucher.Vulnerability, err error) {
	filterStr := kindFilterStr(reference, grafeas.NoteKind_VULNERABILITY)

	err = pollForDiscoveries(ctx, g, reference)
	if nil != err {
		return []voucher.Vulnerability{}, err
	}

	project := projectPath(g.imageProject)
	req := &grafeas.ListOccurrencesRequest{Parent: project, Filter: filterStr}
	occIterator := g.containeranalysis.ListOccurrences(ctx, req)

	for {
		var occ *grafeas.Occurrence

		occ, err = occIterator.Next()
		if nil != err {
			if iterator.Done == err {
				err = nil
			}

			break
		}

		vuln := OccurrenceToVulnerability(occ)
		vulnerabilities = append(vulnerabilities, vuln)
	}

	return
}

// Close closes the containeranalysis Grafeas client.
func (g *Client) Close() {
	g.containeranalysis.Close()
}

// GetBuildDetails gets BuildDetails for the passed image.
func (g *Client) GetBuildDetails(ctx context.Context, reference reference.Canonical) ([]repository.BuildDetail, error) {
	var buildDetails []repository.BuildDetail

	var err error

	filterStr := kindFilterStr(reference, grafeas.NoteKind_BUILD)

	project := projectPath(g.imageProject)
	req := &grafeas.ListOccurrencesRequest{Parent: project, Filter: filterStr}
	occIterator := g.containeranalysis.ListOccurrences(ctx, req)

	for {
		var occ *grafeas.Occurrence

		occ, err = occIterator.Next()
		if err != nil {
			if err == iterator.Done {
				err = nil
			}

			break
		}

		vuln := OccurrenceToBuildDetails(occ)
		buildDetails = append(buildDetails, vuln)
	}

	if len(buildDetails) == 0 && err == nil {
		err = &voucher.NoMetadataError{
			Type: voucher.VulnerabilityType,
			Err:  errNoOccurrences,
		}
	}

	if err != nil {
		return nil, err
	}

	return buildDetails, nil
}

// NewClient creates a new containeranalysis Grafeas Client.
func NewClient(ctx context.Context, imageProject, binauthProject string, keyring *voucher.KeyRing) (*Client, error) {
	var err error

	caClient, err := containeranalysisapi.NewClient(ctx)
	if err != nil {
		return nil, err
	}

	client := &Client{
		containeranalysis: caClient.GetGrafeasClient(),
		keyring:           keyring,
		binauthProject:    binauthProject,
		imageProject:      imageProject,
	}

	return client, nil
}
