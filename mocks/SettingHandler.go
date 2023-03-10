// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	echo "github.com/labstack/echo/v4"
	mock "github.com/stretchr/testify/mock"
)

// SettingHandler is an autogenerated mock type for the SettingHandler type
type SettingHandler struct {
	mock.Mock
}

// EditSetting provides a mock function with given fields:
func (_m *SettingHandler) EditSetting() echo.HandlerFunc {
	ret := _m.Called()

	var r0 echo.HandlerFunc
	if rf, ok := ret.Get(0).(func() echo.HandlerFunc); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(echo.HandlerFunc)
		}
	}

	return r0
}

// GetSetting provides a mock function with given fields:
func (_m *SettingHandler) GetSetting() echo.HandlerFunc {
	ret := _m.Called()

	var r0 echo.HandlerFunc
	if rf, ok := ret.Get(0).(func() echo.HandlerFunc); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(echo.HandlerFunc)
		}
	}

	return r0
}

type mockConstructorTestingTNewSettingHandler interface {
	mock.TestingT
	Cleanup(func())
}

// NewSettingHandler creates a new instance of SettingHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewSettingHandler(t mockConstructorTestingTNewSettingHandler) *SettingHandler {
	mock := &SettingHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
