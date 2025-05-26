package models

type CheckRequest struct {
	Str string `json:"str"`
}

type CheckResponse struct {
	IsPalindrome bool `json:"ispalindrome"`
	Err string `json:"err,omitempty"`
}

