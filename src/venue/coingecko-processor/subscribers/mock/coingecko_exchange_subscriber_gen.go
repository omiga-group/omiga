// Code generated by MockGen. DO NOT EDIT.
// Source: coingecko_exchange_subscriber.go

// Package mock_subscribers is a generated GoMock package.
package mock_subscribers

import (
	gomock "github.com/golang/mock/gomock"
)

// MockCoingeckoExchangeSubscriber is a mock of CoingeckoExchangeSubscriber interface.
type MockCoingeckoExchangeSubscriber struct {
	ctrl     *gomock.Controller
	recorder *MockCoingeckoExchangeSubscriberMockRecorder
}

// MockCoingeckoExchangeSubscriberMockRecorder is the mock recorder for MockCoingeckoExchangeSubscriber.
type MockCoingeckoExchangeSubscriberMockRecorder struct {
	mock *MockCoingeckoExchangeSubscriber
}

// NewMockCoingeckoExchangeSubscriber creates a new mock instance.
func NewMockCoingeckoExchangeSubscriber(ctrl *gomock.Controller) *MockCoingeckoExchangeSubscriber {
	mock := &MockCoingeckoExchangeSubscriber{ctrl: ctrl}
	mock.recorder = &MockCoingeckoExchangeSubscriberMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCoingeckoExchangeSubscriber) EXPECT() *MockCoingeckoExchangeSubscriberMockRecorder {
	return m.recorder
}
