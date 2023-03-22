package service

import (
	"appslab.co.ke/everyshilling/app-db/internal/core/ports"
	"context"
)

type DefaultUserService struct {
	userRepo ports.UserRepo
}

func (d DefaultUserService) CreateUser(ctx context.Context) {
	//TODO implement me
	panic("implement me")
}

var _ ports.UserService = (*DefaultUserService)(nil)

func New(repo ports.UserRepo) *DefaultUserService {
	return &DefaultUserService{userRepo: repo}
}
