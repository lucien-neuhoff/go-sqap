package handlers

import (
	"context"
	"go-sqap/internal/models"
	"go-sqap/internal/services"
	"go-sqap/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService services.AuthService
	logger      *utils.Logger
}

func NewAuthHandler(authService services.AuthService, logger *utils.Logger) *AuthHandler {
	return &AuthHandler{
		authService: authService,
		logger:      logger,
	}
}

func (h *AuthHandler) Login(c *gin.Context) {
	var loginRequest models.LoginRequest

	if err := c.BindJSON(&loginRequest); err != nil {
		h.logger.Errorf("failed to bind login request: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "auth/invalid-request"})
		return
	}

	user, err := h.authService.Authenticate(context.Background(), &loginRequest)
	if err != nil {
		h.logger.Errorf("failed to authenticate user: %v", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "auth/user-not-found"})
		return
	}

	h.logger.Debug("UUID: ", user.UUID)
	session, err := h.authService.CreateSession(context.Background(), user.UUID)
	if err != nil {
		h.logger.Errorf("failed to create session: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "auth/internal-server-error"})
		return
	}

	// TODO: Send the session token with asymetric encryption
	c.JSON(http.StatusOK, gin.H{"user": gin.H{"uuid": user.UUID, "email": user.Email}, "token": session.Token})
}
