package handlers

import (
	"context"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/internal/core/usecase"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/internal/dto"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/internal/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	AuthUC usecase.AuthUseCase
}

func (h Handler) Login(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (h Handler) VerifyLoginOTP(c *gin.Context, trackingUuid string) {
	//TODO implement me
	panic("implement me")
}

func (h Handler) Reset(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (h Handler) ChangePassword(c *gin.Context, trackingUuid string) {
	//TODO implement me
	panic("implement me")
}

func (h Handler) VerifyResetOTP(c *gin.Context, trackingUuid string) {
	//TODO implement me
	panic("implement me")
}

func (h Handler) PostLogin(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (h Handler) Register(c *gin.Context) {
	// get request body
	var requestBody dto.RegisterRequest
	var responseBody dto.DefaultRes[*dto.UserRegistrationRes]

	responseBody.Message = "Registration failed"

	// parse request body
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		responseBody.Code = 400
		responseBody.Data = nil
		responseBody.Error = "Invalid request."
		c.JSON(400, responseBody)
		return
	}

	// process the request
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	usr, err := h.AuthUC.RegisterUser(ctx, requestBody)
	if err == nil {
		responseBody.Message = "Registration success"
		responseBody.Code = 200
		responseBody.Data = usr
		responseBody.Error = ""
		c.JSON(200, responseBody)
		return
	}

	responseBody.Error = err.Error()
	responseBody.Message = "Registration failed"

	switch err.(type) {
	case *errors.ErrUserExists:
		responseBody.Code = http.StatusConflict
	case *errors.ErrUserCreation:
		responseBody.Code = http.StatusInternalServerError
	case *errors.ErrRequestValidation:
		responseBody.Code = http.StatusBadRequest
	default:
		responseBody.Error = ""

	}

	c.JSON(responseBody.Code, responseBody)
}

func NewHandler(authUC *usecase.AuthUseCase) ServerInterface {
	return &Handler{}
}
