package models

type WeatherResponse struct {
	City        string  `json:"city"`
	Temp        float64 `json:"temp"`
	Windspeed   float64 `json:"windspeed"`
	Description string  `json:"description"`
}

type CurrentWeather struct {
	Temperature   float64 `json:"temperature"`
	Windspeed     float64 `json:"windspeed"`
	Winddirection int     `json:"winddirection"`
	Weathercode   int     `json:"weathercode"`
	IsDay         int     `json:"is_day"`
}

type WeatherData struct {
	CurrentWeather CurrentWeather `json:"current_weather"`
}

type Location struct {
	Lat         string `json:"lat"`
	Lon         string `json:"lon"`
	DisplayName string `json:"display_name"`
	Address     struct {
		City    string `json:"city"`
		Country string `json:"country"`
	} `json:"address"`
}
