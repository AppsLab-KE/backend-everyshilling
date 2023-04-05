package handlers

import (
	"context"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/internal/dto"
	"github.com/gin-gonic/gin"
)

func (h Handler) Reset(c *gin.Context) {
	//get the request body
	var requestBody dto.RequestResetCredentials
	var responseBody dto.DefaultRes[*dto.ResetRes]

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
	err, _ := h.AuthUC.ResetPassword(ctx, requestBody)
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
