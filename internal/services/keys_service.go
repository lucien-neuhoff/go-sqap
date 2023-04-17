package services

import (
	"context"
	"crypto/rsa"
	"go-sqap/internal/models"
	"go-sqap/internal/repositories"
	"go-sqap/internal/utils"
)

type KeysService interface {
	SaveUserPublicKey(ctx context.Context, publicKey models.PublicKey) error
	GetUserPublicKey(ctx context.Context, user *models.User) (*string, error)
	GetServerPublicKey() rsa.PublicKey
}

type keysService struct {
	userRepository  repositories.UserRepository
	keysRepository  repositories.KeysRepository
	logger          *utils.Logger
	serverPublicKey *rsa.PublicKey
}

func NewKeysService(userRepo repositories.UserRepository, keysRepo repositories.KeysRepository, logger *utils.Logger, serverPublicKey *rsa.PublicKey) KeysService {
	return &keysService{
		userRepository:  userRepo,
		keysRepository:  keysRepo,
		logger:          logger,
		serverPublicKey: serverPublicKey,
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

func (s *keysService) GetServerPublicKey() rsa.PublicKey {
	return *s.serverPublicKey
}
