package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Joshdike/backend_in_Go/beginner/weather-api/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Get("/weather", handlers.GetWeather)
	fmt.Println("server starting on port 8080...")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}
}
