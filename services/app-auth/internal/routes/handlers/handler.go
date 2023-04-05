package handlers

import (
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/internal/core/usecase"
)

type Handler struct {
	AuthUC usecase.AuthUseCase
}

func NewHandler(authUC *usecase.AuthUseCase) ServerInterface {
	return &Handler{}
}
