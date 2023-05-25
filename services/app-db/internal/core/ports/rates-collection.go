package ports

import (
	"context"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-db/internal/core/models"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-db/internal/dto"
)

type RatesRepo interface {
	CreateRate(context.Context, []*models.ConversionRate) ([]*models.ConversionRate, error)
	UpdateRate(context.Context, *models.ConversionRate) (*models.ConversionRate, error)
	FetchRates(context.Context, dto.FetchRatesRequest) ([]models.ConversionRate, error)
	DeleteRate(context.Context, string) error
}

type RatesStorage interface {
	CreateRate(context.Context, []*models.ConversionRate) ([]*models.ConversionRate, error)
	UpdateRate(context.Context, *models.ConversionRate) (*models.ConversionRate, error)
	FetchRates(context.Context, dto.FetchRatesRequest) ([]models.ConversionRate, error)
	DeleteRate(context.Context, string) error
}
