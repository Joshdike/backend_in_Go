package models

type WeatherResponse struct {
	City        string  `json:"city"`
	Temp        float64 `json:"temp"`
	Windspeed   float64 `json:"windspeed"`
	Description string  `json:"description"`
}

type CurrentWeather struct {
	Temperature   float64 `json:"temperature"`
	Windspeed     float64 `json:"windspeed"`
	Winddirection int     `json:"winddirection"`
	Weathercode   int     `json:"weathercode"`
	IsDay         int     `json:"is_day"`
}

type WeatherData struct {
	CurrentWeather CurrentWeather `json:"current_weather"`
}

type Location []struct {
	PlaceID     int      `json:"place_id"`
	Licence     string   `json:"licence"`
	OsmType     string   `json:"osm_type"`
	OsmID       int      `json:"osm_id"`
	Lat         string   `json:"lat"`
	Lon         string   `json:"lon"`
	Class       string   `json:"class"`
	Type        string   `json:"type"`
	PlaceRank   int      `json:"place_rank"`
	Importance  float64  `json:"importance"`
	AddressType string   `json:"addresstype"`
	Name        string   `json:"name"`
	DisplayName string   `json:"display_name"`
	BoundingBox []string `json:"boundingbox"`
}

type CityResponse struct {
	PlaceID     int     `json:"place_id"`
	Licence     string  `json:"licence"`
	OsmType     string  `json:"osm_type"`
	OsmID       int     `json:"osm_id"`
	Lat         string  `json:"lat"`
	Lon         string  `json:"lon"`
	Class       string  `json:"class"`
	Type        string  `json:"type"`
	PlaceRank   int     `json:"place_rank"`
	Importance  float64 `json:"importance"`
	AddressType string  `json:"addresstype"`
	Name        string  `json:"name"`
	DisplayName string  `json:"display_name"`
	Address     struct {
		Road         string `json:"road"`
		Hamlet       string `json:"hamlet"`
		Municipality string `json:"municipality"`
		County       string `json:"county"`
		State        string `json:"state"`
		ISO3166Lvl4  string `json:"ISO3166-2-lvl4"`
		Region       string `json:"region"`
		Country      string `json:"country"`
		CountryCode  string `json:"country_code"`
	} `json:"address"`
	BoundingBox []string `json:"boundingbox"`
}
