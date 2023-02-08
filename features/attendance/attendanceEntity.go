package attendance

import "github.com/labstack/echo/v4"

type Core struct {
	ID               uint    `json:"id"`
	AttendanceDate   string  `json:"attendance_date"`
	ClockIn          string  `json:"clock_in"`
	ClockInLocation  string  `json:"clock_in_location"`
	ClockInOSM       string  `json:"clock_in_map_location"`
	ClockOut         string  `json:"clock_out"`
	ClockOutLocation string  `json:"clock_out_location"`
	ClockOutOSM      string  `json:"clock_out_map_location"`
	Attendance       string  `json:"attendance"`
	AttendanceStatus string  `json:"attendance_status"`
	WorkTime         float32 `json:"work_time"`
}

type AttendanceHandler interface {
	GetLL() echo.HandlerFunc
	ClockIn() echo.HandlerFunc
	ClockOut() echo.HandlerFunc
	AttendanceFromAdmin() echo.HandlerFunc
	Record() echo.HandlerFunc
	GetPresenceToday() echo.HandlerFunc
	GetPresenceTotalToday() echo.HandlerFunc
	GetPresenceDetail() echo.HandlerFunc
	RecordByID() echo.HandlerFunc
}

type AttendanceService interface {
	ClockIn(token interface{}, latitudeData string, longitudeData string) (Core, error)
	ClockOut(token interface{}, latitudeData string, longitudeData string) (Core, error)
	AttendanceFromAdmin(token interface{}, dateStart string, dateEnd string, attendanceType string, employeeID uint) error
	Record(token interface{}, dateFrom string, dateTo string) ([]Core, error)
	GetPresenceToday(token interface{}) (Core, error)
	GetPresenceTotalToday(token interface{}) ([]Core, error)
	GetPresenceDetail(token interface{}, attendanceID uint) (Core, error)
	RecordByID(employeeID uint, dateFrom string, dateTo string) ([]Core, string, error)
}

type AttendanceData interface {
	ClockIn(employeeID uint, latitudeData string, longitudeData string) (Core, error)
	ClockOut(employeeID uint, latitudeData string, longitudeData string) (Core, error)
	AttendanceFromAdmin(adminID uint, dateStart string, dateEnd string, attendanceType string, employeeID uint) error
	Record(employeeID uint, dateFrom string, dateTo string) ([]Core, string, error)
	GetPresenceToday(employeeID uint) (Core, error)
	GetPresenceTotalToday(adminID uint) ([]Core, error)
	GetPresenceDetail(adminID uint, attendanceID uint) (Core, error)
}
