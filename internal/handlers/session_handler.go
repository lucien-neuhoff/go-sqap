package handlers

import (
	"go-sqap/encryption"
	"go-sqap/internal/models"
	"go-sqap/internal/services"
	"go-sqap/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SessionHandler struct {
	sessionService services.SessionService
	logger         *utils.Logger
}

func NewSessionHandler(sessionService services.SessionService, logger *utils.Logger) *SessionHandler {
	return &SessionHandler{
		sessionService: sessionService,
		logger:         logger,
	}
}

func (h *SessionHandler) CreateSession(c *gin.Context) {
	var createSessionRequest models.CreateSessionRequest

	err := c.BindJSON(&createSessionRequest)
	if err != nil {
		h.logger.Error("Error while parsing public key: ", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "session/missing-key"})
		return
	}

	session, err := h.sessionService.CreateSession(createSessionRequest.UserID, createSessionRequest.PublicKey)
	if err != nil {
		h.logger.Error("Error while creating a new session: ", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	encryptedSessionToken, err := encryption.Encrypt([]byte(session.Token), session.PublicKey)
	if err != nil {
		h.logger.Error("Error while encrypting session token: ", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "session/token-encryption"})
	}

	serverPublicKeyString, err := encryption.GetServerPublicKeyString()
	if err != nil {
		h.logger.Error("Error while getting server public key: ", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "session/token-encryption"})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{"server_public_key": serverPublicKeyString, "token": encryptedSessionToken})
}
