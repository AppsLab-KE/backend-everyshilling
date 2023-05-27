package ports

import (
	"context"
	"github.com/AppsLab-KE/be-go-gen-grpc/db"
)

type ExchangeRepository interface {
	ReadConversionRate(ctx context.Context, in *db.ReadConversionRateRequest) (*db.ReadConversionRateResponse, error)
}

type ExchangeService interface {
	ReadConversionRate(ctx context.Context, in *db.ReadConversionRateRequest) (*db.ReadConversionRateResponse, error)
}
