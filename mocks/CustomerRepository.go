// Code generated by mockery v2.30.1. DO NOT EDIT.

package mocks

import (
	domain "github.com/brightnc/go-hexagonal/internal/core/domain"
	mock "github.com/stretchr/testify/mock"
)

// CustomerRepository is an autogenerated mock type for the CustomerRepository type
type CustomerRepository struct {
	mock.Mock
}

// CreateUser provides a mock function with given fields: user
func (_m *CustomerRepository) CreateUser(user domain.Customer) (*domain.Customer, error) {
	ret := _m.Called(user)

	var r0 *domain.Customer
	var r1 error
	if rf, ok := ret.Get(0).(func(domain.Customer) (*domain.Customer, error)); ok {
		return rf(user)
	}
	if rf, ok := ret.Get(0).(func(domain.Customer) *domain.Customer); ok {
		r0 = rf(user)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Customer)
		}
	}

	if rf, ok := ret.Get(1).(func(domain.Customer) error); ok {
		r1 = rf(user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAll provides a mock function with given fields:
func (_m *CustomerRepository) GetAll() ([]domain.Customer, error) {
	ret := _m.Called()

	var r0 []domain.Customer
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]domain.Customer, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []domain.Customer); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Customer)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetById provides a mock function with given fields: id
func (_m *CustomerRepository) GetById(id int) (*domain.Customer, error) {
	ret := _m.Called(id)

	var r0 *domain.Customer
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (*domain.Customer, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(int) *domain.Customer); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Customer)
		}
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewCustomerRepository creates a new instance of CustomerRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCustomerRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *CustomerRepository {
	mock := &CustomerRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
