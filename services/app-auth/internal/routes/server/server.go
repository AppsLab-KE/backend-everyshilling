package server

import (
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/internal/core/adapters"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/internal/routes/handlers"
	"github.com/gin-gonic/gin"
)

const BaseUrl = "/api"

type Config struct {
	AuthUsecase adapters.AuthUseCase
}

func NewServer(authUseCase adapters.AuthUseCase) *gin.Engine {
	r := gin.Default()
	options := handlers.GinServerOptions{
		BaseURL: BaseUrl,
	}

	// Serve swagger spec

	// Map handlers
	h := handlers.NewHandler(authUseCase)
	handlers.RegisterHandlersWithOptions(r, h, options)

	return r
}
