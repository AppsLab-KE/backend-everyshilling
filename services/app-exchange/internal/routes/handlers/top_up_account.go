package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func TopUpAccount(c *gin.Context) {
	// Access query parameter
	amount := c.Query("amount")

	// Validate the amount parameter
	if amount == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Missing amount parameter",
		})
		return
	}

	// Process the request and return a response
	c.JSON(http.StatusOK, gin.H{
		"message": "TopUpAccount handler",
		"amount":  amount,
	})
}
