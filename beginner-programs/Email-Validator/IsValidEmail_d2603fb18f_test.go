package Validator

import (
	"log"
	"testing"
)

func TestIsValidEmail_d2603fb18f(t *testing.T) {
	tests := []struct {
		name  string
		email string
		want  bool
	}{
		{"valid email", "test@example.com", true},
		{"invalid email", "test@example", false},
		{"empty email", "", false},
		{"long email", "test@example.com" + repeat("long", 100), false},
		{"invalid email with no domain", "test", false},
		{"invalid email with no '@'", "test.com", false},
		{"invalid email with no name", "@.com", false},
		{"invalid email with non-ascii characters", "адрес@пример.рф", false},
		{"invalid email with space before", " space_before@test.com", false},
		{"invalid email with space between", "space between@test.com", false},
		{"invalid email with newline before", "\nnewlinebefore@test.com", false},
		{"invalid email with newline between", "newline\nbetween@test.com", false},
		{"invalid email with dot at the end", "test@test.com.", false},
		{"invalid email with excessive length", "asyouallcanseethisemailaddressexceedsthemaximumnumberofcharactersallowedtobeintheemailaddresswhichisnomorethatn254accordingtovariousrfcokaycanistopnowornotyetnoineedmorecharacterstoadd@i.really.cannot.thinkof.what.else.to.put.into.this.invalid.address.net", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidEmail(tt.email); got != tt.want {
				t.Errorf("IsValidEmail(%s) = %v, want %v", tt.email, got, tt.want)
			} else {
				log.Printf("Success: TestIsValidEmail_d2603fb18f with case '%s'\n", tt.name)
			}
		})
	}
}

func repeat(s string, n int) string {
	var result string
	for i := 0; i < n; i++ {
		result += s
	}
	return result
}
