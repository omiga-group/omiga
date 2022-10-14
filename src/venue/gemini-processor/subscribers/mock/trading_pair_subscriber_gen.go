// Code generated by MockGen. DO NOT EDIT.
// Source: trading_pair_subscriber.go

// Package mock_subscribers is a generated GoMock package.
package mock_subscribers

import (
	gomock "github.com/golang/mock/gomock"
)

// MockGeminiTradingPairSubscriber is a mock of GeminiTradingPairSubscriber interface.
type MockGeminiTradingPairSubscriber struct {
	ctrl     *gomock.Controller
	recorder *MockGeminiTradingPairSubscriberMockRecorder
}

// MockGeminiTradingPairSubscriberMockRecorder is the mock recorder for MockGeminiTradingPairSubscriber.
type MockGeminiTradingPairSubscriberMockRecorder struct {
	mock *MockGeminiTradingPairSubscriber
}

// NewMockGeminiTradingPairSubscriber creates a new mock instance.
func NewMockGeminiTradingPairSubscriber(ctrl *gomock.Controller) *MockGeminiTradingPairSubscriber {
	mock := &MockGeminiTradingPairSubscriber{ctrl: ctrl}
	mock.recorder = &MockGeminiTradingPairSubscriberMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGeminiTradingPairSubscriber) EXPECT() *MockGeminiTradingPairSubscriberMockRecorder {
	return m.recorder
}