package utils

import (
	"time"

	"github.com/Joshdike/backend_in_Go/beginner/weather-api/models"
)

var (
	cache = make(map[string]CacheItem)
)

type CacheItem struct {
	weather models.WeatherResponse
	Expiry  int64
}

func SetCache(city string, data models.WeatherResponse) {
	cache[city] = CacheItem{
		weather: data,
		Expiry:  time.Now().Unix() + 300,
	}
}

func GetCache(city string) models.WeatherResponse {
	return cache[city].weather
}

func IsCached(city string) bool {
	item, ok := cache[city]
	if !ok {
		return false
	}
	if item.Expiry < time.Now().Unix() {
		delete(cache, city)
		return false
	}
	return true
}
