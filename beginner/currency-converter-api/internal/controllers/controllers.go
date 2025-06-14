package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Joshdike/backend_in_Go/beginner/currency-converter-api/internal/models"
	"github.com/Joshdike/backend_in_Go/beginner/currency-converter-api/internal/service"
	"github.com/Joshdike/backend_in_Go/beginner/currency-converter-api/internal/storage"
)

type CurrencyController struct {
	service *service.CurrencyService
	storage *storage.CurrencyStorage
}

func NewCurrencyController(service *service.CurrencyService, sto *storage.CurrencyStorage) *CurrencyController {
	return &CurrencyController{service: service, storage: sto}
}

func (c *CurrencyController) ConvertCurrency(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var req models.ConversionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid request body", "details": err.Error()})
		return
	}
	result, err := c.service.Convert(req.From, req.To, req.Amount)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "currency conversion failed", "details": err.Error()})
		return
	}
	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		http.Error(w, `{"error":"unable to encode response"}`, http.StatusInternalServerError)
	}
}

func (c *CurrencyController) BulkConvertCurrencies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var req models.BulkConversionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid request body", "details": err.Error()})
		return
	}
	result, err := c.service.BulkConvert(req.From, req.To, req.Amount)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "currency conversions failed", "details": err.Error()})
		return
	}
	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		http.Error(w, `{"error":"unable to encode response"}`, http.StatusInternalServerError)
	}
}

func (c *CurrencyController) Rates(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	rates, err := c.storage.GetRates()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "failed to get rates", "details": err.Error()})
		return
	}

	result := models.RatesResponse{Rates: rates}
	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		http.Error(w, `{"error":"unable to encode response"}`, http.StatusInternalServerError)
	}

}

func (c *CurrencyController) UpdateRates(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var req models.UpdateRatesRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid request body", "details": err.Error()})
		return
	}
	err := c.storage.UpdateRates(req.Rates)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "rates update failed", "details": err.Error()})
		return
	}
	w.Write([]byte(`{"message":"successful"}`))
}
