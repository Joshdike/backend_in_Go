package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Joshdike/backend_in_Go/beginner/password-checker-service/models"
	"github.com/Joshdike/backend_in_Go/beginner/password-checker-service/services"
)

func CheckPassword(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var req models.CheckRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid request body", "details": err.Error()})
		return
	}
	strength, suggestions, err := services.StrengthChecker(req.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "password check failed", "details": err.Error()})
		return
	}
	if err = json.NewEncoder(w).Encode(models.CheckResponse{Password: req.Password, Strength: strength, Suggestions: suggestions}); err != nil {
		http.Error(w, `{"error":"unable to encode response"}`, http.StatusInternalServerError)
	}
}
