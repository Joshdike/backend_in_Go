package utils

import (
	"regexp"
	"strings"
	"testing"
)

type TestConfig struct {
	FrontendURL  string
	FromEmail    string
	SMTPHost     string
	SMTPPort     string
	SMTPUser     string
	SMTPPassword string
}

func (t *TestConfig) GetFrontendURL() string  { return t.FrontendURL }
func (t *TestConfig) GetFromEmail() string    { return t.FromEmail }
func (t *TestConfig) GetSMTPHost() string     { return t.SMTPHost }
func (t *TestConfig) GetSMTPPort() string     { return t.SMTPPort }
func (t *TestConfig) GetSMTPUser() string     { return t.SMTPUser }
func (t *TestConfig) GetSMTPPassword() string { return t.SMTPPassword }

func TestMockEmailService(t *testing.T) {
	testConfig := &TestConfig{
		FrontendURL:  "http://localhost:3000",
		FromEmail:    "test@example.com",
		SMTPHost:     "smtp.example.com",
		SMTPPort:     "587",
		SMTPUser:     "testuser",
		SMTPPassword: "testpassword",
	}
	t.Run("sends email succesfully", func(t *testing.T) {
		mockService := NewMockEmailService(testConfig)
		email := "user@example.com"
		token := "testtoken123"

		err := mockService.SendPasswordResetEmail(email, token)
		if err != nil {
			t.Errorf("expected no error but got %v", err)
		}

		if len(mockService.SentEmails) != 1 {
			t.Errorf("expected 1 email to be sent but got %d", len(mockService.SentEmails))
		}

		sentEmail := mockService.SentEmails[0]
		if sentEmail.Email != email {
			t.Errorf("expected email %q but got %q", email, sentEmail.Email)
		}
		if sentEmail.Token != token {
			t.Errorf("expected token %q but got %q", token, sentEmail.Token)
		}

		expectedLink := "http://localhost:3000/reset-password?token=testtoken123"
		if !strings.Contains(sentEmail.Body, expectedLink) {
			t.Errorf("expected email body to contain link %q but got %q", expectedLink, sentEmail.Body)
		}
	})
	t.Run("handles multiple emails", func(t *testing.T) {
		mockService := NewMockEmailService(testConfig)
		emails := []string{"user1@example.com", "user2@example.com"}
		tokens := []string{"token1", "token2"}

		for i, email := range emails {
			err := mockService.SendPasswordResetEmail(email, tokens[i])
			if err != nil {
				t.Errorf("Failed to send email %d: %v", i, err)
			}
		}

		if len(mockService.SentEmails) != 2 {
			t.Errorf("Expected 2 sent emails, got %d", len(mockService.SentEmails))
		}
	})

	t.Run("handles errors when configured", func(t *testing.T) {
		mockService := NewMockEmailService(testConfig)
		mockService.ShouldError = true
		email := "user@example.com"
		token := "test_token"

		err := mockService.SendPasswordResetEmail(email, token)

		if err == nil {
			t.Error("Expected error, got nil")
		}

		expectedError := "failed to send email to user@example.com"
		if err.Error() != expectedError {
			t.Errorf("Expected error '%s', got '%s'", expectedError, err.Error())
		}

		if len(mockService.SentEmails) != 0 {
			t.Errorf("Expected no sent emails when error occurs, got %d", len(mockService.SentEmails))
		}
	})
}

func TestEmailContentWithDifferentConfigs(t *testing.T) {
	tests := []struct {
		name     string
		config   *TestConfig
		token    string
		expected string
	}{
		{
			name: "development environment",
			config: &TestConfig{
				FrontendURL: "http://localhost:3000",
				FromEmail:   "dev@example.com",
			},
			token:    "dev_token",
			expected: "http://localhost:3000/reset-password?token=dev_token",
		},
		{
			name: "production environment",
			config: &TestConfig{
				FrontendURL: "https://myapp.com",
				FromEmail:   "noreply@myapp.com",
			},
			token:    "prod_token",
			expected: "https://myapp.com/reset-password?token=prod_token",
		},
		{
			name: "staging environment",
			config: &TestConfig{
				FrontendURL: "https://staging.myapp.com",
				FromEmail:   "staging@myapp.com",
			},
			token:    "staging_token",
			expected: "https://staging.myapp.com/reset-password?token=staging_token",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockService := NewMockEmailService(tt.config)
			email := "test@example.com"

			err := mockService.SendPasswordResetEmail(email, tt.token)
			if err != nil {
				t.Fatalf("Failed to send email: %v", err)
			}

			sentEmail := mockService.SentEmails[0]
			if !strings.Contains(sentEmail.Body, tt.expected) {
				t.Errorf("Expected reset link %s in email body", tt.expected)
			}
		})
	}
}

func TestGenerateResetToken(t *testing.T) {
	t.Run("generates unique token", func(t *testing.T) {
		token1 := GenerateResetToken()
		token2 := GenerateResetToken()

		if token1 == token2 {
			t.Error("expected unique tokens but got the same")
		}

	})
	t.Run("generates correct length", func(t *testing.T) {
		token := GenerateResetToken()
		if len(token) != 32 {
			t.Errorf("expected token length 32 but got %d", len(token))
		}
	})
	t.Run("contains only valid characters", func(t *testing.T) {
		token := GenerateResetToken()
		validChars := regexp.MustCompile(`^[a-zA-Z0-9]+$`)
		if !validChars.MatchString(token) {
			t.Errorf("expected only valid characters but got %q", token)
		}
	})
}
