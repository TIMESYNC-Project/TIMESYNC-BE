package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"timesync-be/features/attendance"

	"github.com/labstack/echo/v4"
)

type attendanceController struct {
	srv attendance.AttendanceService
}

func New(as attendance.AttendanceService) attendance.AttendanceHandler {
	return &attendanceController{
		srv: as,
	}
}

// ClockIn implements attendance.AttendanceHandler
func (ac *attendanceController) ClockIn() echo.HandlerFunc {
	return func(c echo.Context) error {
		longLat := LongitudeLatitude{}
		err := c.Bind(&longLat)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "input format inccorect"})
		}
		res, err := ac.srv.ClockIn(c.Get("user"), longLat.Latitude, longLat.Longitude)
		if err != nil {
			if strings.Contains(err.Error(), "you already clock in today") {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "clock in fail, you already clock in today"})
			} else if strings.Contains(err.Error(), "invalid") {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "clock in session has ended"})
			} else {
				return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "internal server error"})
			}
		}
		return c.JSON(http.StatusCreated, map[string]interface{}{
			"data":    res,
			"message": "clock in success",
		})
	}
}

// ClockOut implements attendance.AttendanceHandler
func (ac *attendanceController) ClockOut() echo.HandlerFunc {
	return func(c echo.Context) error {
		longLat := LongitudeLatitude{}
		err := c.Bind(&longLat)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "input format inccorect"})
		}
		res, err := ac.srv.ClockOut(c.Get("user"), longLat.Latitude, longLat.Longitude)
		if err != nil {
			if strings.Contains(err.Error(), "already clock out today") {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "clock out fail, you already clock out today"})
			} else if strings.Contains(err.Error(), "invalid") {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "clock out session has ended"})
			} else {
				return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "internal server error"})
			}
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"data":    res,
			"message": "clock out success",
		})
	}
}

// GetLL implements attendance.AttendanceHandler
func (*attendanceController) GetLL() echo.HandlerFunc {
	return func(c echo.Context) error {
		latitude, _ := strconv.ParseFloat(c.QueryParam("latitude"), 64)
		longitude, _ := strconv.ParseFloat(c.QueryParam("longitude"), 64)
		log.Println("cek data latitude, longitude", latitude, longitude)

		url := fmt.Sprintf("https://nominatim.openstreetmap.org/reverse?format=json&lat=%f&lon=%f&zoom=18&addressdetails=1", latitude, longitude)
		response, err := http.Get(url)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "error"})
		}

		defer response.Body.Close()
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "error"})
		}

		var nominatimResponse NominatimResponse
		err = json.Unmarshal(body, &nominatimResponse)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "error"})
		}
		city := nominatimResponse.Address.City
		street := nominatimResponse.Address.Road
		postcode := nominatimResponse.Address.Postcode
		state := nominatimResponse.Address.State
		country := nominatimResponse.Address.Country
		urlLocation := fmt.Sprintf("https://www.openstreetmap.org/#map=19/%f/%f", latitude, longitude)
		result := make(map[string]interface{})
		result["street"] = street
		result["city"] = city
		result["postal_code"] = postcode
		result["state"] = state
		result["country"] = country
		result["latitude"] = latitude
		result["longitude"] = longitude
		result["url_osm"] = urlLocation
		// Return the inserted location
		log.Println("Road : ", street)
		log.Println("City : ", city)
		log.Println("Postal Code:", postcode)
		log.Println("State:", state)
		log.Println("Country:", country)
		log.Println("Latitude: ", latitude)
		log.Println("Longitude: ", longitude)
		log.Println("location : ", urlLocation)
		// Return the inserted location
		return c.JSON(http.StatusOK, map[string]interface{}{
			"data":    result,
			"message": "success get location",
		})
	}
}

