package services

import (
	"errors"
	"log"
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
		} else if strings.Contains(err.Error(), "clockin time was expired") {
			return attendance.Core{}, errors.New("invalid clock in time request")
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
		if strings.Contains(err.Error(), "already") {
			return attendance.Core{}, errors.New("you already clock out today")
		} else if strings.Contains(err.Error(), "you dont have clock in data") {
			log.Println("must clock in")
			return attendance.Core{}, err
		} else if strings.Contains(err.Error(), "clock out time expired") {
			return attendance.Core{}, err
		} else {
			return attendance.Core{}, err
		}
	}
	return res, nil
}

// AttendanceFromAdmin implements attendance.AttendanceService
func (auc *attendanceUseCase) AttendanceFromAdmin(token interface{}, dateStart string, dateEnd string, attendanceType string, employeeID uint) error {
	adminID := helper.ExtractToken(token)
	err := auc.qry.AttendanceFromAdmin(uint(adminID), dateStart, dateEnd, attendanceType, employeeID)
	if err != nil {
		if strings.Contains(err.Error(), "wrong input format") {
			return errors.New("wrong input format")
		} else if strings.Contains(err.Error(), "access denied") {
			return errors.New("access denied")
		} else {
			return errors.New("server error")
		}
	}
	return nil
}

// Record implements attendance.AttendanceService
func (auc *attendanceUseCase) Record(token interface{}, dateFrom string, dateTo string) ([]attendance.Core, error) {
	userID := helper.ExtractToken(token)
	res, _, err := auc.qry.Record(uint(userID), dateFrom, dateTo)
	if err != nil {
		if strings.Contains(err.Error(), "wrong input format") {
			return []attendance.Core{}, errors.New("wrong input format")
		} else {
			return []attendance.Core{}, errors.New("server error")
		}

	}
	return res, nil
}

// GetPresenceToday implements attendance.AttendanceService
func (auc *attendanceUseCase) GetPresenceToday(token interface{}) (attendance.Core, error) {
	employeeID := helper.ExtractToken(token)
	res, err := auc.qry.GetPresenceToday(uint(employeeID))
	if err != nil {
		log.Println("data not found", err.Error())
		return attendance.Core{}, errors.New("data not found")
	}
	return res, nil
}

// GetPresenceTotalToday implements attendance.AttendanceService
func (auc *attendanceUseCase) GetPresenceTotalToday(token interface{}) ([]attendance.Core, error) {
	adminID := helper.ExtractToken(token)
	res, err := auc.qry.GetPresenceTotalToday(uint(adminID))
	if err != nil {
		log.Println("data not found", err.Error())
		return []attendance.Core{}, errors.New("data not found")
	}
	return res, nil
}

// GetPresenceDetail implements attendance.AttendanceService
func (auc *attendanceUseCase) GetPresenceDetail(token interface{}, attendanceID uint) (attendance.Core, error) {
	adminID := helper.ExtractToken(token)
	res, err := auc.qry.GetPresenceDetail(uint(adminID), attendanceID)
	if err != nil {
		log.Println("data not found", err.Error())
		return attendance.Core{}, errors.New("data not found")
	}
	return res, nil
}

// RecordByID implements attendance.AttendanceService
func (auc *attendanceUseCase) RecordByID(employeeID uint, dateFrom string, dateTo string) ([]attendance.Core, string, error) {
	res, nameUser, err := auc.qry.Record(uint(employeeID), dateFrom, dateTo)
	if err != nil {
		if strings.Contains(err.Error(), "wrong input format") {
			return []attendance.Core{}, "", errors.New("wrong input format")
		} else {
			return []attendance.Core{}, "", errors.New("server error")
		}

	}
	return res, nameUser, nil
}

// Graph implements attendance.AttendanceService
func (auc *attendanceUseCase) Graph(token interface{}, param string, yearMonth string) (interface{}, error) {
	adminID := helper.ExtractToken(token)
	res, err := auc.qry.Graph(uint(adminID), param, yearMonth)
	if err != nil {
		if strings.Contains(err.Error(), "access") {
			return []attendance.Core{}, errors.New("access denied")
		} else if strings.Contains(err.Error(), "type") {
			return []attendance.Core{}, errors.New("wrong type parameter")
		} else {
			return []attendance.Core{}, errors.New("server error")
		}

	}
	return res, nil
}
