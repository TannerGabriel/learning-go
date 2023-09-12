package Validator

import (
	"testing"
)

func TestIsValidEmail_d2603fb18f(t *testing.T) {
	t.Run("valid email", func(t *testing.T) {
		email := "test@example.com"
		if !IsValidEmail(email) {
			t.Errorf("Expected true for email '%s', got false", email)
		}
	})

	t.Run("invalid email", func(t *testing.T) {
		email := "test@example"
		if IsValidEmail(email) {
			t.Errorf("Expected false for email '%s', got true", email)
		}
	})

	t.Run("empty email", func(t *testing.T) {
		email := ""
		if IsValidEmail(email) {
			t.Errorf("Expected false for email '%s', got true", email)
		}
	})

	t.Run("long email", func(t *testing.T) {
		email := "test@example.com"
		for i := 0; i < 20; i++ {
			email += "long"
		}
		if IsValidEmail(email) {
			t.Errorf("Expected false for email '%s', got true", email)
		}
	})
}
