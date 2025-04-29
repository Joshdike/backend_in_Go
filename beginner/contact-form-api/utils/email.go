package utils

import (
	"log"

	"github.com/Joshdike/backend_in_Go/beginner/contact-form-api/models"
)

func SendContactEmail(contact models.ContactRequest) error {
	log.Printf(
		"Contact from %s <%s>: %q",
		contact.Name,
		contact.Email,
		contact.Message,
	)
	return nil
}
