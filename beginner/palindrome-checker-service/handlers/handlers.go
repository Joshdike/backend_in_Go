package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/Joshdike/backend_in_Go/beginner/palindrome-checker-service/helpers"
	"github.com/Joshdike/backend_in_Go/beginner/palindrome-checker-service/models"
)

func Check(w http.ResponseWriter, r *http.Request) {
	var req models.CheckRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		helpers.ErrorResponse(w, fmt.Errorf("invalid request body: %v", err))
		return
	}
	if strings.TrimSpace(req.Str) == "" {
		helpers.ErrorResponse(w, errors.New("word cannot be empty"))
		return
	}
	result := helpers.IsPalindrome(req.Str)
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(models.CheckResponse{IsPalindrome: result})
	if err != nil {
		helpers.ErrorResponse(w, err)
		return
	}
}
