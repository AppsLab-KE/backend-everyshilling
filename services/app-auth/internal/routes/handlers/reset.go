package handlers

import (
	"context"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/internal/dto"
	"github.com/gin-gonic/gin"
)

func (h Handler) Reset(c *gin.Context) {
	if c.IsAborted() {
		return
	}
	//get the request body
	var requestBody dto.OtpGenReq
	var responseBody dto.DefaultRes[*dto.OtpGenRes]

	//parse the request body
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		responseBody = handleError[*dto.OtpGenRes](err)
		c.JSON(responseBody.Code, responseBody)
		return
	}
	//process the request
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	// check if the user exists and generate a reset request ID
	res, err := h.AuthUC.SendResetOtp(ctx, requestBody)

	if err != nil {
		responseBody = handleError[*dto.OtpGenRes](err)
		c.JSON(responseBody.Code, responseBody)
		c.JSON(400, responseBody)
		return
	}
	responseBody = okResponse[*dto.OtpGenRes](res, res.Message)
	c.JSON(responseBody.Code, responseBody)
}

func (h Handler) VerifyResetOTP(c *gin.Context, trackingUuid string) {
	if c.IsAborted() {
		return
	}
	//get the request body
	var requestBody dto.RequestOTP
	var responseBody dto.DefaultRes[*dto.OtpVerificationRes]

	//parse the request body
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		responseBody = handleError[*dto.OtpVerificationRes](err)
		c.JSON(responseBody.Code, responseBody)
		return
	}

	otpVerificationReq := dto.OtpVerificationReq{
		TrackingUID: trackingUuid,
		OtpCode:     requestBody.OtpCode,
	}
	//process the request
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	// verify the OTP for password reset
	verificationRes, err := h.AuthUC.VerifyResetOTP(ctx, trackingUuid, otpVerificationReq)
	if err != nil {
		responseBody = handleError[*dto.OtpVerificationRes](err)
		c.JSON(responseBody.Code, responseBody)
		return
	}
	responseBody = okResponse[*dto.OtpVerificationRes](verificationRes, "otp verified successfully")
	c.JSON(responseBody.Code, responseBody)
}

func (h Handler) ResendResetOTP(c *gin.Context, trackingUuid string) {
	if c.IsAborted() {
		return
	}
	var responseBody dto.DefaultRes[*dto.ResendOTPRes]
	var resendOTPReq dto.ResendOTPReq = dto.ResendOTPReq{
		TrackingUID: trackingUuid,
	}
	res, err := h.AuthUC.ResendResetOTP(resendOTPReq)
	if err != nil {
		responseBody = handleError[*dto.ResendOTPRes](err)
		c.JSON(responseBody.Code, responseBody)
	}

	responseBody = okResponse[*dto.ResendOTPRes](res, res.Message)
	c.JSON(responseBody.Code, responseBody)
}

func (h Handler) ChangePassword(c *gin.Context, trackingUuid string) {
	if c.IsAborted() {
		return
	}
	// get the request body
	var requestBody dto.ChangePasswordJSONRequestBody
	var responseBody dto.DefaultRes[*dto.ResetRes]

	//parse the request body
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		responseBody = handleError[*dto.ResetRes](err)
		c.JSON(responseBody.Code, responseBody)
		return
	}
	//process the request
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	// check if the user exists and change the password
	res, err := h.AuthUC.ChangePassword(ctx, trackingUuid, requestBody)
	if err != nil {
		responseBody = handleError[*dto.ResetRes](err)
		c.JSON(responseBody.Code, responseBody)
		return
	}

	responseBody = okResponse[*dto.ResetRes](res, "Password change successful")
	c.JSON(responseBody.Code, responseBody)
}
