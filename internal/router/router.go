package router

import (
	"go-sqap/internal/handlers"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CreateRouter(authHandler *handlers.AuthHandler, sessionHandler *handlers.SessionHandler) *gin.Engine {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Content-Length"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	api := router.Group("/api")
	{
		api.POST("/login", authHandler.LoginUser)
		api.POST("/register", authHandler.RegisterUser)
		api.POST("/session/create", sessionHandler.CreateSession)
		api.PUT("/session/validate", sessionHandler.ValidateSession)
	}

	return router
}
