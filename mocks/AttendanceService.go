// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	attendance "timesync-be/features/attendance"

	mock "github.com/stretchr/testify/mock"
)

// AttendanceService is an autogenerated mock type for the AttendanceService type
type AttendanceService struct {
	mock.Mock
}

// AttendanceFromAdmin provides a mock function with given fields: token, dateStart, dateEnd, attendanceType, employeeID
func (_m *AttendanceService) AttendanceFromAdmin(token interface{}, dateStart string, dateEnd string, attendanceType string, employeeID uint) error {
	ret := _m.Called(token, dateStart, dateEnd, attendanceType, employeeID)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}, string, string, string, uint) error); ok {
		r0 = rf(token, dateStart, dateEnd, attendanceType, employeeID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ClockIn provides a mock function with given fields: token, latitudeData, longitudeData
func (_m *AttendanceService) ClockIn(token interface{}, latitudeData string, longitudeData string) (attendance.Core, error) {
	ret := _m.Called(token, latitudeData, longitudeData)

	var r0 attendance.Core
	if rf, ok := ret.Get(0).(func(interface{}, string, string) attendance.Core); ok {
		r0 = rf(token, latitudeData, longitudeData)
	} else {
		r0 = ret.Get(0).(attendance.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(interface{}, string, string) error); ok {
		r1 = rf(token, latitudeData, longitudeData)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ClockOut provides a mock function with given fields: token, latitudeData, longitudeData
func (_m *AttendanceService) ClockOut(token interface{}, latitudeData string, longitudeData string) (attendance.Core, error) {
	ret := _m.Called(token, latitudeData, longitudeData)

	var r0 attendance.Core
	if rf, ok := ret.Get(0).(func(interface{}, string, string) attendance.Core); ok {
		r0 = rf(token, latitudeData, longitudeData)
	} else {
		r0 = ret.Get(0).(attendance.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(interface{}, string, string) error); ok {
		r1 = rf(token, latitudeData, longitudeData)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Record provides a mock function with given fields: token, dateFrom, dateTo
func (_m *AttendanceService) Record(token interface{}, dateFrom string, dateTo string) ([]attendance.Core, error) {
	ret := _m.Called(token, dateFrom, dateTo)

	var r0 []attendance.Core
	if rf, ok := ret.Get(0).(func(interface{}, string, string) []attendance.Core); ok {
		r0 = rf(token, dateFrom, dateTo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]attendance.Core)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(interface{}, string, string) error); ok {
		r1 = rf(token, dateFrom, dateTo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewAttendanceService interface {
	mock.TestingT
	Cleanup(func())
}

// NewAttendanceService creates a new instance of AttendanceService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewAttendanceService(t mockConstructorTestingTNewAttendanceService) *AttendanceService {
	mock := &AttendanceService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}