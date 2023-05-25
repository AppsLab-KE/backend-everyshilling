package repositories

import (
	"context"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-exchange/internal/core/ports"
	"github.com/AppsLab-KE/be-go-gen-grpc/db"
)

type exchangeRepository struct {
	dbStorage ports.DBStorage
}

func (e exchangeRepository) ReadConversionRate(ctx context.Context, in *db.ReadConversionRateRequest) (*db.ReadConversionRateResponse, error) {
	res, err := e.dbStorage.ReadConversionRate(ctx, in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func NewExchangeRepository(dbStorage ports.DBStorage) ports.ExchangeRepository {
	return &exchangeRepository{
		dbStorage: dbStorage,
	}
}
