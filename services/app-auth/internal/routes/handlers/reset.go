package handlers

import (
	"context"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/internal/dto"
	"github.com/gin-gonic/gin"
)

func (h Handler) Reset(c *gin.Context) {
	//get the request body
	var requestBody dto.OtpGenReq
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
	_, err := h.AuthUC.SendResetOtp(ctx, requestBody)

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

func (h Handler) VerifyResetOTP(c *gin.Context, trackingUuid string) {
	//get the request body
	var requestBody dto.OtpVerificationReq
	var responseBody dto.DefaultRes[*dto.OtpVerificationRes]

	//parse the request body
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		responseBody.Message = "Verify OTP failed"
		responseBody.Code = 400
		responseBody.Data = nil
		responseBody.Error = "Invalid request."
		c.JSON(400, responseBody)
		return
	}

	//process the request
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	// verify the OTP for password reset
	_, err := h.AuthUC.VerifyResetOTP(ctx, trackingUuid, requestBody)
	if err != nil {
		responseBody.Message = "Verify OTP failed"
		responseBody.Code = 400
		responseBody.Data = nil
		responseBody.Error = "Invalid request"
		c.JSON(400, responseBody)
		return
	}
	responseBody.Message = "OTP verified successfully"
	responseBody.Code = 200
	responseBody.Data = nil
	responseBody.Error = ""
	c.JSON(200, responseBody)

}

func (h Handler) ChangePassword(c *gin.Context, trackingUuid string) {
	// get the request body
	var requestBody dto.ChangePasswordJSONRequestBody
	var responseBody dto.DefaultRes[*dto.ResetRes]

	//parse the request body
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		responseBody.Message = "change password failed"
		responseBody.Code = 400
		responseBody.Data = nil
		responseBody.Error = "Invalid request"
		c.JSON(400, responseBody)
		return
	}
	//process the request
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	// check if the user exists and change the password
	_, err := h.AuthUC.ChangePassword(ctx, trackingUuid, requestBody)
	if err != nil {
		responseBody.Message = "Change password failed"
		responseBody.Code = 400
		responseBody.Data = nil
		responseBody.Error = "Invalid request"
		c.JSON(400, responseBody)
		return
	}
	responseBody.Message = "Password changed successfully"
	responseBody.Code = 200
	responseBody.Data = nil
	responseBody.Error = ""
	c.JSON(200, responseBody)
}
