package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/Joshdike/backend_in_Go/beginner/time-zone-service/models"
	"github.com/Joshdike/backend_in_Go/beginner/time-zone-service/service"
)

func TimezoneHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	query := r.URL.Query()
	location := query.Get("loc")
	longitude := query.Get("lng")
	latitude := query.Get("lat")
	if location == "" && longitude == "" && latitude == "" {
		http.Error(w, `{"error":"location parameter is required"}`, http.StatusBadRequest)
		return
	}

	if location != "" {
		timezone, err := service.GetTimezone(location, "", "")
		if err != nil {
			http.Error(w, `{"error":"could not get timezone"}`, http.StatusInternalServerError)
			return
		}
		Location := timezone.CityName + ", " + timezone.CountryName
		err = json.NewEncoder(w).Encode(models.TimezoneResponse{Location: Location, ZoneName: timezone.ZoneName, Abbreviation: timezone.Abbreviation, Utc_offset: timezone.GMTOffset, Current_time: timezone.Formatted})
		if err != nil {
			http.Error(w, `{"error":"failed to encode JSON response"}`, http.StatusInternalServerError)
			return
		}
	}

	if !isValidPosition(longitude, latitude) {
		http.Error(w, `{"error":"invalid longitude or latitude"}`, http.StatusBadRequest)
		return
	}

	timezone, err := service.GetTimezone(location, longitude, latitude)
	if err != nil {
		http.Error(w, `{"error":"could not get timezone"}`, http.StatusInternalServerError)
		return
	}
	Location := timezone.CityName + ", " + timezone.CountryName
	err = json.NewEncoder(w).Encode(models.TimezoneResponse{Location: Location, ZoneName: timezone.ZoneName, Abbreviation: timezone.Abbreviation, Utc_offset: timezone.GMTOffset, Current_time: timezone.Formatted})
	if err != nil {
		log.Fatal(err)
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
