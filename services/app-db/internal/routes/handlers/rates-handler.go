package handlers

import (
	"context"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-db/internal/core/models"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-db/internal/dto"
	"github.com/AppsLab-KE/be-go-gen-grpc/db"
)

// CreateConversionRate creates a conversion rate
func (h *Handler) CreateConversionRate(ctx context.Context, req *db.CreateConversionRateRequest) (*db.CreateConversionRateResponse, error) {
	if req == nil {
		return nil, ErrEmptyRequest
	}

	var rates []*models.ConversionRate
	for _, rate := range req.ConversionRate {
		rates = append(rates, &models.ConversionRate{
			FromCurrency: rate.FromCurrency,
			ToCurrency:   rate.ToCurrency,
			Rate:         rate.Rate,
			TimeStampUTC: rate.DateUpdatedUnixUtc,
		})
	}

	createdRates, err := h.ratesRepo.CreateRate(ctx, rates)
	if err != nil {
		return nil, err
	}

	// Example response
	resp := &db.CreateConversionRateResponse{
		// Set the response fields accordingly
	}

	for _, rate := range createdRates {
		resp.ConversionRate = append(resp.ConversionRate, &db.ConversionRate{
			FromCurrency:       rate.FromCurrency,
			ToCurrency:         rate.ToCurrency,
			Rate:               rate.Rate,
			DateUpdatedUnixUtc: rate.TimeStampUTC,
		})
	}

	return resp, nil
}

// ReadConversionRate reads a conversion rate
func (h *Handler) ReadConversionRate(ctx context.Context, req *db.ReadConversionRateRequest) (*db.ReadConversionRateResponse, error) {
	if req == nil {
		return nil, ErrEmptyRequest
	}

	var fetchRequest dto.FetchRatesRequest = dto.FetchRatesRequest{
		FromUnixUtc: req.FromUnixUtc,
		ToUnixUtc:   req.ToUnixUtc,
	}

	rates, err := h.ratesRepo.FetchRates(ctx, fetchRequest)
	if err != nil {
		return nil, err
	}

	// Example response
	resp := &db.ReadConversionRateResponse{
		// Set the response fields accordingly
	}

	for _, rate := range rates {
		resp.ConversionRate = append(resp.ConversionRate, &db.ConversionRate{
			FromCurrency:       rate.FromCurrency,
			ToCurrency:         rate.ToCurrency,
			Rate:               rate.Rate,
			DateUpdatedUnixUtc: rate.TimeStampUTC,
		})
	}

	return resp, nil
}

// UpdateConversionRate updates a conversion rate
func (h *Handler) UpdateConversionRate(ctx context.Context, req *db.UpdateConversionRateRequest) (*db.UpdateConversionRateResponse, error) {
	if req == nil {
		return nil, ErrEmptyRequest
	}

	var request *models.ConversionRate = &models.ConversionRate{
		FromCurrency: req.ConversionRate.FromCurrency,
		ToCurrency:   req.ConversionRate.ToCurrency,
		Rate:         req.ConversionRate.Rate,
		TimeStampUTC: req.ConversionRate.DateUpdatedUnixUtc,
	}

	_, err := h.ratesRepo.UpdateRate(ctx, request)
	if err != nil {
		return nil, err
	}

	// Example response
	resp := &db.UpdateConversionRateResponse{
		// Set the response fields accordingly
	}

	return resp, nil
}

// DeleteConversionRate deletes a conversion rate
func (h *Handler) DeleteConversionRate(ctx context.Context, req *db.DeleteConversionRateRequest) (*db.DeleteConversionRateResponse, error) {
	if req == nil {
		return nil, ErrEmptyRequest
	}

	var request string = req.Uuid

	err := h.ratesRepo.DeleteRate(ctx, request)
	if err != nil {
		return nil, err
	}

	// Example response
	resp := &db.DeleteConversionRateResponse{
		// Set the response fields accordingly
	}

	return resp, nil
}
