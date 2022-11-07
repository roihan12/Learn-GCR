// Code generated by mockery v2.14.1. DO NOT EDIT.

package mocks

import (
	recipes "echo-recipe/businesses/recipes"

	mock "github.com/stretchr/testify/mock"
)

// Usecase is an autogenerated mock type for the Usecase type
type Usecase struct {
	mock.Mock
}

// Create provides a mock function with given fields: recipeDomain
func (_m *Usecase) Create(recipeDomain *recipes.Domain) recipes.Domain {
	ret := _m.Called(recipeDomain)

	var r0 recipes.Domain
	if rf, ok := ret.Get(0).(func(*recipes.Domain) recipes.Domain); ok {
		r0 = rf(recipeDomain)
	} else {
		r0 = ret.Get(0).(recipes.Domain)
	}

	return r0
}

// Delete provides a mock function with given fields: id
func (_m *Usecase) Delete(id string) bool {
	ret := _m.Called(id)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// GetAll provides a mock function with given fields: name
func (_m *Usecase) GetAll(name string) []recipes.Domain {
	ret := _m.Called(name)

	var r0 []recipes.Domain
	if rf, ok := ret.Get(0).(func(string) []recipes.Domain); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]recipes.Domain)
		}
	}

	return r0
}

// GetByCategoryID provides a mock function with given fields: id
func (_m *Usecase) GetByCategoryID(id string) []recipes.Domain {
	ret := _m.Called(id)

	var r0 []recipes.Domain
	if rf, ok := ret.Get(0).(func(string) []recipes.Domain); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]recipes.Domain)
		}
	}

	return r0
}

// GetByID provides a mock function with given fields: id
func (_m *Usecase) GetByID(id string) recipes.Domain {
	ret := _m.Called(id)

	var r0 recipes.Domain
	if rf, ok := ret.Get(0).(func(string) recipes.Domain); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(recipes.Domain)
	}

	return r0
}

// Update provides a mock function with given fields: id, recipeDomain
func (_m *Usecase) Update(id string, recipeDomain *recipes.Domain) recipes.Domain {
	ret := _m.Called(id, recipeDomain)

	var r0 recipes.Domain
	if rf, ok := ret.Get(0).(func(string, *recipes.Domain) recipes.Domain); ok {
		r0 = rf(id, recipeDomain)
	} else {
		r0 = ret.Get(0).(recipes.Domain)
	}

	return r0
}

type mockConstructorTestingTNewUsecase interface {
	mock.TestingT
	Cleanup(func())
}

// NewUsecase creates a new instance of Usecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUsecase(t mockConstructorTestingTNewUsecase) *Usecase {
	mock := &Usecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
