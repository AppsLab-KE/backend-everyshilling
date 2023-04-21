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
		responseBody = badRequest[*dto.LoginInitRes](err.Error())
		c.JSON(http.StatusBadRequest, responseBody)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	usr, err := h.AuthUC.LoginUser(ctx, requestBody)
	if err != nil {
		responseBody = badRequest[*dto.LoginInitRes](err.Error())
		// TODO Process error types
		c.JSON(http.StatusBadRequest, responseBody)
		return
	}

	responseBody = okResponse[*dto.LoginInitRes](usr)
	c.JSON(200, responseBody)
}

func (h Handler) VerifyLoginOTP(c *gin.Context, trackingUuid string) {
	// Get trackingID
	// Body otpCode
	var responseBody dto.DefaultRes[*dto.OtpVerificationRes]
	var otpBody dto.LoginOTPBody
	mainCtx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()

	if err := c.ShouldBindJSON(&otpBody); err != nil {
		responseBody = badRequest[*dto.OtpVerificationRes](err.Error())
		c.JSON(http.StatusBadRequest, responseBody)
		return
	}

	requestBody := dto.OtpVerificationReq{
		TrackingUID: trackingUuid,
		OtpCode:     trackingUuid,
	}

	res, err := h.AuthUC.VerifyLoginOTP(mainCtx, requestBody)
	if err != nil {
		responseBody = badRequest[*dto.OtpVerificationRes](err.Error())
		// TODO Process error types
		c.JSON(http.StatusBadRequest, responseBody)
		return
	}

	responseBody = okResponse[*dto.OtpVerificationRes](res)
	c.JSON(http.StatusOK, responseBody)
}
