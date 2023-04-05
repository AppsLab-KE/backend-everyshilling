package middleware

import (
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/internal/dto"
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

func (m *Manager) Auth(ctx gin.Context) {
	if ctx.IsAborted() {
		return
	}

	unauthorisedResponse := unauthorisedError()
	bearerToken := ctx.GetHeader(AuthorisationHeader)
	bearerToken = strings.TrimPrefix(AuthorisationHeaderPrefix, bearerToken)
	if bearerToken == "" {
		ctx.JSON(http.StatusUnauthorized, unauthorisedResponse)
		ctx.Abort()
		return
	}

	userId, err := tokens.VerifyToken(bearerToken, m.config.Jwt.Secret)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, unauthorisedResponse)
		ctx.Abort()
		return
	}

	ctx.Set(UserUIDKey, userId)
	ctx.Next()
}
