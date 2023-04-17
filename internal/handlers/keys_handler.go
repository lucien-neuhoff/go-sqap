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
	userService services.UserService
	logger      *utils.Logger
}

func NewKeysHandler(keysService services.KeysService, userSerivce services.UserService, logger *utils.Logger) *KeysHandler {
	return &KeysHandler{
		keysService: keysService,
		userService: userSerivce,
		logger:      logger,
	}
}

func (h *KeysHandler) ExchangeKeys(c *gin.Context) {
	var req models.PublicKeyRequest

	if err := c.BindJSON(&req); err != nil {
		h.logger.Error("failed to bind save public key request: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.userService.GetUserByEmail(req.Email)
	if err != nil {
		h.logger.Error("error while retrieving user by email: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var publicKey_req models.PublicKey
	publicKey_req.UserID = user.UUID
	publicKey_req.Key = req.Key

	err = h.keysService.SaveUserPublicKey(context.Background(), publicKey_req)
	if err != nil {
		h.logger.Error("error while saving user public key: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	serverPublicKey := h.keysService.GetServerPublicKey()
	serverPublicKeyString, err := encryption.PublicKeyToString(&serverPublicKey)
	if err != nil {
		h.logger.Error("error while parsing public key to string: ", err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"server_public_key": serverPublicKeyString})
}
