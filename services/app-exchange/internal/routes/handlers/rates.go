package handlers

import (
	"context"
	"errors"
	"github.com/AppsLab-KE/be-go-gen-grpc/db"
	"github.com/AppsLab-KE/be-go-gen-grpc/exchange"
)

func (h Handler) CreateConversionRate(ctx context.Context, request *exchange.CreateConversionRateRequest) (*exchange.CreateConversionRateResponse, error) {
	if request == nil {
		return nil, ErrEmptyRequest
	}

	return nil, errors.New("feature moved")
}

func (h Handler) ReadConversionRate(ctx context.Context, request *exchange.ReadConversionRateRequest) (*exchange.ReadConversionRateResponse, error) {
	if request == nil {
		return nil, ErrEmptyRequest
	}

	readConversionRateRequest := db.ReadConversionRateRequest{
		FromUnixUtc: request.FromUnixUtc,
		ToUnixUtc:   request.ToUnixUtc,
	}

	readConversionRateResponse, err := h.exchangeService.ReadConversionRate(ctx, &readConversionRateRequest)
	if err != nil {
		return nil, err
	}

	readConversionRateResponseResult := &exchange.ReadConversionRateResponse{
		ConversionRate: []*exchange.ConversionRate{},
	}

	for _, conversionRate := range readConversionRateResponse.ConversionRate {
		readConversionRateResponseResult.ConversionRate = append(readConversionRateResponseResult.ConversionRate, &exchange.ConversionRate{
			FromCurrency:       conversionRate.FromCurrency,
			ToCurrency:         conversionRate.ToCurrency,
			Rate:               conversionRate.Rate,
			DateUpdatedUnixUtc: conversionRate.DateUpdatedUnixUtc,
		})
	}

	return readConversionRateResponseResult, nil
}

func (h Handler) UpdateConversionRate(ctx context.Context, request *exchange.UpdateConversionRateRequest) (*exchange.UpdateConversionRateResponse, error) {
	if request == nil {
		return nil, ErrEmptyRequest
	}

	return nil, errors.New("feature moved")
}

func (h Handler) DeleteConversionRate(ctx context.Context, request *exchange.DeleteConversionRateRequest) (*exchange.DeleteConversionRateResponse, error) {
	if request == nil {
		return nil, ErrEmptyRequest
	}

	return nil, errors.New("feature moved")
}
