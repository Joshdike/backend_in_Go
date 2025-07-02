package handlers

import (
	"encoding/json"
	"net"
	"net/http"

	"github.com/Joshdike/backend_in_Go/beginner/ip-address-tracker-service/models"
)

func GeolocateHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ip := r.URL.Query().Get("ip")
	if ip == "" {
		http.Error(w, `{"error":"ip parameter is required"}`, http.StatusBadRequest)
		return
	}
	if !isValidIP(ip) {
		http.Error(w, `{"error":"invalid ip address"}`, http.StatusBadRequest)
		return
	}

	Geolocation, err := Geolocate(ip)
	if err != nil {
		http.Error(w, `{"error":"could not get geolocation"}`, http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(models.Location{IP: ip, City: Geolocation.City, Region: Geolocation.RegionName, Country: Geolocation.Country})
	if err != nil {
		http.Error(w, `{"error":"failed to encode JSON response"}`, http.StatusInternalServerError)
		return
	}

}

func Geolocate(ip string) (models.GeolocationData, error) {
	var Geolocation models.GeolocationData
	res, err := http.Get("http://ip-api.com/json/" + ip)
	if err != nil {
		return models.GeolocationData{}, err
	}
	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(&Geolocation)
	if err != nil {
		return models.GeolocationData{}, err
	}
	return Geolocation, nil
}

func isValidIP(ipAddress string) bool {
	ip := net.ParseIP(ipAddress)
	return ip != nil
}
