package utils

import "regexp"

// Password validation
var (
	// must have a length of at least 8
	passwordLen = 8

	// must be at least one lowercase letter
	lowercaseRegex = regexp.MustCompile(`[a-z]`)

	// must be at least one uppercase letter
	uppercaseRegex = regexp.MustCompile(`[A-Z]`)

	// must be at least one digit
	digitRegex = regexp.MustCompile(`\d`)

	// must be at least one special char
	specialCharRegex = regexp.MustCompile(`[!@#$%^&*(),.?"{}<>]`)
)

// ValidatePassword() returns true if a given password is valid
func ValidatePassword(password string) bool {
	if len(password) < passwordLen {
		return false
	}

	if !lowercaseRegex.MatchString(password) {
		return false
	}

	if !uppercaseRegex.MatchString(password) {
		return false
	}

	if !digitRegex.MatchString(password) {
		return false
	}

	if !specialCharRegex.MatchString(password) {
		return false
	}

	return true
}

// Username validation, only can include:
// lowercase or uppercase letters, digits, hyphen, underscore
var usernameRegex = regexp.MustCompile(`^[a-zA-Z0-9_-]{1,39}$`)

// ValidateUsername() returns true if a given username is valid
func ValidateUsername(username string) bool {
	return usernameRegex.MatchString(username)
}

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

func ValidateEmail(email string) bool {
	return emailRegex.MatchString(email)
}
