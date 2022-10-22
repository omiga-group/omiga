// Code generated by MockGen. DO NOT EDIT.
// Source: currency_repository.go

// Package mock_repositories is a generated GoMock package.
package mock_repositories

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	models "github.com/omiga-group/omiga/src/venue/shared/models"
)

// MockCurrencyRepository is a mock of CurrencyRepository interface.
type MockCurrencyRepository struct {
	ctrl     *gomock.Controller
	recorder *MockCurrencyRepositoryMockRecorder
}

// MockCurrencyRepositoryMockRecorder is the mock recorder for MockCurrencyRepository.
type MockCurrencyRepositoryMockRecorder struct {
	mock *MockCurrencyRepository
}

// NewMockCurrencyRepository creates a new mock instance.
func NewMockCurrencyRepository(ctrl *gomock.Controller) *MockCurrencyRepository {
	mock := &MockCurrencyRepository{ctrl: ctrl}
	mock.recorder = &MockCurrencyRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCurrencyRepository) EXPECT() *MockCurrencyRepositoryMockRecorder {
	return m.recorder
}

// CreateCurrencies mocks base method.
func (m *MockCurrencyRepository) CreateCurrencies(ctx context.Context, currencies []models.Currency) (map[string]int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCurrencies", ctx, currencies)
	ret0, _ := ret[0].(map[string]int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCurrencies indicates an expected call of CreateCurrencies.
func (mr *MockCurrencyRepositoryMockRecorder) CreateCurrencies(ctx, currencies interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCurrencies", reflect.TypeOf((*MockCurrencyRepository)(nil).CreateCurrencies), ctx, currencies)
}

// CreateCurrency mocks base method.
func (m *MockCurrencyRepository) CreateCurrency(ctx context.Context, currency models.Currency) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCurrency", ctx, currency)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCurrency indicates an expected call of CreateCurrency.
func (mr *MockCurrencyRepositoryMockRecorder) CreateCurrency(ctx, currency interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCurrency", reflect.TypeOf((*MockCurrencyRepository)(nil).CreateCurrency), ctx, currency)
}
