// Code generated by MockGen. DO NOT EDIT.
// Source: coingecko_coin_subscriber.go

// Package mock_subscribers is a generated GoMock package.
package mock_subscribers

import (
	gomock "github.com/golang/mock/gomock"
)

// MockCoingeckoCoinSubscriber is a mock of CoingeckoCoinSubscriber interface.
type MockCoingeckoCoinSubscriber struct {
	ctrl     *gomock.Controller
	recorder *MockCoingeckoCoinSubscriberMockRecorder
}

// MockCoingeckoCoinSubscriberMockRecorder is the mock recorder for MockCoingeckoCoinSubscriber.
type MockCoingeckoCoinSubscriberMockRecorder struct {
	mock *MockCoingeckoCoinSubscriber
}

// NewMockCoingeckoCoinSubscriber creates a new mock instance.
func NewMockCoingeckoCoinSubscriber(ctrl *gomock.Controller) *MockCoingeckoCoinSubscriber {
	mock := &MockCoingeckoCoinSubscriber{ctrl: ctrl}
	mock.recorder = &MockCoingeckoCoinSubscriberMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCoingeckoCoinSubscriber) EXPECT() *MockCoingeckoCoinSubscriberMockRecorder {
	return m.recorder
}
