package main

import (
	"log"
	"net/http"

	"github.com/Joshdike/backend_in_Go/beginner/temperature-converter-api/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Post("/convert", handlers.Convert)

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}
}
