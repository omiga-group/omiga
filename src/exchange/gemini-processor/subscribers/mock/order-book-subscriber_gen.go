// Code generated by MockGen. DO NOT EDIT.
// Source: order-book-subscriber.go

// Package mock_subscribers is a generated GoMock package.
package mock_subscribers

import (
	gomock "github.com/golang/mock/gomock"
)

// MockGeminiOrderBookSubscriber is a mock of GeminiOrderBookSubscriber interface.
type MockGeminiOrderBookSubscriber struct {
	ctrl     *gomock.Controller
	recorder *MockGeminiOrderBookSubscriberMockRecorder
}

// MockGeminiOrderBookSubscriberMockRecorder is the mock recorder for MockGeminiOrderBookSubscriber.
type MockGeminiOrderBookSubscriberMockRecorder struct {
	mock *MockGeminiOrderBookSubscriber
}

// NewMockGeminiOrderBookSubscriber creates a new mock instance.
func NewMockGeminiOrderBookSubscriber(ctrl *gomock.Controller) *MockGeminiOrderBookSubscriber {
	mock := &MockGeminiOrderBookSubscriber{ctrl: ctrl}
	mock.recorder = &MockGeminiOrderBookSubscriberMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGeminiOrderBookSubscriber) EXPECT() *MockGeminiOrderBookSubscriberMockRecorder {
	return m.recorder
}
