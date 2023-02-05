package helper

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

type NominatimResponse struct {
	Address struct {
		City     string `json:"city"`
		Road     string `json:"road"`
		Postcode string `json:"postcode"`
		State    string `json:"state"`
		Country  string `json:"country"`
	} `json:"address"`
}

type AddressLoc struct {
	Location    string
	UrlLocation string
}

func FindLoc(latitudeData string, longitudeData string) (AddressLoc, error) {
	latitude, _ := strconv.ParseFloat(latitudeData, 64)
	longitude, _ := strconv.ParseFloat(longitudeData, 64)

	url := fmt.Sprintf("https://nominatim.openstreetmap.org/reverse?format=json&lat=%f&lon=%f&zoom=18&addressdetails=1", latitude, longitude)
	response, err := http.Get(url)
	if err != nil {
		return AddressLoc{}, err
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return AddressLoc{}, err
	}
	var nominatimResponse NominatimResponse
	err = json.Unmarshal(body, &nominatimResponse)
	if err != nil {
		return AddressLoc{}, err
	}
	city := nominatimResponse.Address.City
	street := nominatimResponse.Address.Road
	postcode := nominatimResponse.Address.Postcode
	state := nominatimResponse.Address.State
	country := nominatimResponse.Address.Country
	urlLocation := fmt.Sprintf("https://www.openstreetmap.org/#map=19/%f/%f", latitude, longitude)
	loc := ""
	if len(street) == 0 {
		loc = fmt.Sprintf("%s,%s,%s,%s", city, state, postcode, country)
	} else {
		loc = fmt.Sprintf("%s %s,%s,%s,%s", street, city, state, postcode, country)
	}
	result := AddressLoc{}
	result.Location = loc
	result.UrlLocation = urlLocation
	return result, nil

}

func GetTimeHourMinute() string {
	t := time.Now()
	hour := strconv.Itoa(t.Hour())
	minute := strconv.Itoa(t.Minute())

	if len(hour) == 1 {
		hour = "0" + hour
	}
	if len(minute) == 1 {
		minute = "0" + minute
	}
	result := fmt.Sprintf("%s:%s", hour, minute)
	return result
}
