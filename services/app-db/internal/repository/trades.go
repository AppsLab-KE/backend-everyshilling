package repository

import (
	"context"
	"errors"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-db/internal/core/models"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-db/internal/core/ports"
	"log"
)

type TradeRepository struct {
	tradeStorage ports.TradeStorage
}

func NewTradeRepository(tradeStorage ports.TradeStorage) *TradeRepository {
	return &TradeRepository{
		tradeStorage: tradeStorage,
	}
}

func (r *TradeRepository) CreateTrade(ctx context.Context, trade *models.Trade) error {
	err := r.tradeStorage.CreateTrade(ctx, trade)
	if err != nil {
		log.Println("Error creating trade:", err)
		return errors.New("failed to create trade")
	}
	return nil
}

func (r *TradeRepository) DeleteTrade(ctx context.Context, tradeID string) error {
	err := r.tradeStorage.DeleteTrade(ctx, tradeID)
	if err != nil {
		log.Println("Error deleting trade:", err)
		return errors.New("failed to delete trade")
	}
	return nil
}

func (r *TradeRepository) UpdateTrade(ctx context.Context, trade *models.Trade) error {
	err := r.tradeStorage.UpdateTrade(ctx, trade)
	if err != nil {
		log.Println("Error updating trade:", err)
		return errors.New("failed to update trade")
	}
	return nil
}

func (r *TradeRepository) GetTrade(ctx context.Context, tradeID string) (*models.Trade, error) {
	trade, err := r.tradeStorage.GetTrade(ctx, tradeID)
	if err != nil {
		log.Println("Error retrieving trade:", err)
		return nil, errors.New("failed to retrieve trade")
	}
	return trade, nil
}

func (r *TradeRepository) GetTradeByAccount(ctx context.Context, accountID string) ([]*models.Trade, error) {
	trades, err := r.tradeStorage.GetTradeByAccount(ctx, accountID)
	if err != nil {
		log.Println("Error retrieving trades by account:", err)
		return nil, errors.New("failed to retrieve trades by account")
	}
	return trades, nil
}
