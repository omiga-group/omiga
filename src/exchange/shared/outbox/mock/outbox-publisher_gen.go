// Code generated by MockGen. DO NOT EDIT.
// Source: outbox-publisher.go

// Package mock_outbox is a generated GoMock package.
package mock_outbox

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	entities "github.com/omiga-group/omiga/src/exchange/shared/entities"
)

// MockOutboxPublisher is a mock of OutboxPublisher interface.
type MockOutboxPublisher struct {
	ctrl     *gomock.Controller
	recorder *MockOutboxPublisherMockRecorder
}

// MockOutboxPublisherMockRecorder is the mock recorder for MockOutboxPublisher.
type MockOutboxPublisherMockRecorder struct {
	mock *MockOutboxPublisher
}

// NewMockOutboxPublisher creates a new mock instance.
func NewMockOutboxPublisher(ctrl *gomock.Controller) *MockOutboxPublisher {
	mock := &MockOutboxPublisher{ctrl: ctrl}
	mock.recorder = &MockOutboxPublisherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOutboxPublisher) EXPECT() *MockOutboxPublisherMockRecorder {
	return m.recorder
}

// Publish mocks base method.
func (m *MockOutboxPublisher) Publish(ctx context.Context, transaction *entities.Tx, topic, key string, headers map[string]string, event interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Publish", ctx, transaction, topic, key, headers, event)
	ret0, _ := ret[0].(error)
	return ret0
}

// Publish indicates an expected call of Publish.
func (mr *MockOutboxPublisherMockRecorder) Publish(ctx, transaction, topic, key, headers, event interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Publish", reflect.TypeOf((*MockOutboxPublisher)(nil).Publish), ctx, transaction, topic, key, headers, event)
}

// PublishWithoutTransaction mocks base method.
func (m *MockOutboxPublisher) PublishWithoutTransaction(ctx context.Context, topic, key string, headers map[string]string, event interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PublishWithoutTransaction", ctx, topic, key, headers, event)
	ret0, _ := ret[0].(error)
	return ret0
}

// PublishWithoutTransaction indicates an expected call of PublishWithoutTransaction.
func (mr *MockOutboxPublisherMockRecorder) PublishWithoutTransaction(ctx, topic, key, headers, event interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublishWithoutTransaction", reflect.TypeOf((*MockOutboxPublisher)(nil).PublishWithoutTransaction), ctx, topic, key, headers, event)
}
