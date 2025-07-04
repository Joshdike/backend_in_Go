package service

import (
	"encoding/json"
	"net/http"

	"github.com/Joshdike/backend_in_Go/beginner/zipcode-lookup-service/models"
)

func GetLocation(zipcode string) (models.ApiData, error) {
	var data models.ApiData
	res, err := http.Get("https://api.zippopotam.us/us/" + zipcode)
	if err != nil {
		return models.ApiData{}, err
	}
	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		return models.ApiData{}, err
	}
	return data, nil
}
