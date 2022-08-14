// Code generated by MockGen. DO NOT EDIT.
// Source: producer_gen.go

// Package mock_syntheticorderv1 is a generated GoMock package.
package mock_syntheticorderv1

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	syntheticorderv1 "github.com/omiga-group/omiga/src/shared/clients/events/omiga/synthetic-order/v1"
)

// MockProducer is a mock of Producer interface.
type MockProducer struct {
	ctrl     *gomock.Controller
	recorder *MockProducerMockRecorder
}

// MockProducerMockRecorder is the mock recorder for MockProducer.
type MockProducerMockRecorder struct {
	mock *MockProducer
}

// NewMockProducer creates a new mock instance.
func NewMockProducer(ctrl *gomock.Controller) *MockProducer {
	mock := &MockProducer{ctrl: ctrl}
	mock.recorder = &MockProducerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProducer) EXPECT() *MockProducerMockRecorder {
	return m.recorder
}

// Produce mocks base method.
func (m *MockProducer) Produce(ctx context.Context, key string, event syntheticorderv1.SyntheticOrderEvent) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Produce", ctx, key, event)
	ret0, _ := ret[0].(error)
	return ret0
}

// Produce indicates an expected call of Produce.
func (mr *MockProducerMockRecorder) Produce(ctx, key, event interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Produce", reflect.TypeOf((*MockProducer)(nil).Produce), ctx, key, event)
}
