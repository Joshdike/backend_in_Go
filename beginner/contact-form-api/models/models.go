package models

type ContactRequest struct {
	Name    string `json:"name" validate:"required,min=2,max=50"`
	Email   string `json:"email" validate:"required,email"`
	Message string `json:"message" validate:"required,min=10,max=500"`
}

type ContactResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
