// Code generated by MockGen. DO NOT EDIT.
// Source: trading_pair_subscriber.go

// Package mock_subscribers is a generated GoMock package.
package mock_subscribers

import (
	gomock "github.com/golang/mock/gomock"
)

// MockBybitTradingPairSubscriber is a mock of BybitTradingPairSubscriber interface.
type MockBybitTradingPairSubscriber struct {
	ctrl     *gomock.Controller
	recorder *MockBybitTradingPairSubscriberMockRecorder
}

// MockBybitTradingPairSubscriberMockRecorder is the mock recorder for MockBybitTradingPairSubscriber.
type MockBybitTradingPairSubscriberMockRecorder struct {
	mock *MockBybitTradingPairSubscriber
}

// NewMockBybitTradingPairSubscriber creates a new mock instance.
func NewMockBybitTradingPairSubscriber(ctrl *gomock.Controller) *MockBybitTradingPairSubscriber {
	mock := &MockBybitTradingPairSubscriber{ctrl: ctrl}
	mock.recorder = &MockBybitTradingPairSubscriberMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBybitTradingPairSubscriber) EXPECT() *MockBybitTradingPairSubscriberMockRecorder {
	return m.recorder
}