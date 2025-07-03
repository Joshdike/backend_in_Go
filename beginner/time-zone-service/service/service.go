package service

import (
	"encoding/json"
	"net/http"

	"github.com/Joshdike/backend_in_Go/beginner/time-zone-service/models"
)

func GetTimezone(location, longitude, latitude string) (models.TimezoneData, error) {
	var timezone models.TimezoneData

	if location != "" {
		res, err := http.Get("http://api.timezonedb.com/v2.1/get-time-zone?format=json&key=YOUR_API_KEY&by=city&city=" + location)
		if err != nil {
			return models.TimezoneData{}, err
		}
		defer res.Body.Close()
		err = json.NewDecoder(res.Body).Decode(&timezone)
		if err != nil {
			return models.TimezoneData{}, err
		}
		return timezone, nil
	} else {
		res, err := http.Get("http://api.timezonedb.com/v2.1/get-time-zone?format=json&key=YOUR_API_KEY&by=position&lat=" + latitude + "&lng=" + longitude)
		if err != nil {
			return models.TimezoneData{}, err
		}
		defer res.Body.Close()
		err = json.NewDecoder(res.Body).Decode(&timezone)
		if err != nil {
			return models.TimezoneData{}, err
		}
		return timezone, nil
	}
}
