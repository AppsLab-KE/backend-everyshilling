package handlers

import (
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/internal/dto"
	"github.com/gin-gonic/gin"
)

func (h Handler) VerifyPhone(c *gin.Context) {
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

	// check if the user exists and generate a reset request ID
	res, err := h.AuthUC.SendVerifyPhoneOTP(requestBody)

	if err != nil {
		responseBody = handleError[*dto.OtpGenRes](err)
		c.JSON(responseBody.Code, responseBody)
		return
	}

	responseBody = okResponse[*dto.OtpGenRes](res, res.Message)
	c.JSON(responseBody.Code, responseBody)
}

func (h Handler) ResendVerificationOTP(c *gin.Context, trackingUuid string) {
	if c.IsAborted() {
		return
	}
	var responseBody dto.DefaultRes[*dto.ResendOTPRes]
	var resendOTPReq dto.ResendOTPReq = dto.ResendOTPReq{
		TrackingUID: trackingUuid,
	}
	res, err := h.AuthUC.ResendVerifyPhoneOTP(resendOTPReq)
	if err != nil {
		responseBody = handleError[*dto.ResendOTPRes](err)
		c.JSON(responseBody.Code, responseBody)
	}

	responseBody = okResponse[*dto.ResendOTPRes](res, res.Message)
	c.JSON(responseBody.Code, responseBody)
}

func (h Handler) VerifyVerificationOTP(c *gin.Context, trackingUuid string) {
	if c.IsAborted() {
		return
	}
	// Body otpCode
	var responseBody dto.DefaultRes[*dto.OtpVerificationRes]
	var otpBody dto.RequestOTP

	if err := c.ShouldBindJSON(&otpBody); err != nil {
		responseBody = handleError[*dto.OtpVerificationRes](err)
		c.JSON(responseBody.Code, responseBody)
		return
	}

	requestBody := dto.OtpVerificationReq{
		TrackingUID: trackingUuid,
		OtpCode:     otpBody.OtpCode,
	}

	res, err := h.AuthUC.VerifyPhoneOTP(requestBody)
	if err != nil {
		responseBody = handleError[*dto.OtpVerificationRes](err)
		c.JSON(responseBody.Code, responseBody)
		return
	}

	responseBody = okResponse[*dto.OtpVerificationRes](res, res.Message)
	c.JSON(responseBody.Code, responseBody)
}
