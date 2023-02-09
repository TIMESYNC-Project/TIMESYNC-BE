package handler

import "timesync-be/features/attendance"

type RecordResponse struct {
	ID               uint   `json:"id"`
	AttendanceDate   string `json:"attendance_date"`
	ClockIn          string `json:"clock_in"`
	ClockOut         string `json:"clock_out"`
	Attendance       string `json:"attendance"`
	AttendanceStatus string `json:"attendance_status"`
}

func RecordResponseCenvert(data attendance.Core) RecordResponse {
	return RecordResponse{
		ID:               data.ID,
		AttendanceDate:   data.AttendanceDate,
		ClockIn:          data.ClockIn,
		ClockOut:         data.ClockOut,
		Attendance:       data.Attendance,
		AttendanceStatus: data.AttendanceStatus,
	}
}
