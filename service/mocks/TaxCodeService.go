// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	service "taxcode-converter/service"

	mock "github.com/stretchr/testify/mock"
)

// TaxCodeService is an autogenerated mock type for the TaxCodeService type
type TaxCodeService struct {
	mock.Mock
}

// CalculatePersonData provides a mock function with given fields: req
func (_m *TaxCodeService) CalculatePersonData(req service.CalculatePersonDataRequest) (*service.CalculatePersonDataResponse, error) {
	ret := _m.Called(req)

	var r0 *service.CalculatePersonDataResponse
	if rf, ok := ret.Get(0).(func(service.CalculatePersonDataRequest) *service.CalculatePersonDataResponse); ok {
		r0 = rf(req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*service.CalculatePersonDataResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(service.CalculatePersonDataRequest) error); ok {
		r1 = rf(req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CalculateTaxCode provides a mock function with given fields: req
func (_m *TaxCodeService) CalculateTaxCode(req service.CalculateTaxCodeRequest) (*service.CalculateTaxCodeResponse, error) {
	ret := _m.Called(req)

	var r0 *service.CalculateTaxCodeResponse
	if rf, ok := ret.Get(0).(func(service.CalculateTaxCodeRequest) *service.CalculateTaxCodeResponse); ok {
		r0 = rf(req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*service.CalculateTaxCodeResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(service.CalculateTaxCodeRequest) error); ok {
		r1 = rf(req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewTaxCodeService interface {
	mock.TestingT
	Cleanup(func())
}

// NewTaxCodeService creates a new instance of TaxCodeService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTaxCodeService(t mockConstructorTestingTNewTaxCodeService) *TaxCodeService {
	mock := &TaxCodeService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
