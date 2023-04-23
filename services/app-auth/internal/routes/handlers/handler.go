package handlers

import (
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/internal/core/adapters"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/internal/dto"
	"net/http"
)

type Handler struct {
	AuthUC adapters.AuthUseCase
}

func handleError[T any](err error) dto.DefaultRes[T] {
	return dto.DefaultRes[T]{
		Message: "failed",
		Error:   err.Error(),
		Code:    http.StatusBadRequest,
	}
}

func okResponse[T any](data T, message string) dto.DefaultRes[T] {
	return dto.DefaultRes[T]{
		Message: message,
		Error:   "",
		Code:    200,
		Data:    data,
	}
}

func NewHandler(authUC adapters.AuthUseCase) ServerInterface {
	return &Handler{
		AuthUC: authUC,
	}
}
