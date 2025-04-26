package main

import (
	"log"
	"net/http"

	"github.com/Joshdike/backend_in_Go/beginner/calculator-api/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Post("/add", handlers.AddHandler)
	r.Post("/subtract", handlers.SubtractHandler)
	r.Post("/divide", handlers.DivideHandler)
	r.Post("/multiply", handlers.MultiplyHandler)
	r.Post("/exponent", handlers.ExponentHandler)
	r.Post("/nthroot", handlers.NthRootHandler)

	log.Println("Server starting on port 8080...")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}

}
