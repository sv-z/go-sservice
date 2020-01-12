package validator

import (
	valid "github.com/go-playground/validator/v10"
	"regexp"
	"unicode"
)

func IsPasswordValid(fl valid.FieldLevel) bool {
	pw := fl.Field().String()
	// https://stackoverflow.com/questions/19605150/regex-for-password-must-contain-at-least-eight-characters-at-least-one-number-a
	regex := regexp.MustCompile("^[A-Za-z\\d+@$!%*?&]+$")
	if !regex.MatchString(pw) {
		return false
	}

	var num, lower, upper, spec bool
	for _, r := range pw {
		switch {
		case unicode.IsDigit(r):
			num = true
		case unicode.IsUpper(r):
			upper = true
		case unicode.IsLower(r):
			lower = true
		case unicode.IsSymbol(r), unicode.IsPunct(r):
			spec = true
		}
	}
	if num && lower && upper && spec {
		return true
	}

	return false
}
