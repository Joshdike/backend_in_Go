package main

import (
	"log"
	"net/http"

	"github.com/Joshdike/backend_in_Go/beginner/palindrome-checker-service/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Post("/check", handlers.Check)

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}
}
