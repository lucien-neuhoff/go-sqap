package handlers

import (
	"errors"
	"go-sqap/internal/models"
	"regexp"
)

func validCreateUserRequest(req *models.CreateUserRequest) error {
	// Check that the email is valid
	if !isValidEmail(req.Email) {
		return errors.New("invalid email address")
	}

	// Check that the password meets complexity requirements
	if !isValidPassword(req.Password) {
		return errors.New("password must be at least 8 characters long and contain at least one uppercase letter, one lowercase letter, and one digit")
	}

	return nil
}

func isValidEmail(email string) bool {
	// A simple regex to validate the email format
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}

func isValidPassword(password string) bool {
	// Check length
	if len(password) < 8 {
		return false
	}

	// Check for at least one uppercase letter, one lowercase letter, and one digit
	hasUpper := false
	hasLower := false
	hasDigit := false

	for _, char := range password {
		if char >= 'A' && char <= 'Z' {
			hasUpper = true
		} else if char >= 'a' && char <= 'z' {
			hasLower = true
		} else if char >= '0' && char <= '9' {
			hasDigit = true
		}
	}

	return hasUpper && hasLower && hasDigit
}
