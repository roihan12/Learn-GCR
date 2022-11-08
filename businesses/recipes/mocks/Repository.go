// Code generated by mockery v2.14.1. DO NOT EDIT.

package mocks

import (
	recipes "echo-recipe/businesses/recipes"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// Create provides a mock function with given fields: recipeDomain
func (_m *Repository) Create(recipeDomain *recipes.Domain) recipes.Domain {
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
func (_m *Repository) Delete(id string) bool {
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
func (_m *Repository) GetAll(name string) []recipes.Domain {
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
func (_m *Repository) GetByCategoryID(id string) []recipes.Domain {
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
func (_m *Repository) GetByID(id string) recipes.Domain {
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
func (_m *Repository) Update(id string, recipeDomain *recipes.Domain) recipes.Domain {
	ret := _m.Called(id, recipeDomain)

	var r0 recipes.Domain
	if rf, ok := ret.Get(0).(func(string, *recipes.Domain) recipes.Domain); ok {
		r0 = rf(id, recipeDomain)
	} else {
		r0 = ret.Get(0).(recipes.Domain)
	}

	return r0
}

type mockConstructorTestingTNewRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewRepository creates a new instance of Repository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRepository(t mockConstructorTestingTNewRepository) *Repository {
	mock := &Repository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}