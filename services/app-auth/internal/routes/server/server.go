package server

import (
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/internal/core/adapters"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/internal/routes/handlers"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

const BaseUrl = "/api"

type Config struct {
	AuthUsecase adapters.AuthUseCase
	Logger      *logrus.Logger
}

func NewServer(cfg Config) *gin.Engine {
	r := gin.Default()
	options := handlers.GinServerOptions{
		BaseURL: BaseUrl,
	}

	// Map handlers
	h := handlers.NewHandler(cfg.AuthUsecase)
	handlers.RegisterHandlersWithOptions(r, h, options)

	return r
}
