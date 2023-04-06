package handlers

import (
	"context"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/internal/dto"
	"github.com/gin-gonic/gin"
)

func (h Handler) Login(c *gin.Context) {
	var requestBody dto.RequestLogin
	var responseBody dto.DefaultRes[*dto.UserLoginRes]

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

	usr, err := h.AuthUC.LoginUser(ctx, requestBody)
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

func (h Handler) VerifyLoginOTP(c *gin.Context, trackingUuid string) {
	//TODO implement me
	panic("implement me")
}
