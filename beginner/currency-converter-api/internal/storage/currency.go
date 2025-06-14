package storage

import (
	"sync"
)

type CurrencyStorage struct {
	mu           sync.RWMutex
	currentRates map[string]float64
}

func NewCurrencyStorage() *CurrencyStorage {
	storage := &CurrencyStorage{
		currentRates: map[string]float64{
			"USD": 1.0,
			"EUR": 0.93,
			"GBP": 0.80,
			"JPY": 151.50,
			"CAD": 1.36,
			"AUD": 1.52,
			"CNY": 7.23,
			"INR": 83.30,
		},
	}
	return storage
}

func (s *CurrencyStorage) GetRates() (map[string]float64, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	rates := make(map[string]float64, len(s.currentRates))
	for k, v := range s.currentRates {
		rates[k] = v
	}
	return rates, nil
}

func (s *CurrencyStorage) UpdateRates(newRates map[string]float64) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	for currency, rate := range newRates {
		s.currentRates[currency] = rate
	}
	return nil
}
