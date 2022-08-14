// Code generated by MockGen. DO NOT EDIT.
// Source: order-book-publisher.go

// Package mock_publishers is a generated GoMock package.
package mock_publishers

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	models "github.com/omiga-group/omiga/src/exchange/shared/models"
)

// MockOrderBookPublisher is a mock of OrderBookPublisher interface.
type MockOrderBookPublisher struct {
	ctrl     *gomock.Controller
	recorder *MockOrderBookPublisherMockRecorder
}

// MockOrderBookPublisherMockRecorder is the mock recorder for MockOrderBookPublisher.
type MockOrderBookPublisherMockRecorder struct {
	mock *MockOrderBookPublisher
}

// NewMockOrderBookPublisher creates a new mock instance.
func NewMockOrderBookPublisher(ctrl *gomock.Controller) *MockOrderBookPublisher {
	mock := &MockOrderBookPublisher{ctrl: ctrl}
	mock.recorder = &MockOrderBookPublisherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrderBookPublisher) EXPECT() *MockOrderBookPublisherMockRecorder {
	return m.recorder
}

// Publish mocks base method.
func (m *MockOrderBookPublisher) Publish(ctx context.Context, orderBook models.OrderBook) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Publish", ctx, orderBook)
	ret0, _ := ret[0].(error)
	return ret0
}

// Publish indicates an expected call of Publish.
func (mr *MockOrderBookPublisherMockRecorder) Publish(ctx, orderBook interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Publish", reflect.TypeOf((*MockOrderBookPublisher)(nil).Publish), ctx, orderBook)
}
