// Code generated by MockGen. DO NOT EDIT.
// Source: http.go

// Package mock_http is a generated GoMock package.
package mock_http

import (
	http "net/http"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockHttpServer is a mock of HttpServer interface.
type MockHttpServer struct {
	ctrl     *gomock.Controller
	recorder *MockHttpServerMockRecorder
}

// MockHttpServerMockRecorder is the mock recorder for MockHttpServer.
type MockHttpServerMockRecorder struct {
	mock *MockHttpServer
}

// NewMockHttpServer creates a new mock instance.
func NewMockHttpServer(ctrl *gomock.Controller) *MockHttpServer {
	mock := &MockHttpServer{ctrl: ctrl}
	mock.recorder = &MockHttpServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHttpServer) EXPECT() *MockHttpServerMockRecorder {
	return m.recorder
}

// GetGraphQLHandler mocks base method.
func (m *MockHttpServer) GetGraphQLHandler() http.Handler {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGraphQLHandler")
	ret0, _ := ret[0].(http.Handler)
	return ret0
}

// GetGraphQLHandler indicates an expected call of GetGraphQLHandler.
func (mr *MockHttpServerMockRecorder) GetGraphQLHandler() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGraphQLHandler", reflect.TypeOf((*MockHttpServer)(nil).GetGraphQLHandler))
}

// GetHandler mocks base method.
func (m *MockHttpServer) GetHandler() http.Handler {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHandler")
	ret0, _ := ret[0].(http.Handler)
	return ret0
}

// GetHandler indicates an expected call of GetHandler.
func (mr *MockHttpServerMockRecorder) GetHandler() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHandler", reflect.TypeOf((*MockHttpServer)(nil).GetHandler))
}

// ListenAndServe mocks base method.
func (m *MockHttpServer) ListenAndServe() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListenAndServe")
	ret0, _ := ret[0].(error)
	return ret0
}

// ListenAndServe indicates an expected call of ListenAndServe.
func (mr *MockHttpServerMockRecorder) ListenAndServe() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListenAndServe", reflect.TypeOf((*MockHttpServer)(nil).ListenAndServe))
}
