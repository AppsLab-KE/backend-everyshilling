package handlers

import (
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/internal/core/usecase"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	AuthUC usecase.AuthUseCase
}

func (h Handler) PostLogin(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (h Handler) Register(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (h Handler) PostReset(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (h Handler) PostResetRequestIdChange(c *gin.Context, requestId string) {
	//TODO implement me
	panic("implement me")
}

func (h Handler) PostResetRequestIdVerify(c *gin.Context, requestId string) {
	//TODO implement me
	panic("implement me")
}

func NewHandler(authUC usecase.AuthUseCase) ServerInterface {
	return &Handler{}
}
