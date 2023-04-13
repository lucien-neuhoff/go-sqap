package main

import (
	"fmt"
	"go-sqap/internal/config"
	"go-sqap/internal/database"
	"go-sqap/internal/handlers"
	"go-sqap/internal/repositories"
	"go-sqap/internal/services"
	"go-sqap/internal/utils"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig("../.env")
	logger := utils.NewLogger()
	db := database.ConnectDatabase(cfg)

	userRepo := repositories.NewUserRepository(db, logger)
	sessionRepo := repositories.NewSessionRepository(db, logger)

	userService := services.NewUserService(userRepo, logger)
	authService := services.NewAuthService(userRepo, sessionRepo, logger)

	userHandler := handlers.NewUserHandler(userService, logger)
	authHandler := handlers.NewAuthHandler(authService, logger)

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
		api.POST("/login", authHandler.Login)
		api.POST("/register", userHandler.CreateUser)
		api.GET("/users", userHandler.GetUsers)
	}

	logger.Info(fmt.Sprintf("Starting server on %s:%s", cfg.APIHost, cfg.APIPort))
	if err := http.ListenAndServe(fmt.Sprintf(":%s", cfg.APIPort), router); err != nil {
		log.Fatal(err)
	}
}
