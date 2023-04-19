package handlers

import (
	"context"
	"go-sqap/encryption"
	"go-sqap/internal/models"
	"go-sqap/internal/services"
	"go-sqap/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type KeysHandler struct {
	keysService services.KeysService
	authService services.AuthService
	logger      *utils.Logger
}

func NewKeysHandler(keysService services.KeysService, authService services.AuthService, logger *utils.Logger) *KeysHandler {
	return &KeysHandler{
		keysService: keysService,
		authService: authService,
		logger:      logger,
	}
}

func (h *KeysHandler) ExchangeKeys(c *gin.Context) {
	var req models.PublicKeyRequest

	if err := c.BindJSON(&req); err == nil {
		h.logger.Error("failed to bind save public key request: ", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	user, err := h.authService.GetUserByEmail(req.Email)
	if err != nil {
		h.logger.Error("error while retrieving user by email: ", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var publicKey_req models.PublicKey
	publicKey_req.UserID = user.UUID
	publicKey_req.Key = req.Key

	err = h.keysService.SaveUserPublicKey(context.Background(), publicKey_req)
	if err != nil {
		h.logger.Error("error while saving user public key: ", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "error while saving public key"})
		return
	}

	serverPublicKeyStr, err := encryption.GetServerPublicKeyString()
	if err != nil {
		h.logger.Error("error while parsing public key to string: ", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "error while getting server's public key"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"server_public_key": serverPublicKeyStr})
}
