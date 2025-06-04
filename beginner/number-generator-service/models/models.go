package models

type RandRequest struct{
	Min int `json:"min"`
	Max int `json:"max"`
	Quantity int `json:"quantity"`
}

type RandResponse struct{
	Numbers []int `json:"numbers,omitempty"`
	Err string `json:"error,omitempty"`
}