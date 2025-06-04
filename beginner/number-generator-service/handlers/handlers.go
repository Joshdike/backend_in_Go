package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Joshdike/backend_in_Go/beginner/number-generator-service/models"
	"github.com/Joshdike/backend_in_Go/beginner/number-generator-service/service"
)

func GetRandomNumbers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var req models.RandRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.RandResponse{Err: "invalid request body"})
		return
	}
	numbers, err := service.RandomNumbers(req.Min, req.Max, req.Quantity)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.RandResponse{Err: err.Error()})
		return
	}
	err = json.NewEncoder(w).Encode(models.RandResponse{Numbers: numbers})
	if err != nil {
		http.Error(w, `{"error":"unable to encode response"}`, http.StatusInternalServerError)
	}
}
