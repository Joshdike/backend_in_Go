package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Joshdike/backend_in_Go/beginner/currency-converter-api/internal/controllers"
	"github.com/Joshdike/backend_in_Go/beginner/currency-converter-api/internal/service"
	"github.com/Joshdike/backend_in_Go/beginner/currency-converter-api/internal/storage"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	sto := storage.NewCurrencyStorage()
	s := service.NewCurrencyService(sto)
	c := controllers.NewCurrencyController(s, sto)
	r.Post("/convert", c.ConvertCurrency)
	r.Post("/bulkconvert", c.BulkConvertCurrencies)
	r.Get("/rates", c.Rates)
	r.Post("/updaterates", c.UpdateRates)

	port := ":8080"
	fmt.Printf("server starting on port %s...", port)
	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatal(err)
	}
}
