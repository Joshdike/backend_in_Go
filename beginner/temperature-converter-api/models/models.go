package models

type ConversionRequest struct{
	From string `json:"from"`
	To string `json:"to"`
	Val float64 `json:"val"`
}

