package handlers

import (
	"go-sqap/internal/models"
	"go-sqap/internal/services"
	"go-sqap/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService services.UserService
	logger      *utils.Logger
}

func NewUserHandler(userService services.UserService, logger *utils.Logger) *UserHandler {
	return &UserHandler{
		userService: userService,
		logger:      logger,
	}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var req models.CreateUserRequest

	if err := c.BindJSON(&req); err != nil {
		h.logger.Errorf("failed to bind create user request: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validCreateUserRequest(&req); err != nil {
		h.logger.Errorf("could not valided create user request: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user_req models.User
	user_req.Email = req.Email
	user_req.Password = req.Password

	err := h.userService.CreateUser(&user_req)
	if err != nil {
		h.logger.Errorf("failed: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.userService.GetUserByEmail(req.Email)
	if err != nil {
		h.logger.Errorf("error while retrieving newly create user: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusCreated, user)
}

func (h *UserHandler) GetUsers(c *gin.Context) {

}
