package server

import (
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/config"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/internal/core/adapters"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/internal/routes/handlers"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/internal/routes/middleware"
	"github.com/gin-gonic/gin"
)

const BaseUrl = "/api"

type Config struct {
	AuthUsecase     adapters.AuthUseCase
	ExchangeUsecase adapters.ExchangeStorageUsecase
}

func NewServer(authUseCase adapters.AuthUseCase, exchangeUseCase adapters.ExchangeStorageUsecase, cfg config.Config) *gin.Engine {
	r := gin.Default()

	middlewareManager := middleware.NewManager(cfg, authUseCase)

	middlewares := []handlers.MiddlewareFunc{
		middlewareManager.Auth,
		//middlewareManager.RateLimiter,
		//middlewareManager.Log,
	}

	options := handlers.GinServerOptions{
		BaseURL:     BaseUrl,
		Middlewares: middlewares,
	}

	// Serve swagger spec

	// Map handlers
	//h := handlers.NewHandler(authUseCase,exchangeUseCase)
	h := handlers.NewHandler(authUseCase)

	handlers.RegisterHandlersWithOptions(r, h, options)

	return r
}
