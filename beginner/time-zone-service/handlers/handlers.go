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
	location := query.Get("location")
	longitude := query.Get("longitude")
	latitude := query.Get("latitude")
	if location == "" && longitude == "" && latitude == "" {
		http.Error(w, `{"error":"location parameter is required"}`, http.StatusBadRequest)
		return
	}

	if location != "" {
		timezone, err := service.GetTimezone(location, 0, 0)
		if err != nil {
			http.Error(w, `{"error":"could not get timezone"}`, http.StatusInternalServerError)
			return
		}
		err = json.NewEncoder(w).Encode(models.TimezoneResponse{Location: timezone.Location, Timezone: timezone.Timezone, Utc_offset: timezone.Utc_offset, Current_time: timezone.Current_time})
		if err != nil {
			http.Error(w, `{"error":"failed to encode JSON response"}`, http.StatusInternalServerError)
			return
		}
	}
	lon, lat, err := stringToFloat64(longitude, latitude)
	if err != nil {
		http.Error(w, `{"error":"invalid longitude or latitude"}`, http.StatusBadRequest)
		return
	}
	if lon > 180 || lon < -180 {
		http.Error(w, `{"error":"invalid longitude"}`, http.StatusBadRequest)
		return
	}
	if lat > 90 || lat < -90 {
		http.Error(w, `{"error":"invalid latitude"}`, http.StatusBadRequest)
		return
	}

	timezone, err := service.GetTimezone(location, lon, lat)
	if err != nil {
		http.Error(w, `{"error":"could not get timezone"}`, http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(models.TimezoneResponse{Location: timezone.Location, Timezone: timezone.Timezone, Utc_offset: timezone.Utc_offset, Current_time: timezone.Current_time})
	if err != nil {
		log.Fatal(err)
	}
}

func stringToFloat64(longitude, latitude string) (float64, float64, error) {
	lon, err := strconv.ParseFloat(longitude, 64)
	if err != nil {
		return 0, 0, err
	}
	lat, err := strconv.ParseFloat(latitude, 64)
	if err != nil {
		return 0, 0, err
	}
	return lon, lat, nil
}
