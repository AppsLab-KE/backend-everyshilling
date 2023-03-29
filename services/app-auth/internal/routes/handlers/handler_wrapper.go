package handlers

import (
	"context"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/internal/core/usecase"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/internal/dto"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	AuthUC usecase.AuthUseCase
}

func (h Handler) PostLogin(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (h Handler) Register(c *gin.Context) {
	// get request body
	var requestBody dto.RegisterRequest
	var responseBody dto.DefaultRes

	// parse request body
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		responseBody.Message = "Registration failed"
		responseBody.Code = 400
		responseBody.Data = nil
		responseBody.Error = "Invalid request."
		c.JSON(400, responseBody)
		return
	}

	// process the request
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	usr, err := h.AuthUC.RegisterUser(ctx, &requestBody)
	if err != nil {
		responseBody.Message = "Registration failed"
		responseBody.Code = 400
		responseBody.Data = nil
		responseBody.Error = err.Error()
		c.JSON(400, responseBody)
		return
	}

	responseBody.Message = "Registration failed"
	responseBody.Code = 200
	responseBody.Data = usr
	responseBody.Error = ""
	c.JSON(200, responseBody)
}

func (h Handler) PostReset(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (h Handler) PostResetRequestIdChange(c *gin.Context, requestId string) {
	//TODO implement me
	panic("implement me")
}

func (h Handler) PostResetRequestIdVerify(c *gin.Context, requestId string) {
	//TODO implement me
	panic("implement me")
}

func NewHandler(authUC *usecase.AuthUseCase) ServerInterface {
	return &Handler{}
}
