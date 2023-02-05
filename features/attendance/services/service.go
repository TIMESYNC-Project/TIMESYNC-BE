package services

import (
	"errors"
	"strings"
	"timesync-be/features/attendance"
	"timesync-be/helper"
)

type attendanceUseCase struct {
	qry attendance.AttendanceData
}

func New(ad attendance.AttendanceData) attendance.AttendanceService {
	return &attendanceUseCase{
		qry: ad,
	}
}

// ClockIn implements attendance.AttendanceService
func (auc *attendanceUseCase) ClockIn(token interface{}, latitude string, longitude string) (attendance.Core, error) {
	employeeID := helper.ExtractToken(token)
	res, err := auc.qry.ClockIn(uint(employeeID), latitude, longitude)
	if err != nil {
		if strings.Contains(err.Error(), "you already clock in today") {
			return attendance.Core{}, errors.New("you already clock in today")
		} else {
			return attendance.Core{}, errors.New("server error")
		}
	}
	return res, nil
}

// ClockOut implements attendance.AttendanceService
func (auc *attendanceUseCase) ClockOut(token interface{}, latitude string, longitude string) (attendance.Core, error) {
	employeeID := helper.ExtractToken(token)
	res, err := auc.qry.ClockOut(uint(employeeID), latitude, longitude)
	if err != nil {
		if strings.Contains(err.Error(), "already clock out today") {
			return attendance.Core{}, errors.New("you already clock out today")
		} else {
			return attendance.Core{}, errors.New("server error")
		}
	}
	return res, nil
}

// AttendanceFromAdmin implements attendance.AttendanceService
func (auc *attendanceUseCase) AttendanceFromAdmin(token interface{}, dateStart string, dateEnd string, attendanceType string, employeeID uint) error {
	adminID := helper.ExtractToken(token)
	err := auc.qry.AttendanceFromAdmin(uint(adminID), dateStart, dateEnd, attendanceType, employeeID)
	if err != nil {
		if strings.Contains(err.Error(), "already clock out today") {
			return errors.New("already creating attendance")
		} else if strings.Contains(err.Error(), "wrong input format") {
			return errors.New("wrong input format")
		} else {
			return errors.New("server error")
		}
	}
	return nil
}

// Record implements attendance.AttendanceService
func (auc *attendanceUseCase) Record(token interface{}, dateFrom string, dateTo string) ([]attendance.Core, error) {
	userID := helper.ExtractToken(token)
	res, err := auc.qry.Record(uint(userID), dateFrom, dateTo)
	if err != nil {
		if strings.Contains(err.Error(), "wrong input format") {
			return []attendance.Core{}, errors.New("wrong input format")
		} else {
			return []attendance.Core{}, errors.New("server error")
		}

	}
	return res, nil
}
