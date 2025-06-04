package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Joshdike/backend_in_Go/beginner/number-generator-service/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Post("/random", handlers.GetRandomNumbers)
	fmt.Println("server starting on port 8080...")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}
}
