// Code generated by mockery v2.30.1. DO NOT EDIT.

package mocks

import (
	domain "github.com/brightnc/go-hexagonal/internal/core/domain"
	mock "github.com/stretchr/testify/mock"
)

// AccountRepository is an autogenerated mock type for the AccountRepository type
type AccountRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: _a0
func (_m *AccountRepository) Create(_a0 domain.Account) (*domain.Account, error) {
	ret := _m.Called(_a0)

	var r0 *domain.Account
	var r1 error
	if rf, ok := ret.Get(0).(func(domain.Account) (*domain.Account, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(domain.Account) *domain.Account); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Account)
		}
	}

	if rf, ok := ret.Get(1).(func(domain.Account) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAll provides a mock function with given fields: customerId
func (_m *AccountRepository) GetAll(customerId int) ([]domain.Account, error) {
	ret := _m.Called(customerId)

	var r0 []domain.Account
	var r1 error
	if rf, ok := ret.Get(0).(func(int) ([]domain.Account, error)); ok {
		return rf(customerId)
	}
	if rf, ok := ret.Get(0).(func(int) []domain.Account); ok {
		r0 = rf(customerId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Account)
		}
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(customerId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewAccountRepository creates a new instance of AccountRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAccountRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *AccountRepository {
	mock := &AccountRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}