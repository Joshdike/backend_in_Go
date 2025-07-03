package models

type TimezoneResponse struct {
	Location     string `json:"location"`
	ZoneName    string `json:"zone_name"`
	Abbreviation string `json:"abbreviation"`
	Utc_offset   int `json:"utc_offset"`
	Current_time string `json:"current_time"`
}

type TimezoneData struct {
	Status           string `json:"status"`
	Message          string `json:"message"`
	CountryCode      string `json:"countryCode"`
	CountryName      string `json:"countryName"`
	RegionName       string `json:"regionName"`
	CityName         string `json:"cityName"`
	ZoneName         string `json:"zoneName"`
	Abbreviation     string `json:"abbreviation"`
	GMTOffset        int    `json:"gmtOffset"`
	DST              string `json:"dst"`
	ZoneStart        int64  `json:"zoneStart"`
	ZoneEnd          int64  `json:"zoneEnd"`
	NextAbbreviation string `json:"nextAbbreviation"`
	Timestamp        int64  `json:"timestamp"`
	Formatted        string `json:"formatted"`
}
