package services

import (
	"context"
	"go-sqap/encryption"
	"go-sqap/internal/models"
	"go-sqap/internal/repositories"
	"go-sqap/internal/utils"
)

type SessionService interface {
	CreateSession(publicKey string, user_id *string) (*models.Session, error)
}

type sessionService struct {
	userRepository    repositories.UserRepository
	sessionRepository repositories.SessionRepository
	logger            *utils.Logger
}

func NewSessionService(userRepo repositories.UserRepository, sessionRepository repositories.SessionRepository, logger *utils.Logger) SessionService {
	return &sessionService{
		userRepository:    userRepo,
		sessionRepository: sessionRepository,
		logger:            logger,
	}
}

func (s *sessionService) CreateSession(publicKey string, userId *string) (*models.Session, error) {
	var session models.Session

	session.Token = utils.GenerateToken(128)
	publicKeyStr, err := encryption.StringToPublicKey(publicKey)
	if err != nil {
		return nil, err
	}
	session.PublicKey = *publicKeyStr

	err = s.sessionRepository.SaveSession(context.Background(), &session)
	if err != nil {
		return nil, err
	}

	return &session, nil
}
