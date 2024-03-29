package main

import (
	"context"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/config"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/internal/core/repository"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/internal/core/service"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/internal/core/storage"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/internal/core/usecase"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/internal/routes/server"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	ServiceName = "app-auth"
)

func main() {
	// Initialise Logger
	log := logrus.New()
	// Dependency initialisation

	// register service

	// Load config
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Loading config failed", err)
	}

	// initilise storage
	dbStorage, err := storage.NewDbStorage(cfg.Database)
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}

	otpStorage, err := storage.NewOtpStorage(cfg.OTP)
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}

	cacheStorage, err := storage.NewCacheStorage(cfg.Redis)
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}
	// Repos
	authRepo := repository.NewAuthRepo(cacheStorage, dbStorage, otpStorage)
	// services
	authService := service.NewDefaultAuthService(cfg.Jwt, authRepo)

	// Initialise usecases
	authUC := usecase.NewAuthUsecase(authService, nil)

	// server config
	handler := server.NewServer(authUC, *cfg)

	serviceAddress := ":" + os.Getenv("PORT")
	srv := &http.Server{
		Addr:    serviceAddress,
		Handler: handler,
	}
	//run db migrations

	go func() {

		log.Println("Starting server on Address ", serviceAddress)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Error(err)
			os.Exit(1)
		}
	}()

	// Shutdown
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	log.Println("Shutting down server")
	if err := srv.Shutdown(ctx); err != nil {
		log.Error(err)
		os.Exit(1)
	}
}
