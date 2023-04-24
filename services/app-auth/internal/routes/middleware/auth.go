package middleware

import (
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/internal/dto"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/internal/routes/handlers"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/pkg/tokens"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	AuthorisationHeader       = "Authorisation"
	AuthorisationHeaderPrefix = "Bearer"
	UserUIDKey                = "UserUID"
)

func unauthorisedError() dto.DefaultRes[any] {
	return dto.DefaultRes[any]{
		Message: "Request failed",
		Error:   "Unauthorised request",
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
		bearerToken = strings.TrimPrefix(AuthorisationHeaderPrefix, bearerToken)

		// if token is missing, abort
		if bearerToken == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, unauthorisedResponse)
			return
		}

		// validate token
		userId, err := tokens.VerifyToken(bearerToken, m.config.Jwt.Secret)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, unauthorisedResponse)
			return
		}

		ctx.Set(UserUIDKey, userId)
	}
	ctx.Next()
}
