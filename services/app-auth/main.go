package main

import (
	"context"
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

	// TODO Initialise repositories & services

	// Initialise usecases
	authUC := usecase.NewAuthUsecase(nil, nil)

	// server config
	serverConfig := server.Config{
		AuthUsecase: authUC,
		Logger:      log,
	}
	handler := server.NewServer(serverConfig)

	port := ":" + os.Getenv("PORT")
	srv := &http.Server{
		Addr:    port,
		Handler: handler,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalln("Failed to start server.")
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
		log.Fatalln("Forced to shutdown")
	}
}
