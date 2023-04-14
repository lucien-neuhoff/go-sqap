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

	return nil
}

func isValidEmail(email string) bool {
	// A simple regex to validate the email format
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}
