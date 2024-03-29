// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	service "taxcode-converter/service"

	mock "github.com/stretchr/testify/mock"
)

// CsvProcessor is an autogenerated mock type for the CsvProcessor type
type CsvProcessor struct {
	mock.Mock
}

// CityFromCode provides a mock function with given fields: code
func (_m *CsvProcessor) CityFromCode(code string) *service.CityCSV {
	ret := _m.Called(code)

	var r0 *service.CityCSV
	if rf, ok := ret.Get(0).(func(string) *service.CityCSV); ok {
		r0 = rf(code)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*service.CityCSV)
		}
	}

	return r0
}

// CityFromPlace provides a mock function with given fields: place
func (_m *CsvProcessor) CityFromPlace(place service.Place) *service.CityCSV {
	ret := _m.Called(place)

	var r0 *service.CityCSV
	if rf, ok := ret.Get(0).(func(service.Place) *service.CityCSV); ok {
		r0 = rf(place)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*service.CityCSV)
		}
	}

	return r0
}

type mockConstructorTestingTNewCsvProcessor interface {
	mock.TestingT
	Cleanup(func())
}

// NewCsvProcessor creates a new instance of CsvProcessor. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCsvProcessor(t mockConstructorTestingTNewCsvProcessor) *CsvProcessor {
	mock := &CsvProcessor{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
