package service

import (
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/internal/core/adapters"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/internal/dto"
)

type DefaultSessionService struct {
	repo adapters.AuthRepo
}

func (d DefaultSessionService) Invalidate() dto.DefaultRes {
	//TODO implement me
	panic("implement me")
}

func NewDefaultSessionService(repo adapters.AuthRepo) adapters.SessionService {
	return &DefaultSessionService{repo: repo}
}
