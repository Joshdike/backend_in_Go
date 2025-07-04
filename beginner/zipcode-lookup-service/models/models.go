package models

type Location struct {
	Zipcode string `json:"zipcode"`
	City    string `json:"city"`
	State   string `json:"state"`
	Country string `json:"country"`
}

type ApiData struct {
	PostCode    string  `json:"post code"`
	Country     string  `json:"country"`
	CountryAbbr string  `json:"country abbreviation"`
	Places      []Place `json:"places"`
}

type Place struct {
	PlaceName string `json:"place name"`
	Longitude string `json:"longitude"`
	State     string `json:"state"`
	StateAbbr string `json:"state abbreviation"`
	Latitude  string `json:"latitude"`
}
