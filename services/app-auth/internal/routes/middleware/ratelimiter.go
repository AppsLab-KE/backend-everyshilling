package middleware

import (
	"github.com/gin-gonic/gin"
)

// RateLimiter Uses LeakyBuckets Rate limiting Algorithm
func (m *Manager) RateLimiter(ctx *gin.Context) {
	if ctx.IsAborted() {
		return
	}

	// TODO
	ctx.Next()
}
