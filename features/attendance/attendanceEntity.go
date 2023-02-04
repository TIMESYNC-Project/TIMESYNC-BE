package attendance

import "github.com/labstack/echo/v4"

type Core struct {
	ID               uint
	ClockIn          string
	ClockOut         string
	Attendance       string
	AttendanceStatus string
	WorkTime         int
	ClockInLocation  string
	ClockOutLocation string
	ClockInOSM       string
	ClockOutOSM      string
	User             User
}

type User struct {
	ID   uint
	Name string
	Nip  string
}

type AttendanceHandler interface {
	GetLL() echo.HandlerFunc
	ClockIn() echo.HandlerFunc
	ClockOut() echo.HandlerFunc
}

type AttendanceService interface {
	ClockIn(token interface{}, latitude interface{}, longitude interface{}) (Core, error)
	ClockOut(token interface{}) (Core, error)
}

type AttendanceData interface {
	ClockIn(userID uint, latitude interface{}, longitude interface{}) (Core, error)
	ClockOut(userID uint) (Core, error)
}
