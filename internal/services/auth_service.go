package services

import (
	"go-sqap/internal/models"
	"go-sqap/internal/repositories"
	"go-sqap/internal/utils"

	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
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

func (s *authService) Authenticate(loginRequest *models.LoginRequest) (*models.Session, error) {
	user, err := s.userRepository.FindByEmail(loginRequest.Email)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password))
	if err != nil {
		return nil, err
	}

	session, err := s.sessionRepository.Create(user.ID)
	if err != nil {
		return nil, err
	}

	return session, err
}
