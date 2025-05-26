package helpers

import (
	"encoding/json"
	"net/http"
	"unicode"

	"github.com/Joshdike/backend_in_Go/beginner/palindrome-checker-service/models"
)

func IsPalindrome(word string) bool {
	l := 0
	r := len(word) - 1
	for l < r {
		left := rune(word[l])
		right := rune(word[r])
		if !unicode.IsLetter(left) && !unicode.IsDigit(left) {
			l++
			continue
		}
		if !unicode.IsLetter(right) && !unicode.IsDigit(right) {
			r--
			continue
		}
		if unicode.ToLower(left) != unicode.ToLower(right) {
			return false
		}
		l++
		r--
	}
	return true
}

func ErrorResponse(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	err = json.NewEncoder(w).Encode(models.CheckResponse{IsPalindrome: false, Err: err.Error()})
	if err != nil {
		http.Error(w, `{Err: "failed to encode error response"}`, http.StatusInternalServerError)
	}

}
