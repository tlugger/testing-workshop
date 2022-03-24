package pwdvalidator

import (
	"errors"
	"regexp"
)

// https://en.wikipedia.org/wiki/List_of_the_most_common_passwords
var commonPasswords = map[string]bool{
	"123456":      true,
	"123456789":   true,
	"12345":       true,
	"qwerty":      true,
	"password":    true,
	"12345678":    true,
	"111111":      true,
	"123123":      true,
	"1234567890":  true,
	"1234567":     true,
	"qwerty123":   true,
	"000000":      true,
	"1q2w3e":      true,
	"aa12345678":  true,
	"abc123":      true,
	"password1":   true,
	"1234":        true,
	"qwertyuiop":  true,
	"123321":      true,
	"password123": true,
}

var capitalRegex = regexp.MustCompile(`\p{Lu}`)
var lowerRegex = regexp.MustCompile(`\p{Ll}`)
var specialRegex = regexp.MustCompile(`^.*[!@#$%^&].+$`)

func ValidatePassword(password string) error {
	if password == "" {
		return errors.New("MUST_PROVIDE_PASSWORD")
	}
	if len(password) < 8 {
		return errors.New("PWD_TOO_SHOT")
	}
	if len(password) > 144 {
		return errors.New("PWD_TOO_LONG")
	}
	if commonPasswords[password] {
		return errors.New("PWD_NOT_SECURE")
	}
	if !capitalRegex.MatchString(password) {
		return errors.New("PWD_MUST_INCLUDE_UPPERCASE_LETTER")
	}
	if !lowerRegex.MatchString(password) {
		return errors.New("PWD_MUST_INCLUDE_LOWERCASE_LETTER")
	}
	if !specialRegex.MatchString(password) {
		return errors.New("PWD_MUST_INCLUDE_SPECIAL_CHARACTER")
	}
	return nil
}
