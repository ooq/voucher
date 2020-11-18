// Code generated by MockGen. DO NOT EDIT.
// Source: grafeas/grafeas_service.go

// Package mock_grafeas is a generated GoMock package.
package mock_grafeas

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	objects "github.com/grafeas/voucher/grafeas/objects"
)

// MockGrafeasAPIService is a mock of GrafeasAPIService interface.
type MockGrafeasAPIService struct {
	ctrl     *gomock.Controller
	recorder *MockGrafeasAPIServiceMockRecorder
}

// MockGrafeasAPIServiceMockRecorder is the mock recorder for MockGrafeasAPIService.
type MockGrafeasAPIServiceMockRecorder struct {
	mock *MockGrafeasAPIService
}

// NewMockGrafeasAPIService creates a new mock instance.
func NewMockGrafeasAPIService(ctrl *gomock.Controller) *MockGrafeasAPIService {
	mock := &MockGrafeasAPIService{ctrl: ctrl}
	mock.recorder = &MockGrafeasAPIServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGrafeasAPIService) EXPECT() *MockGrafeasAPIServiceMockRecorder {
	return m.recorder
}

// CreateOccurrence mocks base method.
func (m *MockGrafeasAPIService) CreateOccurrence(arg0 context.Context, arg1 string, arg2 objects.Occurrence) (objects.Occurrence, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOccurrence", arg0, arg1, arg2)
	ret0, _ := ret[0].(objects.Occurrence)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOccurrence indicates an expected call of CreateOccurrence.
func (mr *MockGrafeasAPIServiceMockRecorder) CreateOccurrence(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOccurrence", reflect.TypeOf((*MockGrafeasAPIService)(nil).CreateOccurrence), arg0, arg1, arg2)
}

// ListNotes mocks base method.
func (m *MockGrafeasAPIService) ListNotes(arg0 context.Context, arg1 string, arg2 *objects.ListOpts) (objects.ListNotesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListNotes", arg0, arg1, arg2)
	ret0, _ := ret[0].(objects.ListNotesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListNotes indicates an expected call of ListNotes.
func (mr *MockGrafeasAPIServiceMockRecorder) ListNotes(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListNotes", reflect.TypeOf((*MockGrafeasAPIService)(nil).ListNotes), arg0, arg1, arg2)
}

// ListOccurrences mocks base method.
func (m *MockGrafeasAPIService) ListOccurrences(arg0 context.Context, arg1 string, arg2 *objects.ListOpts) (objects.ListOccurrencesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListOccurrences", arg0, arg1, arg2)
	ret0, _ := ret[0].(objects.ListOccurrencesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListOccurrences indicates an expected call of ListOccurrences.
func (mr *MockGrafeasAPIServiceMockRecorder) ListOccurrences(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListOccurrences", reflect.TypeOf((*MockGrafeasAPIService)(nil).ListOccurrences), arg0, arg1, arg2)
}
