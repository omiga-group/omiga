// Code generated by MockGen. DO NOT EDIT.
// Source: order-book-subscriber.go

// Package mock_subscribers is a generated GoMock package.
package mock_subscribers

import (
	gomock "github.com/golang/mock/gomock"
)

// MockFtxOrderBookSubscriber is a mock of FtxOrderBookSubscriber interface.
type MockFtxOrderBookSubscriber struct {
	ctrl     *gomock.Controller
	recorder *MockFtxOrderBookSubscriberMockRecorder
}

// MockFtxOrderBookSubscriberMockRecorder is the mock recorder for MockFtxOrderBookSubscriber.
type MockFtxOrderBookSubscriberMockRecorder struct {
	mock *MockFtxOrderBookSubscriber
}

// NewMockFtxOrderBookSubscriber creates a new mock instance.
func NewMockFtxOrderBookSubscriber(ctrl *gomock.Controller) *MockFtxOrderBookSubscriber {
	mock := &MockFtxOrderBookSubscriber{ctrl: ctrl}
	mock.recorder = &MockFtxOrderBookSubscriberMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFtxOrderBookSubscriber) EXPECT() *MockFtxOrderBookSubscriberMockRecorder {
	return m.recorder
}
