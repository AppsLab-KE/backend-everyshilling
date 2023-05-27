package services

import (
	"context"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-exchange/internal/core/ports"
	"github.com/AppsLab-KE/be-go-gen-grpc/db"
)

type exchangeService struct {
	exchangeRepo ports.ExchangeRepository
}

func (e exchangeService) ReadConversionRate(ctx context.Context, in *db.ReadConversionRateRequest) (*db.ReadConversionRateResponse, error) {
	return e.exchangeRepo.ReadConversionRate(ctx, in)
}

func NewExchangeService(exchangeRepo ports.ExchangeRepository) ports.ExchangeService {
	return &exchangeService{exchangeRepo: exchangeRepo}
}
