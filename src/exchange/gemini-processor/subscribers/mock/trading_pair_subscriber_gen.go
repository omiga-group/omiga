// Code generated by MockGen. DO NOT EDIT.
// Source: trading_pair_subscriber.go

// Package mock_subscribers is a generated GoMock package.
package mock_subscribers

import (
	gomock "github.com/golang/mock/gomock"
)

// MockTradingPairsSubscriber is a mock of TradingPairsSubscriber interface.
type MockTradingPairsSubscriber struct {
	ctrl     *gomock.Controller
	recorder *MockTradingPairsSubscriberMockRecorder
}

// MockTradingPairsSubscriberMockRecorder is the mock recorder for MockTradingPairsSubscriber.
type MockTradingPairsSubscriberMockRecorder struct {
	mock *MockTradingPairsSubscriber
}

// NewMockTradingPairsSubscriber creates a new mock instance.
func NewMockTradingPairsSubscriber(ctrl *gomock.Controller) *MockTradingPairsSubscriber {
	mock := &MockTradingPairsSubscriber{ctrl: ctrl}
	mock.recorder = &MockTradingPairsSubscriberMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTradingPairsSubscriber) EXPECT() *MockTradingPairsSubscriberMockRecorder {
	return m.recorder
}
