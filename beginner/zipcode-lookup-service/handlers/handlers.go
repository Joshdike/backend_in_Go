package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Joshdike/backend_in_Go/beginner/zipcode-lookup-service/models"
	"github.com/Joshdike/backend_in_Go/beginner/zipcode-lookup-service/service"
	"github.com/Joshdike/backend_in_Go/beginner/zipcode-lookup-service/utils"
)

func ZipcodeLookupHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var location models.Location
	zipcode := r.URL.Query().Get("zipcode")
	if zipcode == "" {
		http.Error(w, `{"error":"zipcode parameter is required"}`, http.StatusBadRequest)
		return
	}

	if !isValidZipcode(zipcode) {
		http.Error(w, `{"error":"invalid zipcode"}`, http.StatusBadRequest)
		return
	}
	if utils.HasZipcode(zipcode) {
		location = utils.GetLocation(zipcode)
		err := json.NewEncoder(w).Encode(location)
		if err != nil {
			http.Error(w, `{"error":"failed to encode JSON response"}`, http.StatusInternalServerError)
			return
		}
		return
	}
	data, err := service.GetLocation(zipcode)
	if err != nil {
		http.Error(w, `{"error":"could not get location"}`, http.StatusInternalServerError)
		return
	}
	location = models.Location{
		Zipcode: data.PostCode,
		City:    data.Places[0].PlaceName,
		State:   data.Places[0].State,
		Country: data.Country,
	}
	utils.SetLocation(zipcode, location)
	err = json.NewEncoder(w).Encode(location)
	if err != nil {
		http.Error(w, `{"error":"failed to encode JSON response"}`, http.StatusInternalServerError)
		return
	}
}

func isValidZipcode(zipcode string) bool {
	return len(zipcode) == 5
}
