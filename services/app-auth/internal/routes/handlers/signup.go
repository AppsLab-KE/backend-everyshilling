package handlers

import (
	"context"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/internal/core/service"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/internal/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

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
	switch err {
	case service.ErrUserExists:
		responseBody.Code = http.StatusConflict
	case service.ErrUserCreation:
		responseBody.Code = http.StatusInternalServerError
	case service.ErrRequestValidation:
		responseBody.Code = http.StatusBadRequest
	default:
		responseBody.Error = ""
	}

	c.JSON(responseBody.Code, responseBody)
}
