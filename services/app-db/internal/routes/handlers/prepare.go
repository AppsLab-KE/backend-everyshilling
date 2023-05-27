package handlers

import (
	"context"
	"errors"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-db/internal/core/ports"
	"github.com/AppsLab-KE/be-go-gen-grpc/db"
)

var (
	ErrEmptyRequest = errors.New("empty request")
)

type Handler struct {
	db.UnimplementedDbServiceServer
	userRepo  ports.UserRepo
	ratesRepo ports.RatesRepo
}

func (s *Handler) HealthCheck(context.Context, *db.DefaultRequest) (*db.HealthResponse, error) {
	return &db.HealthResponse{
		Message: "healthy",
	}, nil
}

func NewHandler(userRepo ports.UserRepo, ratesRepo ports.RatesRepo) *Handler {
	return &Handler{
		userRepo:  userRepo,
		ratesRepo: ratesRepo,
	}
}
