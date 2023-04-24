package handlers

import (
	"context"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/internal/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h Handler) Login(c *gin.Context) {
	var requestBody dto.LoginInitReq
	var responseBody dto.DefaultRes[*dto.LoginInitRes]

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		responseBody = handleError[*dto.LoginInitRes](err)
		c.JSON(http.StatusBadRequest, responseBody)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	usr, err := h.AuthUC.LoginUser(ctx, requestBody)
	if err != nil {
		responseBody = handleError[*dto.LoginInitRes](err)
		c.JSON(responseBody.Code, responseBody)
		return
	}

	responseBody = okResponse[*dto.LoginInitRes](usr, usr.Message)
	c.JSON(responseBody.Code, responseBody)
}

func (h Handler) ResendLoginOTP(c *gin.Context, trackingUuid string) {
	var responseBody dto.DefaultRes[*dto.ResendOTPRes]
	var resendOTPReq dto.ResendOTPReq = dto.ResendOTPReq{
		TrackingUID: trackingUuid,
	}
	res, err := h.AuthUC.ResendLoginOTP(resendOTPReq)
	if err != nil {
		responseBody = handleError[*dto.ResendOTPRes](err)
		c.JSON(responseBody.Code, responseBody)
	}

	responseBody = okResponse[*dto.ResendOTPRes](res, res.Message)
	c.JSON(responseBody.Code, responseBody)
}

func (h Handler) VerifyLoginOTP(c *gin.Context, trackingUuid string) {
	// Get trackingID
	// Body otpCode
	var responseBody dto.DefaultRes[*dto.LoginRes]
	var otpBody dto.LoginOTPBody
	mainCtx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()

	if err := c.ShouldBindJSON(&otpBody); err != nil {
		responseBody = handleError[*dto.LoginRes](err)
		c.JSON(responseBody.Code, responseBody)
		return
	}

	requestBody := dto.OtpVerificationReq{
		TrackingUID: trackingUuid,
		OtpCode:     otpBody.OtpCode,
	}

	res, err := h.AuthUC.VerifyLoginOTP(mainCtx, requestBody)
	if err != nil {
		responseBody = handleError[*dto.LoginRes](err)
		c.JSON(responseBody.Code, responseBody)
		return
	}

	responseBody = okResponse[*dto.LoginRes](res, "login success")
	c.JSON(responseBody.Code, responseBody)
}
