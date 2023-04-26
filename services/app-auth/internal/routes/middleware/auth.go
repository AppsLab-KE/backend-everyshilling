package middleware

import (
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/internal/dto"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/internal/routes/handlers"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

const (
	AuthorisationHeader       = "Authorization"
	AuthorizationBearerPrefix = "Bearer"
	UserUUIDKey               = "UserUUID"
)

func unauthorisedError() dto.DefaultRes[any] {
	return dto.DefaultRes[any]{
		Message: "request failed: unauthorised",
		Error:   "request not authorised: missing a valid token",
		Code:    http.StatusUnauthorized,
		Data:    nil,
	}
}

func (m *Manager) Auth(ctx *gin.Context) {
	if ctx.IsAborted() {
		return
	}

	if _, exists := ctx.Get(handlers.BearerScopes); exists {
		// generate default error when unauthorised
		unauthorisedResponse := unauthorisedError()

		// get token from header
		bearerToken := ctx.GetHeader(AuthorisationHeader)
		bearerToken = strings.TrimPrefix(bearerToken, AuthorizationBearerPrefix)
		bearerToken = strings.TrimSpace(bearerToken)

		// if token is missing, abort
		if bearerToken == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, unauthorisedResponse)
			return
		}

		// validate token
		userId, err := m.authUC.VerifyAccessToken(bearerToken)
		if err != nil {
			log.Error(err)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, unauthorisedResponse)
			return
		}

		ctx.Set(UserUUIDKey, userId)
	}
	ctx.Next()
}
