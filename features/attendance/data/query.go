package data

import (
	"timesync-be/features/attendance"

	"gorm.io/gorm"
)

type attendanceQuery struct {
	db *gorm.DB
}

// ClockIn implements attendance.AttendanceData
func (*attendanceQuery) ClockIn(userID uint, latitude interface{}, longitude interface{}) (attendance.Core, error) {
	panic("unimplemented")
}

// ClockOut implements attendance.AttendanceData
func (*attendanceQuery) ClockOut(userID uint) (attendance.Core, error) {
	panic("unimplemented")
}

func New(db *gorm.DB) attendance.AttendanceData {
	return &attendanceQuery{
		db: db,
	}
}
