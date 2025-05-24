package main

import (
	"log"
	"net/http"

	"github.com/Joshdike/backend_in_Go/beginner/hello-world-api/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	route := chi.NewRouter()

	route.Use(middleware.Logger)

	route.HandleFunc("/", handlers.HelloHandler)
	route.Get("/hello/{name}", handlers.Greetings)

	err := http.ListenAndServe(":8080", route)
	if err != nil {
		log.Fatal(err)
	}
}
