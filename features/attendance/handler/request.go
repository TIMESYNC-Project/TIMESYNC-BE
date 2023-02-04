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
	Latitude  interface{} `json:"latitude" form:"latitude"`
	Longitude interface{} `json:"longitude" form:"longitude"`
}
