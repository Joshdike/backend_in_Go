package service

import (
	"time"

	"github.com/Joshdike/backend_in_Go/beginner/currency-converter-api/internal/storage"

	"github.com/Joshdike/backend_in_Go/beginner/currency-converter-api/internal/models"
)

type CurrencyService struct {
	storage *storage.CurrencyStorage
}

func NewCurrencyService(storage *storage.CurrencyStorage) *CurrencyService {
	return &CurrencyService{storage: storage}
}

func (s *CurrencyService) Convert(from, to string, amount float64) (*models.ConversionResult, error) {
	rates, err := s.storage.GetRates()
	if err != nil {
		return nil, err
	}
	rate, exists := rates[to]
	if !exists {
		return nil, models.ErrCurrencyNotSupported
	}
	var converted_amount float64
	if from == to {
		converted_amount = amount
	} else if from == "USD" {
		converted_amount = amount * rate
	} else if to == "USD" {
		converted_amount = amount / rate
	} else {
		rateFrom, exists := rates[from]
		if !exists {
			return nil, models.ErrCurrencyNotSupported
		}
		converted_amount = amount / rateFrom * rate
	}

	return &models.ConversionResult{
		From:            from,
		To:              to,
		Amount:          amount,
		ConvertedAmount: converted_amount,
		Rate:            rate,
		Date:            time.Now().Format("02-01-2006"),
	}, nil
}

func (s *CurrencyService) BulkConvert(from string, to []string, amount float64) (*models.BulkConversionResult, error) {
	conversions := make([]models.ConversionResult, 0, len(to))
	for _, toCurrency := range to {
		conversion, err := s.Convert(from, toCurrency, amount)
		if err != nil {
			return nil, err
		}
		conversions = append(conversions, *conversion)
	}
	return &models.BulkConversionResult{
		From:        from,
		Amount:      amount,
		Conversions: conversions,
		Date:        time.Now().Format("02-01-2006"),
	}, nil
}
