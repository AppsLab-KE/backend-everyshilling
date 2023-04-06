package handlers

import (
	"context"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/internal/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h Handler) Login(c *gin.Context) {
	var requestBody dto.RequestLogin
	var responseBody dto.DefaultRes[*dto.UserLoginRes]

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		responseBody = badRequest[*dto.UserLoginRes](err.Error())
		c.JSON(http.StatusBadRequest, responseBody)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	usr, err := h.AuthUC.LoginUser(ctx, requestBody)
	if err != nil {
		responseBody = badRequest[*dto.UserLoginRes](err.Error())
		// TODO Process error types
		c.JSON(http.StatusBadRequest, responseBody)
		return
	}

	responseBody = okResponse[*dto.UserLoginRes](usr)
	c.JSON(200, responseBody)
}

func (h Handler) VerifyLoginOTP(c *gin.Context, trackingUuid string) {
	// Get trackingID
	// Body otpCode
	var requestBody dto.OtpVerificationReq
	var responseBody dto.DefaultRes[*dto.OtpVerificationRes]

	mainCtx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		responseBody = badRequest[*dto.OtpVerificationRes](err.Error())
		c.JSON(http.StatusBadRequest, responseBody)
		return
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
