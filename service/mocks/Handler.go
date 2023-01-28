// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	fiber "github.com/gofiber/fiber/v2"
	mock "github.com/stretchr/testify/mock"
)

// Handler is an autogenerated mock type for the Handler type
type Handler struct {
	mock.Mock
}

// CalculatePersonData provides a mock function with given fields: c
func (_m *Handler) CalculatePersonData(c *fiber.Ctx) error {
	ret := _m.Called(c)

	var r0 error
	if rf, ok := ret.Get(0).(func(*fiber.Ctx) error); ok {
		r0 = rf(c)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CalculateTaxCode provides a mock function with given fields: c
func (_m *Handler) CalculateTaxCode(c *fiber.Ctx) error {
	ret := _m.Called(c)

	var r0 error
	if rf, ok := ret.Get(0).(func(*fiber.Ctx) error); ok {
		r0 = rf(c)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// HandleError provides a mock function with given fields: c, err
func (_m *Handler) HandleError(c *fiber.Ctx, err error) error {
	ret := _m.Called(c, err)

	var r0 error
	if rf, ok := ret.Get(0).(func(*fiber.Ctx, error) error); ok {
		r0 = rf(c, err)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewHandler interface {
	mock.TestingT
	Cleanup(func())
}

// NewHandler creates a new instance of Handler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewHandler(t mockConstructorTestingTNewHandler) *Handler {
	mock := &Handler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
