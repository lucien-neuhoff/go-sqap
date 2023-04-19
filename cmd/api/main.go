package main

import (
	"fmt"
	"go-sqap/encryption"
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
	cfg := config.LoadConfig("/home/scaffus/dev/go/go-sqap/.env")
	logger := utils.NewLogger(&cfg)
	db := database.ConnectDatabase(cfg)

	err := encryption.Init()
	if err != nil {
		logger.Error("Error while generating RSA keypair: ", err)
	}

	keysRepo := repositories.NewKeysRepository(db, logger)
	userRepo := repositories.NewUserRepository(db, logger)
	sessionRepo := repositories.NewSessionRepository(db, logger)

	keysService := services.NewKeysService(userRepo, keysRepo, logger)
	sessionService := services.NewSessionService(userRepo, sessionRepo, logger)
	authService := services.NewAuthService(userRepo, sessionRepo, logger)

	keysHandler := handlers.NewKeysHandler(keysService, authService, logger)
	authHandler := handlers.NewAuthHandler(authService, sessionService, logger)

	router := router.CreateRouter(*authHandler, *keysHandler)

	logger.Infof("Starting server on %s:%s", cfg.APIHost, cfg.APIPort)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", cfg.APIPort), router); err != nil {
		log.Fatal(err)
	}
}
