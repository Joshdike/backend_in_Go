package models

type PasswordRequest struct {
	Length           int  `json:"length"`
	IncludeUppercase bool `json:"include_uppercase"`
	IncludeNumbers   bool `json:"include_numbers"`
	IncludeSpecial   bool `json:"include_special"`
}

type PasswordResponse struct {
	Password string `json:"password"`
	Strength string `json:"strength"`
}
