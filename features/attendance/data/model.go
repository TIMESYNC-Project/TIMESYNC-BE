package data

import (
	"timesync-be/features/attendance"

	"gorm.io/gorm"
)

type Attendance struct {
	gorm.Model
	UserId           uint
	ClockIn          string
	ClockOut         string
	Attendance       string
	AttendanceStatus string
	WorkTime         int
	ClockInLocation  string
	ClockOutLocation string
	ClockInOSM       string
	ClockOutOSM      string
}

func DataToCore(data Attendance) attendance.Core {
	return attendance.Core{
		ID:               data.ID,
		ClockIn:          data.ClockIn,
		ClockOut:         data.ClockIn,
		Attendance:       data.Attendance,
		AttendanceStatus: data.AttendanceStatus,
		WorkTime:         data.WorkTime,
		ClockInLocation:  data.ClockInLocation,
		ClockOutLocation: data.ClockOutLocation,
		ClockInOSM:       data.ClockInOSM,
		ClockOutOSM:      data.ClockOutOSM,
	}
}

func CoreToData(core attendance.Core) Attendance {
	return Attendance{
		Model:            gorm.Model{ID: core.ID},
		ClockIn:          core.ClockIn,
		ClockOut:         core.ClockOut,
		Attendance:       core.Attendance,
		AttendanceStatus: core.AttendanceStatus,
		WorkTime:         core.WorkTime,
		ClockInLocation:  core.ClockInLocation,
		ClockOutLocation: core.ClockOutLocation,
		ClockInOSM:       core.ClockInOSM,
		ClockOutOSM:      core.ClockOutLocation,
	}
}
