// Code generated by MockGen. DO NOT EDIT.
// Source: outbox-background-service.go

// Package mock_outbox is a generated GoMock package.
package mock_outbox

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockOutboxBackgroundService is a mock of OutboxBackgroundService interface.
type MockOutboxBackgroundService struct {
	ctrl     *gomock.Controller
	recorder *MockOutboxBackgroundServiceMockRecorder
}

// MockOutboxBackgroundServiceMockRecorder is the mock recorder for MockOutboxBackgroundService.
type MockOutboxBackgroundServiceMockRecorder struct {
	mock *MockOutboxBackgroundService
}

// NewMockOutboxBackgroundService creates a new mock instance.
func NewMockOutboxBackgroundService(ctrl *gomock.Controller) *MockOutboxBackgroundService {
	mock := &MockOutboxBackgroundService{ctrl: ctrl}
	mock.recorder = &MockOutboxBackgroundServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOutboxBackgroundService) EXPECT() *MockOutboxBackgroundServiceMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockOutboxBackgroundService) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close.
func (mr *MockOutboxBackgroundServiceMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockOutboxBackgroundService)(nil).Close))
}

// RunAsync mocks base method.
func (m *MockOutboxBackgroundService) RunAsync() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RunAsync")
}

// RunAsync indicates an expected call of RunAsync.
func (mr *MockOutboxBackgroundServiceMockRecorder) RunAsync() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunAsync", reflect.TypeOf((*MockOutboxBackgroundService)(nil).RunAsync))
}
