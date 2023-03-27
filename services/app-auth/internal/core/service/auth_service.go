package service

import (
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/internal/core/adapters"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/internal/dto"
)

type DefaultAuthService struct {
	repo adapters.AuthRepo
}

func (d DefaultAuthService) RequestOtp(request dto.OtpReq) dto.DefaultRes {
	//TODO implement me
	panic("implement me")
}

func NewDefaultAuthService(repo adapters.AuthRepo) adapters.AuthService {
	return &DefaultAuthService{repo: repo}
}
