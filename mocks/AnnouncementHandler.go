// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	echo "github.com/labstack/echo/v4"
	mock "github.com/stretchr/testify/mock"
)

// AnnouncementHandler is an autogenerated mock type for the AnnouncementHandler type
type AnnouncementHandler struct {
	mock.Mock
}

// DeleteAnnouncement provides a mock function with given fields:
func (_m *AnnouncementHandler) DeleteAnnouncement() echo.HandlerFunc {
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

// EmployeeInbox provides a mock function with given fields:
func (_m *AnnouncementHandler) EmployeeInbox() echo.HandlerFunc {
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

// GetAnnouncement provides a mock function with given fields:
func (_m *AnnouncementHandler) GetAnnouncement() echo.HandlerFunc {
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

// GetAnnouncementDetail provides a mock function with given fields:
func (_m *AnnouncementHandler) GetAnnouncementDetail() echo.HandlerFunc {
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

// PostAnnouncement provides a mock function with given fields:
func (_m *AnnouncementHandler) PostAnnouncement() echo.HandlerFunc {
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

type mockConstructorTestingTNewAnnouncementHandler interface {
	mock.TestingT
	Cleanup(func())
}

// NewAnnouncementHandler creates a new instance of AnnouncementHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewAnnouncementHandler(t mockConstructorTestingTNewAnnouncementHandler) *AnnouncementHandler {
	mock := &AnnouncementHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
