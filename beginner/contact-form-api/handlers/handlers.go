package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Joshdike/backend_in_Go/beginner/contact-form-api/models"
	"github.com/Joshdike/backend_in_Go/beginner/contact-form-api/utils"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func HandleContactForm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var contactReq models.ContactRequest
	if err := json.NewDecoder(r.Body).Decode(&contactReq); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ContactResponse{
			Success: false,
			Message: "Invalid Request format",
		})
		return
	}

	if err := validate.Struct(contactReq); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ContactResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	if err := utils.SendContactEmail(contactReq); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ContactResponse{
			Success: false,
			Message: "Failed to process your message",
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.ContactResponse{
		Success: true,
		Message: "Message recieved successfully",
	})
}
