package middleware

import (
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/config"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/internal/core/adapters"
)

type Manager struct {
	authUC adapters.AuthUseCase
	config config.Config
}

func NewManager(config config.Config, authUC adapters.AuthUseCase) *Manager {
	return &Manager{
		authUC,
		config,
	}
}
