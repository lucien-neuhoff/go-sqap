package services

import (
	"context"
	"errors"
	"go-sqap/internal/models"
	"go-sqap/internal/repositories"
	"go-sqap/internal/utils"

	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Authenticate(ctx context.Context, loginRequest *models.LoginRequest) (*models.Session, error)
	CreateSession(ctx context.Context, user_id string) (*models.Session, error)
}

type authService struct {
	userRepository    repositories.UserRepository
	sessionRepository repositories.SessionRepository
	logger            *utils.Logger
}

func NewAuthService(userRepo repositories.UserRepository, sessionRepository repositories.SessionRepository, logger *utils.Logger) AuthService {
	return &authService{
		userRepository:    userRepo,
		sessionRepository: sessionRepository,
		logger:            logger,
	}
}

func (s *authService) Authenticate(ctx context.Context, loginRequest *models.LoginRequest) (*models.Session, error) {
	user, err := s.userRepository.GetUserByEmail(context.Background(), loginRequest.Email)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password)); err != nil {
		return nil, errors.New("auth/invalid-credentials")
	}

	session, err := s.sessionRepository.CreateSession(ctx, user.UUID)
	if err != nil {
		return nil, err
	}

	return session, err
}

func (s *authService) CreateSession(ctx context.Context, user_id string) (*models.Session, error) {
	session, err := s.sessionRepository.CreateSession(ctx, user_id)
	if err != nil {
		return nil, err
	}

	return session, nil
}
