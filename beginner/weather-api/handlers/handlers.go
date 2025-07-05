package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Joshdike/backend_in_Go/beginner/weather-api/models"
	"github.com/Joshdike/backend_in_Go/beginner/weather-api/service"
	"github.com/Joshdike/backend_in_Go/beginner/weather-api/utils"
)

func GetWeather(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	query := r.URL.Query()
	location := query.Get("loc")
	longitude := query.Get("lng")
	latitude := query.Get("lat")
	if location == "" && longitude == "" && latitude == "" {
		http.Error(w, `{"error":"location parameter is required"}`, http.StatusBadRequest)
		return
	}
	var weather models.WeatherResponse

	if (longitude == "" && latitude == "") || !isValidPosition(longitude, latitude) {
	}

	data, err := service.GetWeather(latitude, longitude)
	if err != nil {
		http.Error(w, `{"error":"could not get weather"}`, http.StatusInternalServerError)
		return
	}
	city, err := service.GetCity(latitude, longitude)
	if err != nil {
		http.Error(w, `{"error":"could not get city"}`, http.StatusInternalServerError)
		return
	}
	weather = models.WeatherResponse{
		City: city,
		Temp: data.CurrentWeather.Temperature,
		Windspeed: data.CurrentWeather.Windspeed,
		Description: GetDescription(data.CurrentWeather.Weathercode),
	}
	utils.SetCache(city, weather)
	err = json.NewEncoder(w).Encode(weather)
	if err != nil {
		http.Error(w, `{"error":"failed to encode JSON response"}`, http.StatusInternalServerError)
		return
	}
}

func isValidPosition(longitude, latitude string) bool {
	lon, err := strconv.ParseFloat(longitude, 64)
	if err != nil {
		return false
	}
	if lon > 180 || lon < -180 {
		return false
	}
	lat, err := strconv.ParseFloat(latitude, 64)
	if err != nil {
		return false
	}
	if lat > 90 || lat < -90 {
		return false
	}
	return true
}

func GetDescription(code int) string {
	switch code {
	case 0:
		return "Clear sky"
	case 1:
		return "Mainly clear"
	case 2:
		return "Partly cloudy"
	case 3:
		return "Overcast"
	case 45, 48:
		return "Fog"
	case 51, 53, 55:
		return "Drizzle"
	case 61, 63, 65, 66, 67:
		return "Rain"
	case 71, 73, 75, 77:
		return "Snow"
	case 80:
		return "Slight rain showers"
	case 81, 82:
		return "Heavy rain showers"
	case 95, 96, 99:
		return "Thunderstorms"
	default:
		return "Unknown weather condition"
	}
}
