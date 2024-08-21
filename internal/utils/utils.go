package utils

import (
	"regexp"
	"strings"
)

// ValidateEmail validates an email address format
func ValidateEmail(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}

// ToLowerCase converts a string to lowercase
func ToLowerCase(s string) string {
	return strings.ToLower(s)
}
