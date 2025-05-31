package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Joshdike/backend_in_Go/beginner/password-generator-service/models"
	"github.com/Joshdike/backend_in_Go/beginner/password-generator-service/service"
)

func GeneratePassword(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var req models.PasswordRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error":"invalid request body"}`, http.StatusBadRequest)
		return
	}
	resp, err := service.GeneratePassword(req.Length, req.IncludeUppercase, req.IncludeNumbers, req.IncludeSpecial)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	if err := json.NewEncoder(w).Encode(models.PasswordResponse{Password: resp}); err != nil {
		http.Error(w, `{"error":"unable to encode response"}`, http.StatusInternalServerError)
	}
}
