package models

type CheckRequest struct {
	Password string `json:"password"`
}

type CheckResponse struct {
	Password    string   `json:"password"`
	Strength    string   `json:"strength"`
	Suggestions []string `json:"suggestions,omitempty"`
}
