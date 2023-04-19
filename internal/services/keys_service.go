package services

import (
	"context"
	"go-sqap/internal/models"
	"go-sqap/internal/repositories"
	"go-sqap/internal/utils"
)

type KeysService interface {
	SaveUserPublicKey(ctx context.Context, publicKey models.PublicKey) error
	GetUserPublicKey(ctx context.Context, user *models.User) (*string, error)
}

type keysService struct {
	userRepository repositories.UserRepository
	keysRepository repositories.KeysRepository
	logger         *utils.Logger
}

func NewKeysService(userRepo repositories.UserRepository, keysRepo repositories.KeysRepository, logger *utils.Logger) KeysService {
	return &keysService{
		userRepository: userRepo,
		keysRepository: keysRepo,
		logger:         logger,
	}
}

func (s *keysService) SaveUserPublicKey(ctx context.Context, publicKey models.PublicKey) error {
	err := s.keysRepository.SaveUserPublicKey(ctx, publicKey)
	if err != nil {
		return err
	}

	return nil
}

func (s *keysService) GetUserPublicKey(ctx context.Context, user *models.User) (*string, error) {
	publicKey, err := s.keysRepository.GetUserPublicKey(ctx, user)
	if err != nil {
		return nil, err
	}

	return publicKey, nil
}
