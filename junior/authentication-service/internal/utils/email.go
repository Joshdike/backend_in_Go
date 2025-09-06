package utils

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/Joshdike/backend_in_Go/junior/authentication_service/internal/config"
	"gopkg.in/gomail.v2"
)

type EmailConfig interface {
	GetFrontendURL() string
	GetFromEmail() string
	GetSMTPHost() string
	GetSMTPPort() string
	GetSMTPUser() string
	GetSMTPPassword() string
}

type EmailService interface {
	SendPasswordResetEmail(email, token string) error
}

type SMTPEmailService struct {
	config EmailConfig
}

func NewSMTPEmailService(config EmailConfig) *SMTPEmailService {
	return &SMTPEmailService{config: config}
}

func (s *SMTPEmailService) SendPasswordResetEmail(email, token string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", s.config.GetFromEmail())
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Password Reset Request")

	resetLink := fmt.Sprintf("%s/reset-password?token=%s", s.config.GetFrontendURL(), token)
	body := fmt.Sprintf(`
		<h2>Password Reset Request</h2>
		<p>You requested to reset your password. Click the link below to reset it:</p>
		<a href="%s">Reset Password</a>
		<p>This link will expire in 1 hour.</p>
		<p>If you didn't request this, please ignore this email.</p>
	`, resetLink)

	m.SetBody("text/html", body)

	portStr := s.config.GetSMTPPort()
	port, err := strconv.Atoi(portStr)
	if err != nil {
		port = 587
	}

	d := gomail.NewDialer(
		s.config.GetSMTPHost(),
		port,
		s.config.GetSMTPUser(),
		s.config.GetSMTPPassword(),
	)

	return d.DialAndSend(m)
}

type MockEmailService struct {
	config      EmailConfig
	SentEmails  []SentEmail
	ShouldError bool
	ErrorCount  int
}

type SentEmail struct {
	Email string
	Token string
	Body  string
}

func NewMockEmailService(config EmailConfig) *MockEmailService {
	return &MockEmailService{config: config}
}

func (m *MockEmailService) SendPasswordResetEmail(email, token string) error {
	if m.ShouldError {
		m.ErrorCount++
		return fmt.Errorf("failed to send email to %s", email)
	}

	resetLink := fmt.Sprintf("%s/reset-password?token=%s", m.config.GetFrontendURL(), token)

	body := fmt.Sprintf(`
		<h2>Password Reset Request</h2>
		<p>Reset token: %s</p>
		<a href="%s">Reset Password</a>
		<p>This link will expire in 1 hour.</p>
	`, token, resetLink)

	m.SentEmails = append(m.SentEmails, SentEmail{Email: email, Token: token, Body: body})
	return nil
}

type RealConfig struct{}

func (r *RealConfig) GetFrontendUrl() string {
	return config.GetFrontendURL()
}

func (r *RealConfig) GetFromEmail() string {
	return config.GetFromEmail()
}

func (r *RealConfig) GetSMTPHost() string {
	return config.GetSMTPHost()
}

func (r *RealConfig) GetSMTPPort() string {
	return config.GetSMTPPort()
}

func (r *RealConfig) GetSMTPUser() string {
	return config.GetSMTPUser()
}

func (r *RealConfig) GetSMTPPassword() string {
	return config.GetSMTPPassword()
}

func GenerateResetToken() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	token := make([]byte, 32)
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	for i := range token {
		token[i] = charset[r.Intn(len(charset))]
	}
	return string(token)
}
