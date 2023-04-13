package services

import (
	"context"
	"errors"
	"fmt"
	"go-sqap/internal/models"
	"go-sqap/internal/repositories"
	"go-sqap/internal/utils"
)

type UserService interface {
	CreateUser(user *models.User) error
	GetUserByEmail(email string) (*models.User, error)
}

type userSerivce struct {
	userRepository repositories.UserRepository
	logger         *utils.Logger
}

func NewUserService(userRepository repositories.UserRepository, logger *utils.Logger) UserService {
	return &userSerivce{
		userRepository: userRepository,
		logger:         logger,
	}
}

func (s *userSerivce) CreateUser(user *models.User) error {
	// Check if the user already exists
	existingUser, err := s.userRepository.GetUserByEmail(context.Background(), user.Email)
	if existingUser != nil {
		return errors.New("auth/user-already-exists")
	}

	if err != nil {
		return fmt.Errorf("failed to create user: %v", err)
	}

	// Create the user
	if err := s.userRepository.CreateUser(context.Background(), user); err != nil {
		return fmt.Errorf("failed to create user: %v", err)
	}

	return nil
}

func (s *userSerivce) GetUserByEmail(email string) (*models.User, error) {
	return s.userRepository.GetUserByEmail(context.Background(), email)
}
