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
	//get the request body
	var requestBody dto.RequestResetCredentials
	var responseBody dto.DefaultRes[*dto.UserLoginRes]

	//parse the request body
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		responseBody.Message = "Password reset failed"
		responseBody.Code = 400
		responseBody.Data = nil
		responseBody.Error = "Invalid request."
		c.JSON(400, responseBody)
		return
	}
	//process the request
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	// check if the user exists and generate a reset request ID
	err, _ := h.AuthUC.ResetPassword(ctx, &requestBody)
	if err != nil {
		responseBody.Message = "Password reset failed"
		responseBody.Code = 400
		responseBody.Data = nil
		responseBody.Error = "Invalid request"
		c.JSON(400, responseBody)
		return
	}
	responseBody.Message = "Password reset successful"
	responseBody.Code = 200
	responseBody.Data = nil
	responseBody.Error = ""
	c.JSON(200, responseBody)
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

	var requestBody dto.RequestLogin
	var responseBody dto.DefaultRes

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		responseBody.Message = "Bad request: missing email or phone number"
		responseBody.Code = 400
		responseBody.Data = nil
		responseBody.Error = "Missing or incomplete credentials"
		c.JSON(400, responseBody)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	usr, err := h.AuthUC.LoginUser(ctx, &requestBody)
	if err != nil {
		responseBody.Message = "Login Process failed"
		responseBody.Code = 400
		responseBody.Data = nil
		responseBody.Error = err.Error()
		c.JSON(400, responseBody)
		return
	}

	responseBody.Message = "Login Process successful"
	responseBody.Code = 200
	responseBody.Data = usr
	responseBody.Error = ""
	c.JSON(200, responseBody)
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
