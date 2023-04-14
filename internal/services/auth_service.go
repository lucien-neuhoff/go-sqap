package services

import (
	"context"
	"errors"
	"go-sqap/internal/models"
	"go-sqap/internal/repositories"
	"go-sqap/internal/utils"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Authenticate(ctx context.Context, loginRequest *models.LoginRequest) (*models.User, error)
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

func (s *authService) Authenticate(ctx context.Context, loginRequest *models.LoginRequest) (*models.User, error) {
	user, err := s.userRepository.GetUserByEmail(context.Background(), loginRequest.Email)
	if err != nil {
		s.logger.Error("Error while getting user")
		return nil, err
	}

	if user == nil {
		s.logger.Error("User not found")
		return nil, errors.New("auth/user-not-found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password)); err != nil {
		s.logger.Error("Error during password comparison")
		return nil, errors.New("auth/invalid-credentials")
	}

	return user, nil
}

func (s *authService) CreateSession(ctx context.Context, user_id string) (*models.Session, error) {
	var session *models.Session

	now := time.Now().UTC()

	session.CreatedAt = now
	session.UpdatedAt = now
	session.UserID = user_id
	session.UUID = uuid.New().String()
	session.Token = utils.GenerateToken(255)

	err := s.sessionRepository.CreateSession(ctx, session)
	if err != nil {
		return nil, err
	}

	if session == nil {
		return nil, errors.New("failed to create session")
	}

	return session, nil
}
