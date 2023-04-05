package middleware

import (
	"github.com/gin-gonic/gin"
)
import log "github.com/sirupsen/logrus"

func (m *Manager) Log(ctx gin.Context) {
	if ctx.IsAborted() {
		return
	}
	ctx.Next()
	status := ctx.Writer.Status()
	size := ctx.Writer.Size()
	url := ctx.Request.URL

	log.Infof("URL: %s Status: %d Size: %d", url, status, size)
}
