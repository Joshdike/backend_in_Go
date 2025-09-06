package config

import (
	"os"
)

func GetFrontendURL() string {
	return os.Getenv("FRONTEND_URL")
}

func GetSMTPHost() string {
	return os.Getenv("SMTP_HOST")
}

func GetSMTPPort() string {
	return os.Getenv("SMTP_PORT")
}

func GetSMTPUser() string {
	return os.Getenv("SMTP_USER")
}

func GetSMTPPassword() string {
	return os.Getenv("SMTP_PASSWORD")
}

func GetFromEmail() string {
	return os.Getenv("FROM_EMAIL")
}
