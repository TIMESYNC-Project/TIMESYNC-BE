package services

import "timesync-be/features/attendance"

type attendanceUseCase struct {
	qry attendance.AttendanceData
}

func New(ad attendance.AttendanceData) attendance.AttendanceService {
	return &attendanceUseCase{
		qry: ad,
	}
}

// ClockIn implements attendance.AttendanceService
func (*attendanceUseCase) ClockIn(token interface{}, latitude interface{}, longitude interface{}) (attendance.Core, error) {
	panic("unimplemented")
}

// ClockOut implements attendance.AttendanceService
func (*attendanceUseCase) ClockOut(token interface{}) (attendance.Core, error) {
	panic("unimplemented")
}
