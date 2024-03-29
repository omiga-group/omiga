// Code generated by MockGen. DO NOT EDIT.
// Source: trading_pair_subscriber.go

// Package mock_subscribers is a generated GoMock package.
package mock_subscribers

import (
	gomock "github.com/golang/mock/gomock"
)

// MockKucoinTradingPairSubscriber is a mock of KucoinTradingPairSubscriber interface.
type MockKucoinTradingPairSubscriber struct {
	ctrl     *gomock.Controller
	recorder *MockKucoinTradingPairSubscriberMockRecorder
}

// MockKucoinTradingPairSubscriberMockRecorder is the mock recorder for MockKucoinTradingPairSubscriber.
type MockKucoinTradingPairSubscriberMockRecorder struct {
	mock *MockKucoinTradingPairSubscriber
}

// NewMockKucoinTradingPairSubscriber creates a new mock instance.
func NewMockKucoinTradingPairSubscriber(ctrl *gomock.Controller) *MockKucoinTradingPairSubscriber {
	mock := &MockKucoinTradingPairSubscriber{ctrl: ctrl}
	mock.recorder = &MockKucoinTradingPairSubscriberMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockKucoinTradingPairSubscriber) EXPECT() *MockKucoinTradingPairSubscriberMockRecorder {
	return m.recorder
}
