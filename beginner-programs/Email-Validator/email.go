package Validator

import (
	"regexp"
)

var emailRegexp = regexp.MustCompile("(?i)" + // Case insensitive
	"^[a-z0-9!#$%&'*+/=?^_`{|}~.-]+" + // Validate local part
	"@" +
	"[a-z0-9-]+(\\.[a-z0-9-]+)*$") // Validate domain part

func IsValidEmail(email string) bool {
	if len(email) > 254 {
		return false
	}
	return emailRegexp.MatchString(email)
}
