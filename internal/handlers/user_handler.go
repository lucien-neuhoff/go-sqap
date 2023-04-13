package handlers

import (
	"go-sqap/internal/services"
	"go-sqap/internal/utils"
)

type UserHandler struct {
	userService *services.UserService
	logger      *utils.Logger
}

func NewUserHandler(userService *services.UserService, logger *utils.Logger) *UserHandler {
	return &UserHandler{
		userService: userService,
		logger:      logger,
	}
}
