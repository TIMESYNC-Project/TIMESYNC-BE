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
	AttendanceDate   string
	Attendance       string
	AttendanceStatus string
	WorkTime         int
	ClockInLocation  string
	ClockOutLocation string
	ClockInOSM       string
	ClockOutOSM      string
}

type User struct {
	gorm.Model
	ProfilePicture string
	Name           string
	BirthOfDate    string
	Nip            string `gorm:"not null"`
	Email          string `gorm:"unique"`
	Gender         string
	Position       string
	Phone          string
	Address        string
	Password       string
	Role           string
	AnnualLeave    int
}

type Setting struct {
	gorm.Model
	Start       string
	End         string
	Tolerance   int
	AnnualLeave int
}

func DataToCore(data Attendance) attendance.Core {
	return attendance.Core{
		ID:               data.ID,
		ClockIn:          data.ClockIn,
		ClockOut:         data.ClockOut,
		AttendanceDate:   data.AttendanceDate,
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
		AttendanceDate:   core.AttendanceDate,
		Attendance:       core.Attendance,
		AttendanceStatus: core.AttendanceStatus,
		WorkTime:         core.WorkTime,
		ClockInLocation:  core.ClockInLocation,
		ClockOutLocation: core.ClockOutLocation,
		ClockInOSM:       core.ClockInOSM,
		ClockOutOSM:      core.ClockOutLocation,
	}
}

type NominatimResponse struct {
	Address struct {
		City     string `json:"city"`
		Road     string `json:"road"`
		Postcode string `json:"postcode"`
		State    string `json:"state"`
		Country  string `json:"country"`
	} `json:"address"`
}
