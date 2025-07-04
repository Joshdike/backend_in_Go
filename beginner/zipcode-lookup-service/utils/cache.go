package utils

import "github.com/Joshdike/backend_in_Go/beginner/zipcode-lookup-service/models"

var cache = make(map[string]models.Location)

func GetLocation(zipcode string) models.Location {
	return cache[zipcode]
}

func SetLocation(zipcode string, data models.Location) {
	cache[zipcode] = data
}

func HasZipcode(zipcode string) bool {
	_, ok := cache[zipcode]
	return ok
}

func DeleteZipcode(zipcode string) {
	delete(cache, zipcode)
}

func ClearCache() {
	cache = make(map[string]models.Location)
}
