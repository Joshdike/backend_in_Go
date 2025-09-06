package utils

import (
	"testing"
)

func TestValidateUsername(t *testing.T) {
	t.Run("valid username", func(t *testing.T) {
		err := ValidateUsername("username")
		assertNoError(t, err)
	})

	t.Run("invalid username", func(t *testing.T) {
		err := ValidateUsername("")
		assertError(t, err, ErrInvalidUsername)
	})
}

func TestValidateEmail(t *testing.T) {
	t.Run("invalid email", func(t *testing.T) {
		err := ValidateEmail("invalid_email")
		assertError(t, err, ErrInvalidEmail)
	})

	t.Run("valid email", func(t *testing.T) {
		err := ValidateEmail("HjyXq@example.com")
		assertNoError(t, err)
	})
}

func TestValidatePassword(t *testing.T) {
	t.Run("valid password", func(t *testing.T) {
		err := ValidatePassword("password")
		assertNoError(t, err)
	})

	t.Run("invalid password", func(t *testing.T) {
		err := ValidatePassword("short")
		assertError(t, err, ErrInvalidPassword)
	})
}

func assertNoError(t *testing.T, got error) {
	t.Helper()
	if got != nil {
		t.Fatalf("expected no error but got %q", got)
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()
	if got == nil {
		t.Errorf("expected error but got nil")
	}

	if got != want {
		t.Errorf("expected error %q but got %q", want, got)
	}
}
