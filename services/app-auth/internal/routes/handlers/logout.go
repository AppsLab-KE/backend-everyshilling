package handlers

import (
	"errors"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/internal/dto"
	"github.com/gin-gonic/gin"
)

func (h Handler) Logout(c *gin.Context) {
	if c.IsAborted() {
		return
	}
	var responseBody dto.DefaultRes[any]

	// Get the token from the header
	uuid, exists := c.Get("UserUUID")
	if !exists {
		err := errors.New("fatal error: user uuid not found in context")
		responseBody = handleError[any](err)
		c.JSON(responseBody.Code, responseBody)
		return
	}

	// cast to string
	userUUID, ok := uuid.(string)
	if !ok {
		err := errors.New("fatal error: user uuid not found in context")
		responseBody = handleError[any](err)
		c.JSON(responseBody.Code, responseBody)
		return
	}

	// Return response
	err := h.AuthUC.Logout(userUUID)
	if err != nil {
		responseBody = handleError[any](err)
		c.JSON(responseBody.Code, responseBody)
		return
	}

	responseBody = okResponse[any](nil, "user logged out successfully")
	c.JSON(responseBody.Code, responseBody)
}
