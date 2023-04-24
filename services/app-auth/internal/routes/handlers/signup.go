package handlers

import (
	"context"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/internal/dto"
	"github.com/gin-gonic/gin"
)

func (h Handler) Register(c *gin.Context) {
	// get request body
	var requestBody dto.RegisterRequest
	var responseBody dto.DefaultRes[*dto.UserRegistrationRes]

	// parse request body
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		responseBody = handleError[*dto.UserRegistrationRes](err)
		c.JSON(responseBody.Code, responseBody)
		return
	}

	// process the request
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	usr, err := h.AuthUC.RegisterUser(ctx, requestBody)
	if err != nil {
		responseBody = handleError[*dto.UserRegistrationRes](err)
		c.JSON(responseBody.Code, responseBody)
		return
	}
	responseBody = okResponse[*dto.UserRegistrationRes](usr, "registration success")
	c.JSON(responseBody.Code, responseBody)
}
