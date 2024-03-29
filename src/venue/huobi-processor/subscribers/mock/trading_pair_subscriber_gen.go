// Code generated by MockGen. DO NOT EDIT.
// Source: trading_pair_subscriber.go

// Package mock_subscribers is a generated GoMock package.
package mock_subscribers

import (
	gomock "github.com/golang/mock/gomock"
)

// MockHuobiTradingPairSubscriber is a mock of HuobiTradingPairSubscriber interface.
type MockHuobiTradingPairSubscriber struct {
	ctrl     *gomock.Controller
	recorder *MockHuobiTradingPairSubscriberMockRecorder
}

// MockHuobiTradingPairSubscriberMockRecorder is the mock recorder for MockHuobiTradingPairSubscriber.
type MockHuobiTradingPairSubscriberMockRecorder struct {
	mock *MockHuobiTradingPairSubscriber
}

// NewMockHuobiTradingPairSubscriber creates a new mock instance.
func NewMockHuobiTradingPairSubscriber(ctrl *gomock.Controller) *MockHuobiTradingPairSubscriber {
	mock := &MockHuobiTradingPairSubscriber{ctrl: ctrl}
	mock.recorder = &MockHuobiTradingPairSubscriberMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHuobiTradingPairSubscriber) EXPECT() *MockHuobiTradingPairSubscriberMockRecorder {
	return m.recorder
}
