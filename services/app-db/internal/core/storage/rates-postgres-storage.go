package storage

import (
	"context"
	"errors"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-db/internal/core/models"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-db/internal/core/ports"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-db/internal/dto"
	"gorm.io/gorm"
	"sort"
)

type ratesPostgresStorage struct {
	db *gorm.DB
}

// SortConversionRatesByTimestamp sorts the ConversionRate slice in descending order based on TimeStampUTC
func sortConversionRatesByTimestamp(rates []*models.ConversionRate) {
	sort.Slice(rates, func(i, j int) bool {
		return rates[i].TimeStampUTC > rates[j].TimeStampUTC
	})
}

func (r ratesPostgresStorage) CreateRate(ctx context.Context, rates []*models.ConversionRate) ([]*models.ConversionRate, error) {
	// sort the rates by timestamp starting with the latest
	sortConversionRatesByTimestamp(rates)

	// check if the range of rates already exists and return the counts of the rates
	var count int64
	r.db.Model(&models.ConversionRate{}).Where("time_stamp_utc BETWEEN ? AND ?", rates[len(rates)-1].TimeStampUTC, rates[0].TimeStampUTC).Count(&count)

	if count > 0 {
		return nil, errors.New("rates already exist")
	}

	// create the rates
	if err := r.db.Create(&rates).Error; err != nil {
		return nil, err
	}

	return rates, nil
}

func (r ratesPostgresStorage) UpdateRate(ctx context.Context, rate *models.ConversionRate) (*models.ConversionRate, error) {
	err := r.db.Model(&models.ConversionRate{}).Where("time_stamp_utc = ?", rate.TimeStampUTC).Updates(rate).Error
	if err != nil {
		return nil, err
	}

	return rate, nil
}

func (r ratesPostgresStorage) FetchRates(ctx context.Context, request dto.FetchRatesRequest) ([]models.ConversionRate, error) {
	var rates []models.ConversionRate

	// fetch in the specified range
	if request.ToUnixUtc != 0 && request.FromUnixUtc != 0 {
		err := r.db.Where("time_stamp_utc BETWEEN ? AND ?", request.FromUnixUtc, request.ToUnixUtc).Find(&rates).Error
		if err != nil {
			return nil, err
		}
		return rates, nil
	}

	// fetch the latest
	err := r.db.Order("time_stamp_utc desc").Limit(50).Find(&rates).Error
	if err != nil {
		return nil, err
	}

	return rates, nil
}

func (r ratesPostgresStorage) DeleteRate(ctx context.Context, s string) error {
	err := r.db.Where("time_stamp_utc = ?", s).Delete(&models.ConversionRate{}).Error
	if err != nil {
		return err
	}
	return nil
}

func NewRatesPostgresStorage(db *gorm.DB) ports.RatesStorage {
	return &ratesPostgresStorage{db: db}
}
