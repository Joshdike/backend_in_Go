package main

import (
	"log"
	"net/http"

	"github.com/Joshdike/backend_in_Go/beginner/contact-form-api/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Post("/contact", handlers.HandleContactForm)

	log.Println("Server running on port 8080...")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Panic(err)
	}
}
