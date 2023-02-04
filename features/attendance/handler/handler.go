package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
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
			return c.JSON(http.StatusBadRequest, "input format incorrect")
		}
		res, err := ac.srv.ClockIn(c.Get("user"), longLat.Latitude, longLat.Longitude)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "error"})
		}
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"data":    res,
			"message": "error",
		})
	}
}

// ClockOut implements attendance.AttendanceHandler
func (*attendanceController) ClockOut() echo.HandlerFunc {
	panic("unimplemented")
}

// GetLL implements attendance.AttendanceHandler
func (*attendanceController) GetLL() echo.HandlerFunc {
	return func(c echo.Context) error {
		latitude, _ := strconv.ParseFloat(c.QueryParam("latitude"), 64)
		longitude, _ := strconv.ParseFloat(c.QueryParam("longitude"), 64)

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
