package Validator

import (
	"regexp"
	"testing"
)

var emailRegexp = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

func IsValidEmail(email string) bool {
	if len(email) > 254 {
		return false
	}
	return emailRegexp.MatchString(email)
}

func TestIsValidEmail_d2603fb18f(t *testing.T) {
	t.Run("valid email", func(t *testing.T) {
		email := "test@example.com"
		if !IsValidEmail(email) {
			t.Error("Expected true, got false")
		}
	})

	t.Run("invalid email", func(t *testing.T) {
		email := "test@example"
		if IsValidEmail(email) {
			t.Error("Expected false, got true")
		}
	})

	t.Run("empty email", func(t *testing.T) {
		email := ""
		if IsValidEmail(email) {
			t.Error("Expected false, got true")
		}
	})

	t.Run("long email", func(t *testing.T) {
		email := "test@example.com"
		for i := 0; i < 20; i++ {
			email += "long"
		}
		if IsValidEmail(email) {
			t.Error("Expected false, got true")
		}
	})
}
