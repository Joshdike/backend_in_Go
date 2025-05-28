package handlers

import (
	"encoding/json"
	"net/http"
)

func Echo(w http.ResponseWriter, r *http.Request) {
	var req any
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, `{"error":"Invalid request body"}`, http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(req)
	if err != nil {
		http.Error(w, `{"error":"failed to encode JSON response"}`, http.StatusInternalServerError)
		return
	}
}
