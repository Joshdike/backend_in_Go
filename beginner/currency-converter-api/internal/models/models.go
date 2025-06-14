package models

import (
	"errors"
)

var (
	ErrCurrencyNotSupported = errors.New("currency not supported")
	ErrInvalidAmount        = errors.New("invalid amount")
)

type ConversionRequest struct {
	From   string  `json:"from"`
	To     string  `json:"to"`
	Amount float64 `json:"amount"`
}

type BulkConversionRequest struct {
	From   string   `json:"from"`
	To     []string `json:"to"`
	Amount float64  `json:"amount"`
}

type ConversionResult struct {
	From            string  `json:"from"`
	To              string  `json:"to"`
	Amount          float64 `json:"amount"`
	ConvertedAmount float64 `json:"converted_amount"`
	Rate            float64 `json:"rate"`
	Date            string  `json:"date"`
}

type BulkConversionResult struct {
	From        string             `json:"from"`
	Amount      float64            `json:"amount"`
	Conversions []ConversionResult `json:"conversions"`
	Date        string             `json:"date"`
}
type UpdateRatesRequest struct {
	Rates map[string]float64 `json:"rates"`
}

type RatesResponse struct {
	Rates map[string]float64 `json:"rates"`
}
