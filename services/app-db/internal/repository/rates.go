package repository

import (
	"context"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-db/internal/core/models"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-db/internal/core/ports"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-db/internal/dto"
)

type ratesRepo struct {
	ratesStorage ports.RatesStorage
}

func (r ratesRepo) CreateRate(ctx context.Context, rates []*models.ConversionRate) ([]*models.ConversionRate, error) {
	return r.ratesStorage.CreateRate(ctx, rates)
}

func (r ratesRepo) UpdateRate(ctx context.Context, rate *models.ConversionRate) (*models.ConversionRate, error) {
	return r.ratesStorage.UpdateRate(ctx, rate)
}

func (r ratesRepo) FetchRates(ctx context.Context, request dto.FetchRatesRequest) ([]models.ConversionRate, error) {
	return r.ratesStorage.FetchRates(ctx, request)
}

func (r ratesRepo) DeleteRate(ctx context.Context, s string) error {
	return r.ratesStorage.DeleteRate(ctx, s)
}

func NewRatesRepo(ratesStorage ports.RatesStorage) ports.RatesRepo {
	return &ratesRepo{ratesStorage: ratesStorage}
}
