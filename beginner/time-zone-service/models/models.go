package models

type TimezoneResponse struct {
	Location     string `json:"location"`
	Timezone     string `json:"timezone"`
	Utc_offset   string `json:"utc_offset"`
	Current_time string `json:"current_time"`
}

type TimezoneData struct {
	Location     string `json:"location"`
	Timezone     string `json:"timezone"`
	Region       string `json:"region"`
	Utc_offset   string `json:"utc_offset"`
	Current_time string `json:"current_time"`
}
