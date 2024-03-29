// Code generated by MockGen. DO NOT EDIT.
// Source: trading_pair_subscriber.go

// Package mock_subscribers is a generated GoMock package.
package mock_subscribers

import (
	gomock "github.com/golang/mock/gomock"
)

// MockBittrexTradingPairSubscriber is a mock of BittrexTradingPairSubscriber interface.
type MockBittrexTradingPairSubscriber struct {
	ctrl     *gomock.Controller
	recorder *MockBittrexTradingPairSubscriberMockRecorder
}

// MockBittrexTradingPairSubscriberMockRecorder is the mock recorder for MockBittrexTradingPairSubscriber.
type MockBittrexTradingPairSubscriberMockRecorder struct {
	mock *MockBittrexTradingPairSubscriber
}

// NewMockBittrexTradingPairSubscriber creates a new mock instance.
func NewMockBittrexTradingPairSubscriber(ctrl *gomock.Controller) *MockBittrexTradingPairSubscriber {
	mock := &MockBittrexTradingPairSubscriber{ctrl: ctrl}
	mock.recorder = &MockBittrexTradingPairSubscriberMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBittrexTradingPairSubscriber) EXPECT() *MockBittrexTradingPairSubscriberMockRecorder {
	return m.recorder
}
