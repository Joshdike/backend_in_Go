package service

import (
	"encoding/json"
	"net/http"

	"github.com/Joshdike/backend_in_Go/beginner/weather-api/models"
)

func GetWeather(latitude, longitude string) (models.WeatherData, error) {
	var data models.WeatherData
	res, err := http.Get("api.open-meteo.com/v1/forecast?latitude=" + latitude + "&longitude=" + longitude + "&current_weather=true")
	if err != nil {
		return models.WeatherData{}, err
	}
	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		return models.WeatherData{}, err
	}
	return data, nil
}

func GetCity(latitude, longitude string) (string, error) {
	var data models.Location
	res, err := http.Get("https://nominatim.openstreetmap.org/reverse?format=json&lat=" + latitude + "&lon=" + longitude)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		return "", err
	}
	return data.Address.City, nil
}
