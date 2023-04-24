package handlers

import (
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/internal/core/adapters"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/internal/core/service"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/internal/dto"
	"net/http"
)

const (
	BearerScopes = "Bearer.Scopes"
)

type Handler struct {
	AuthUC adapters.AuthUseCase
}

func handleError[T any](err error) dto.DefaultRes[T] {
	var responseCode int

	switch err {
	case service.ErrIncorrectPassword:
		responseCode = http.StatusUnauthorized
	case service.ErrValidationError:
		responseCode = http.StatusBadRequest
	case service.ErrCacheFetch:
		responseCode = http.StatusInternalServerError
	case service.ErrCacheSave:
		responseCode = http.StatusInternalServerError
	case service.ErrRequestValidation:
		responseCode = http.StatusBadRequest
	case service.ErrHashGeneration:
		responseCode = http.StatusInternalServerError
	case service.ErrTokenGeneration:
		responseCode = http.StatusInternalServerError
	case service.ErrUserNotFound:
		responseCode = http.StatusNotFound
	case service.ErrDatabaseWrite:
		responseCode = http.StatusInternalServerError
	case service.ErrOTPNotInitialied:
		responseCode = http.StatusUnauthorized
	case service.ErrUserExists:
		responseCode = http.StatusConflict
	default:
		responseCode = http.StatusInternalServerError

	}

	responseMessage := http.StatusText(responseCode)

	return dto.DefaultRes[T]{
		Message: "failed: " + responseMessage,
		Error:   err.Error(),
		Code:    responseCode,
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
