package services

import (
	"go-sqap/internal/repositories"
	"go-sqap/internal/utils"
)

type UserService interface {
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
