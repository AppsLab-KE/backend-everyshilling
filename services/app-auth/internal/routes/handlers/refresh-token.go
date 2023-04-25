package handlers

import (
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/internal/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h Handler) RefreshToken(c *gin.Context) {
	if c.IsAborted() {
		return
	}

	// Body otpCode
	var responseBody dto.DefaultRes[*dto.RefreshTokenRes]
	var tokenReq dto.RefreshTokenReq

	// Bind request body to struct
	if err := c.ShouldBindJSON(&tokenReq); err != nil {
		responseBody = handleError[*dto.RefreshTokenRes](err)
		c.JSON(http.StatusBadRequest, responseBody)
		return
	}

	// Call usecase
	refreshToken, err := h.AuthUC.RefreshToken(tokenReq)
	if err != nil {
		responseBody = handleError[*dto.RefreshTokenRes](err)
		c.JSON(responseBody.Code, responseBody)
		return
	}

	// Return response
	responseBody = okResponse[*dto.RefreshTokenRes](refreshToken, "token refreshed successfully")
	c.JSON(responseBody.Code, responseBody)

}
