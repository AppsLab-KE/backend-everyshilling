package handlers

import (
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/internal/core/usecase"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/internal/dto"
	"net/http"
)

type Handler struct {
	AuthUC usecase.AuthUseCase
}

func badRequest[T any](err string) dto.DefaultRes[T] {
	return dto.DefaultRes[T]{
		Message: "failed",
		Error:   err,
		Code:    http.StatusBadRequest,
	}
}

func okResponse[T any](data T) dto.DefaultRes[T] {
	return dto.DefaultRes[T]{
		Message: "Success",
		Error:   "",
		Code:    200,
		Data:    data,
	}
}

func NewHandler(authUC *usecase.AuthUseCase) ServerInterface {
	return &Handler{}
}
