package main

import (
	"fmt"
	"go-sqap/internal/config"
	"go-sqap/internal/database"
	"go-sqap/internal/handlers"
	"go-sqap/internal/repositories"
	"go-sqap/internal/router"
	"go-sqap/internal/services"
	"go-sqap/internal/utils"
	"log"
	"net/http"
)

func main() {
	cfg := config.LoadConfig("D:/go/go-sqap/.env")
	logger := utils.NewLogger(&cfg)
	db := database.ConnectDatabase(cfg)

	userRepo := repositories.NewUserRepository(db, logger)
	sessionRepo := repositories.NewSessionRepository(db, logger)

	userService := services.NewUserService(userRepo, logger)
	authService := services.NewAuthService(userRepo, sessionRepo, logger)

	userHandler := handlers.NewUserHandler(userService, logger)
	authHandler := handlers.NewAuthHandler(authService, logger)

	router := router.CreateRouter(*authHandler, *userHandler)

	logger.Infof("Starting server on %s:%s", cfg.APIHost, cfg.APIPort)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", cfg.APIPort), router); err != nil {
		log.Fatal(err)
	}
}
