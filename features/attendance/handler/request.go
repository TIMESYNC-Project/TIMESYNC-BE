package handler

type NominatimResponse struct {
	Address struct {
		City     string `json:"city"`
		Road     string `json:"road"`
		Postcode string `json:"postcode"`
		State    string `json:"state"`
		Country  string `json:"country"`
	} `json:"address"`
}

type LongitudeLatitude struct {
	Latitude  string `json:"latitude" form:"latitude"`
	Longitude string `json:"longitude" form:"longitude"`
}

type CreateAttendance struct {
	Attendance string `json:"attendance" form:"attendance"`
	DateStart  string `json:"date_start" form:"date_start"`
	DateEnd    string `json:"date_end" form:"date_end"`
}

type RecordRequest struct {
	DateFrom string `json:"date_from" form:"date_from"`
	DateTo   string `json:"date_to" form:"date_to"`
}
