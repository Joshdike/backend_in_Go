package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Joshdike/backend_in_Go/beginner/echo-api/handlers"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func main() {

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Post("/echo", handlers.Echo)
	fmt.Println("server starting on port 8080...")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}
}
