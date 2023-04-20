package handlers

import (
	"go-sqap/encryption"
	"go-sqap/internal/models"
	"go-sqap/internal/services"
	"go-sqap/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService    services.AuthService
	sessionService services.SessionService
	logger         *utils.Logger
}

func NewAuthHandler(authService services.AuthService, sessionService services.SessionService, logger *utils.Logger) *AuthHandler {
	return &AuthHandler{
		authService:    authService,
		sessionService: sessionService,
		logger:         logger,
	}
}

func (h *AuthHandler) LoginUser(c *gin.Context) {
	var loginRequest models.LoginRequest

	if err := c.BindJSON(&loginRequest); err != nil {
		h.logger.Errorf("Error while binding login request: %v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "auth/invalid-request"})
		return
	}

	user, err := h.authService.Authenticate(&loginRequest)
	if err != nil {
		h.logger.Errorf("Error while authenticating user: %v", err)
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "auth/user-not-found"})
		return
	}

	h.logger.Debug("UUID: ", user.UUID)
	session, err := h.sessionService.CreateSession(user.PublicKey, &user.UUID)
	if err != nil {
		h.logger.Errorf("Error while creating session: %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "auth/internal-server-error"})
		return
	}

	encryptedToken, err := encryption.Encrypt([]byte(session.Token), &session.PublicKey)
	if err != nil {
		h.logger.Error("Error while encrypting token: ", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "auth/internal-server-error"})
		return
	}

	encryptedUser, err := encryption.EncryptUser(*user, &session.PublicKey)
	if err != nil {
		h.logger.Error("Error while encrypting user: ", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "auth/internal-server-error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": encryptedUser, "token": encryptedToken})
}

func (h *AuthHandler) RegisterUser(c *gin.Context) {
	var req models.CreateUserRequest

	if err := c.BindJSON(&req); err != nil {
		h.logger.Error("failed to bind create user request: ")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validCreateUserRequest(&req); err != nil {
		h.logger.Error("could not valided create user request: ")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user_req models.User
	user_req.Email = req.Email
	user_req.Password = req.Password

	err := h.authService.RegisterUser(&user_req)
	if err != nil {
		h.logger.Errorf("failed: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.authService.GetUserByEmail(req.Email)
	if err != nil {
		h.logger.Error("error while retrieving newly create user: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusCreated, user)
}
