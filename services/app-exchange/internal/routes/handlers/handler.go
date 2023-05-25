package handlers

import (
	"github.com/AppsLab-KE/backend-everyshilling/services/app-exchange/internal/core/adapters"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-exchange/internal/core/entity"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-exchange/internal/core/service"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-exchange/internal/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	BearerScopes = "Bearer.Scopes"
)

type Handler struct {
	AuthUC adapters.Currency
}

func (h Handler) GetAccountOverview(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (h Handler) PostBuyCurrency(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (h Handler) GetMarketplaceOffers(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (h Handler) PostTopUpAccount(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (h Handler) GetUserById(c *gin.Context, userId string) {
	//TODO implement me
	panic("implement me")
}

func handleError[T any](err error) dto.GeneralResponse[T] {
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
	case service.ErrTokenInvalid:
		responseCode = http.StatusUnauthorized
	case service.ErrVerificationOnWrongPhone:
		responseCode = http.StatusForbidden
	case service.ErrUserLoggedOut:
		responseCode = http.StatusUnauthorized
	case service.ErrUserNotFoundReset:
		responseCode = http.StatusOK
	default:
		responseCode = http.StatusInternalServerError

	}

	// custom error types
	switch err.(type) {
	case entity.ValidationError:
		responseCode = http.StatusBadRequest

	}

	responseMessage := http.StatusText(responseCode)

	return dto.GeneralResponse[T]{
		Message: "failed: " + responseMessage,
		Error:   err.Error(),
		Code:    responseCode,
	}
}

func okResponse[T any](data T, message string) dto.GeneralResponse[T] {
	return dto.GeneralResponse[T]{
		Message: message,
		Error:   "",
		Code:    200,
		Data:    data,
	}
}

func NewHandler(authUC adapters.Currency) ServerInterface {
	return &Handler{
		AuthUC: authUC,
	}
}
