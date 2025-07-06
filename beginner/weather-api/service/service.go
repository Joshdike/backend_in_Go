package service

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Joshdike/backend_in_Go/beginner/weather-api/models"
)

func GetWeather(latitude, longitude string) (models.WeatherData, error) {
	var data models.WeatherData
	res, err := http.Get("https://api.open-meteo.com/v1/forecast?latitude=" + latitude + "&longitude=" + longitude + "&current_weather=true")
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
	var data models.CityResponse
	link := fmt.Sprintf("https://nominatim.openstreetmap.org/reverse?format=json&lat=%s&lon=%s", latitude, longitude)
	res, err := http.Get(link)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		return "", err
	}
	city := data.Name + ", " + data.Address.Country
	return city, nil
}

func GetCoordinates(city string) (string, string, error) {
	var data models.Location
	res, err := http.Get("https://nominatim.openstreetmap.org/search?q=" + city + "&format=json")
	if err != nil {
		return "", "", err
	}
	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		return "", "", err
	}
	return data[0].Lat, data[0].Lon, nil
}
