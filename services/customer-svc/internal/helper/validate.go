package helper

import (
	"regexp"
)

func ValidateEmail(email string) bool {
	regex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return regex.MatchString(email)
}

func ValidatePassword(password string) bool {
	return len(password) >= 6
}

func ValidateNotEmptyStr(value string) bool {
	return len(value) > 0
}
