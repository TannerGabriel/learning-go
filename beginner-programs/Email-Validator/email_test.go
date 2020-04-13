package Validator

import (
	"testing"
)

var validEmails = []string{
	"a@test.com",
	"postmaster@test.com",
	"president@kremlin.gov.ru",
	"test@test.co.uk",
}

var invalidEmails = []string{
	"",
	"test",
	"test.com",
	".com",
	"адрес@пример.рф",
	" space_before@test.com",
	"space between@test.com",
	"\nnewlinebefore@test.com",
	"newline\nbetween@test.com",
	"test@test.com.",
	"asyouallcanseethisemailaddressexceedsthemaximumnumberofcharactersallowedtobeintheemailaddresswhichisnomorethatn254accordingtovariousrfcokaycanistopnowornotyetnoineedmorecharacterstoadd@i.really.cannot.thinkof.what.else.to.put.into.this.invalid.address.net",
}

func TestIsValidEmail(t *testing.T) {
	for i, v := range validEmails {
		if !IsValidEmail(v) {
			t.Errorf("%d: didn't accept valid email: %s", i, v)
		}
	}
	for i, v := range invalidEmails {
		if IsValidEmail(v) {
			t.Errorf("%d: accepted invalid email: %s", i, v)
		}
	}
}
