// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	announcement "timesync-be/features/announcement"

	mock "github.com/stretchr/testify/mock"
)

// AnnouncementService is an autogenerated mock type for the AnnouncementService type
type AnnouncementService struct {
	mock.Mock
}

// DeleteAnnouncement provides a mock function with given fields: token, announcementID
func (_m *AnnouncementService) DeleteAnnouncement(token interface{}, announcementID uint) error {
	ret := _m.Called(token, announcementID)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}, uint) error); ok {
		r0 = rf(token, announcementID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAnnouncement provides a mock function with given fields:
func (_m *AnnouncementService) GetAnnouncement() ([]announcement.Core, error) {
	ret := _m.Called()

	var r0 []announcement.Core
	if rf, ok := ret.Get(0).(func() []announcement.Core); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]announcement.Core)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAnnouncementDetail provides a mock function with given fields: token, announcementID
func (_m *AnnouncementService) GetAnnouncementDetail(token interface{}, announcementID uint) (announcement.Core, error) {
	ret := _m.Called(token, announcementID)

	var r0 announcement.Core
	if rf, ok := ret.Get(0).(func(interface{}, uint) announcement.Core); ok {
		r0 = rf(token, announcementID)
	} else {
		r0 = ret.Get(0).(announcement.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(interface{}, uint) error); ok {
		r1 = rf(token, announcementID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PostAnnouncement provides a mock function with given fields: token, newAnnouncement
func (_m *AnnouncementService) PostAnnouncement(token interface{}, newAnnouncement announcement.Core) (announcement.Core, error) {
	ret := _m.Called(token, newAnnouncement)

	var r0 announcement.Core
	if rf, ok := ret.Get(0).(func(interface{}, announcement.Core) announcement.Core); ok {
		r0 = rf(token, newAnnouncement)
	} else {
		r0 = ret.Get(0).(announcement.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(interface{}, announcement.Core) error); ok {
		r1 = rf(token, newAnnouncement)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewAnnouncementService interface {
	mock.TestingT
	Cleanup(func())
}

// NewAnnouncementService creates a new instance of AnnouncementService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewAnnouncementService(t mockConstructorTestingTNewAnnouncementService) *AnnouncementService {
	mock := &AnnouncementService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}