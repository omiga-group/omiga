// Code generated by MockGen. DO NOT EDIT.
// Source: exchange-repository.go

// Package mock_repositories is a generated GoMock package.
package mock_repositories

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	models "github.com/omiga-group/omiga/src/exchange/shared/models"
)

// MockExchangeRepository is a mock of ExchangeRepository interface.
type MockExchangeRepository struct {
	ctrl     *gomock.Controller
	recorder *MockExchangeRepositoryMockRecorder
}

// MockExchangeRepositoryMockRecorder is the mock recorder for MockExchangeRepository.
type MockExchangeRepositoryMockRecorder struct {
	mock *MockExchangeRepository
}

// NewMockExchangeRepository creates a new mock instance.
func NewMockExchangeRepository(ctrl *gomock.Controller) *MockExchangeRepository {
	mock := &MockExchangeRepository{ctrl: ctrl}
	mock.recorder = &MockExchangeRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockExchangeRepository) EXPECT() *MockExchangeRepositoryMockRecorder {
	return m.recorder
}

// CreateExchange mocks base method.
func (m *MockExchangeRepository) CreateExchange(ctx context.Context, exchange models.Exchange) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateExchange", ctx, exchange)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateExchange indicates an expected call of CreateExchange.
func (mr *MockExchangeRepositoryMockRecorder) CreateExchange(ctx, exchange interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateExchange", reflect.TypeOf((*MockExchangeRepository)(nil).CreateExchange), ctx, exchange)
}

// CreateExchanges mocks base method.
func (m *MockExchangeRepository) CreateExchanges(ctx context.Context, exchanges []models.Exchange) (map[string]int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateExchanges", ctx, exchanges)
	ret0, _ := ret[0].(map[string]int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateExchanges indicates an expected call of CreateExchanges.
func (mr *MockExchangeRepositoryMockRecorder) CreateExchanges(ctx, exchanges interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateExchanges", reflect.TypeOf((*MockExchangeRepository)(nil).CreateExchanges), ctx, exchanges)
}
