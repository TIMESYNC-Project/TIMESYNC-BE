package attendance

import "github.com/labstack/echo/v4"

type Core struct {
	ID               uint `json:"id"`
	AttendanceDate   string
	ClockIn          string
	ClockInLocation  string
	ClockInOSM       string
	ClockOut         string
	ClockOutLocation string
	ClockOutOSM      string
	Attendance       string
	AttendanceStatus string
	WorkTime         int
}

type AttendanceHandler interface {
	GetLL() echo.HandlerFunc
	ClockIn() echo.HandlerFunc
	ClockOut() echo.HandlerFunc
	AttendanceFromAdmin() echo.HandlerFunc
}

type AttendanceService interface {
	ClockIn(token interface{}, latitudeData string, longitudeData string) (Core, error)
	ClockOut(token interface{}, latitudeData string, longitudeData string) (Core, error)
	AttendanceFromAdmin(token interface{}, dateStart string, dateEnd string, attendanceType string, employeeID uint) error
}

type AttendanceData interface {
	ClockIn(employeeID uint, latitudeData string, longitudeData string) (Core, error)
	ClockOut(employeeID uint, latitudeData string, longitudeData string) (Core, error)
	AttendanceFromAdmin(adminID uint, dateStart string, dateEnd string, attendanceType string, employeeID uint) error
}
