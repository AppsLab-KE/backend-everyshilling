package usecase

import (
	"github.com/AppsLab-KE/backend-everyshilling/services/app-ratescron/internal/core/ports"
	"github.com/AppsLab-KE/be-go-gen-grpc/db"
)

type usecase struct {
	dbRepo    ports.DBRepository
	ratesRepo ports.RatesAPIRepository
}

func NewUsecase(dbRepo ports.DBRepository, ratesRepo ports.RatesAPIRepository) ports.RatesUsecase {
	return &usecase{
		dbRepo:    dbRepo,
		ratesRepo: ratesRepo,
	}
}

func (u usecase) FetchAndStoreRates() error {
	rates, err := u.ratesRepo.GetRates("USD")
	if err != nil {
		return err
	}

	var ratesRequest db.CreateConversionRateRequest
	ratesRequest.ConversionRate = make([]*db.ConversionRate, len(rates))

	for i, rate := range rates {
		ratesRequest.ConversionRate[i] = &db.ConversionRate{
			FromCurrency:       rate.FromCurrency,
			ToCurrency:         rate.ToCurrency,
			Rate:               rate.Rate,
			DateUpdatedUnixUtc: rate.DateUpdatedUnixUTC,
		}
	}

	_, err = u.dbRepo.CreateConversionRate(nil, &ratesRequest)
	if err != nil {
		return err
	}

	return nil
}
