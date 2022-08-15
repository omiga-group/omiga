// Code generated by MockGen. DO NOT EDIT.
// Source: order-book-simulator.go

// Package mock_simulators is a generated GoMock package.
package mock_simulators

import (
	gomock "github.com/golang/mock/gomock"
)

// MockOrderBookSimulator is a mock of OrderBookSimulator interface.
type MockOrderBookSimulator struct {
	ctrl     *gomock.Controller
	recorder *MockOrderBookSimulatorMockRecorder
}

// MockOrderBookSimulatorMockRecorder is the mock recorder for MockOrderBookSimulator.
type MockOrderBookSimulatorMockRecorder struct {
	mock *MockOrderBookSimulator
}

// NewMockOrderBookSimulator creates a new mock instance.
func NewMockOrderBookSimulator(ctrl *gomock.Controller) *MockOrderBookSimulator {
	mock := &MockOrderBookSimulator{ctrl: ctrl}
	mock.recorder = &MockOrderBookSimulatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrderBookSimulator) EXPECT() *MockOrderBookSimulatorMockRecorder {
	return m.recorder
}