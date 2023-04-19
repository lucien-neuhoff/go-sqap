package services

import (
	"context"
	"errors"
	"fmt"
	"go-sqap/internal/models"
	"go-sqap/internal/repositories"
	"go-sqap/internal/utils"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Authenticate(loginRequest *models.LoginRequest) (*models.User, error)
	RegisterUser(user *models.User) error
	GetUserByEmail(email string) (*models.User, error)
}

type authService struct {
	userRepository repositories.UserRepository
	logger         *utils.Logger
}

func NewAuthService(userRepo repositories.UserRepository, sessionRepository repositories.SessionRepository, logger *utils.Logger) AuthService {
	return &authService{
		userRepository: userRepo,
		logger:         logger,
	}
}

func (s *authService) Authenticate(loginRequest *models.LoginRequest) (*models.User, error) {
	user, err := s.userRepository.GetUserByEmail(context.Background(), loginRequest.Email)
	if err != nil {
		s.logger.Debug("Error while getting user")
		return nil, err
	}

	if user == nil {
		s.logger.Debug("User not found")
		return nil, errors.New("user not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password)); err != nil {
		s.logger.Debug("Error during password comparison")
		return nil, errors.New("auth/invalid-credentials")
	}

	return user, nil
}

func (s *authService) RegisterUser(user *models.User) error {
	// Check if the user already exists
	existingUser, err := s.userRepository.GetUserByEmail(context.Background(), user.Email)
	if existingUser != nil {
		return errors.New("auth/user-already-exists")
	}

	if err != nil {
		return fmt.Errorf("failed to get user by email: %v", err)
	}

	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		return fmt.Errorf("error while hashing password: %v", err)
	}

	now := time.Now().UTC()

	user.Password = string(password)
	user.CreatedAt.Time = now
	user.UpdatedAt.Time = now
	user.UUID = uuid.NewString()

	s.logger.Infof("Creating user with email '%s' & uuid '%s'", user.Email, user.UUID)
	// Create the user
	if err := s.userRepository.CreateUser(context.Background(), user); err != nil {
		return fmt.Errorf("failed to create user: %v", err)
	}

	return nil
}

func (s *authService) GetUserByEmail(email string) (*models.User, error) {
	return s.userRepository.GetUserByEmail(context.Background(), email)
}
