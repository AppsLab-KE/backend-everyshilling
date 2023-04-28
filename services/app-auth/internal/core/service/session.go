package service

import (
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/internal/core/adapters"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/internal/dto"
	"net/http"
)

type DefaultSessionService struct {
	repo adapters.AuthRepo
}

func (d DefaultSessionService) Invalidate() dto.DefaultRes[interface{}] {
	err := d.repo.InvalidateSession()
	if err != nil {
		return dto.DefaultRes[interface{}]{
			Message: "Failed to invalidate session",
			Error:   err.Error(),
			Code:    http.StatusInternalServerError,
			Data:    nil,
		}
	}
	return dto.DefaultRes[interface{}]{
		Message: "session invalidated succesfully",
		Error:   "",
		Code:    http.StatusOK,
		Data:    nil,
	}

}

func NewDefaultSessionService(repo adapters.AuthRepo) adapters.SessionService {
	return &DefaultSessionService{repo: repo}
}
