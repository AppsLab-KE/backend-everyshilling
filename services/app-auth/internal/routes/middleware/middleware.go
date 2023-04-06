package middleware

import (
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/config"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/internal/core/usecase"
)

type Manager struct {
	authUC usecase.AuthUseCase
	config config.Config
}

func NewManager(config config.Config, authUC usecase.AuthUseCase) *Manager {
	return &Manager{
		authUC,
		config,
	}
}
