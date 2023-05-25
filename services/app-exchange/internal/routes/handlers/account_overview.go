package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AccountOverview(c *gin.Context) {
	// Process the request and retrieve account details
	account := map[string]interface{}{
		"account_number": 1234567890,
		"balance":        1000.0,
		"currency":       "USD",
	}

	// Return the account overview as a JSON response
	c.JSON(http.StatusOK, account)
}