// AttendanceFromAdmin implements attendance.AttendanceHandler
func (ac *attendanceController) AttendanceFromAdmin() echo.HandlerFunc {
	return func(c echo.Context) error {
		paramID := c.Param("id")
		employeeID, _ := strconv.Atoi(paramID)
		input := CreateAttendance{}
		err := c.Bind(&input)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "input format inccorect"})
		}
		err = ac.srv.AttendanceFromAdmin(c.Get("user"), input.DateStart, input.DateEnd, input.Attendance, uint(employeeID))
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "user is not admin"})
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success create attendance",
		})
	}
}

// Record implements attendance.AttendanceHandler
func (ac *attendanceController) Record() echo.HandlerFunc {
	return func(c echo.Context) error {
		dateFrom := c.QueryParam("date_from")
		dateTo := c.QueryParam("date_to")
		res, err := ac.srv.Record(c.Get("user"), dateFrom, dateTo)
		if err != nil {
			if strings.Contains(err.Error(), "input format") {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "input format inccorect"})
			} else {
				return c.JSON(http.StatusNotFound, map[string]interface{}{"message": "data not found"})
			}
		}
		result := []RecordResponse{}
		for _, val := range res {
			result = append(result, RecordResponseCenvert(val))
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"data":    result,
			"message": "success show employee attendance record",
		})
	}
}

// GetPresenceToday implements attendance.AttendanceHandler
func (ac *attendanceController) GetPresenceToday() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := ac.srv.GetPresenceToday(c.Get("user"))
		if err != nil {
			return c.JSON(http.StatusNotFound, map[string]interface{}{"message": "data not found"})
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"data":    res,
			"message": "success show presence data today",
		})
	}
}

// GetPresenceTotalToday implements attendance.AttendanceHandler
func (ac *attendanceController) GetPresenceTotalToday() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := ac.srv.GetPresenceTotalToday(c.Get("user"))
		if err != nil {
			return c.JSON(http.StatusNotFound, map[string]interface{}{"message": "data not found"})
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"data":    res,
			"message": "success show all employe presence data today",
		})
	}
}

// GetPresenceDetail implements attendance.AttendanceHandler
func (ac *attendanceController) GetPresenceDetail() echo.HandlerFunc {
	return func(c echo.Context) error {
		paramID := c.Param("id")
		attendanceID, _ := strconv.Atoi(paramID)
		res, err := ac.srv.GetPresenceDetail(c.Get("user"), uint(attendanceID))
		if err != nil {
			return c.JSON(http.StatusNotFound, map[string]interface{}{"message": "data not found"})
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"data":    res,
			"message": "success show presence data detail",
		})
	}
}

// RecordByID implements attendance.AttendanceHandler
func (ac *attendanceController) RecordByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		dateFrom := c.QueryParam("date_from")
		dateTo := c.QueryParam("date_to")
		paramID := c.Param("id")
		employeeID, _ := strconv.Atoi(paramID)
		res, nameUser, err := ac.srv.RecordByID(uint(employeeID), dateFrom, dateTo)
		if err != nil {
			if strings.Contains(err.Error(), "input format") {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "input format inccorect"})
			} else {
				return c.JSON(http.StatusNotFound, map[string]interface{}{"message": "data not found"})
			}
		}
		result := []RecordResponse{}
		for _, val := range res {
			result = append(result, RecordResponseCenvert(val))
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"data": map[string]interface{}{
				"employee_name": nameUser,
				"record":        result,
			},
			"message": "success show employee attendance record",
		})
	}
}

// Graph implements attendance.AttendanceHandler
func (ac *attendanceController) Graph() echo.HandlerFunc {
	return func(c echo.Context) error {
		typeGrapgh := c.QueryParam("type")
		yearMonth := c.QueryParam("year_month")
		res, err := ac.srv.Graph(c.Get("user"), typeGrapgh, yearMonth)
		if err != nil {
			if strings.Contains(err.Error(), "access denied") {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": err.Error()})
			} else if strings.Contains(err.Error(), "type") {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": err.Error()})
			} else {
				return c.JSON(http.StatusNotFound, map[string]interface{}{"message": "data not found"})
			}
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"data":    res,
			"message": "success show graph data",
		})
	}
}
