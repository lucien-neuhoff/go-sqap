package services

import (
	"context"
	"go-sqap/internal/models"
	"go-sqap/internal/repositories"
	"go-sqap/internal/utils"
)

type SessionService interface {
	CreateSession(user_id *string, publicKey *string) (*models.Session, error)
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

func (s *sessionService) CreateSession(userId *string, publicKey *string) (*models.Session, error) {
	var session models.Session

	session.Token = utils.GenerateToken(255)
	session.PublicKey = publicKey

	err := s.sessionRepository.SaveSession(context.Background(), &session)
	if err != nil {
		s.logger.Error("Error while creating session: ", err)
		return nil, err
	}

	return &session, nil
}
