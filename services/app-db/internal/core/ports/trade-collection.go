package ports

import (
	"context"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-db/internal/core/models"
)

// TradeStorage defines the methods for interacting with the trade collection in the data store.
type TradeStorage interface {
	CreateTrade(ctx context.Context, trade *models.Trade) error
	DeleteTrade(ctx context.Context, tradeID string) error
	UpdateTrade(ctx context.Context, trade *models.Trade) error
	GetTrade(ctx context.Context, tradeID string) (*models.Trade, error)
	GetTradeByAccount(ctx context.Context, accountID string) ([]*models.Trade, error)
}

// TradeRepo defines the methods for the trade repository.
type TradeRepo interface {
	CreateTrade(ctx context.Context, trade *models.Trade) error
	DeleteTrade(ctx context.Context, tradeID string) error
	UpdateTrade(ctx context.Context, trade *models.Trade) error
	GetTrade(ctx context.Context, tradeID string) (*models.Trade, error)
	GetTradeByAccount(ctx context.Context, accountID string) ([]*models.Trade, error)
}

// TradeUsecase defines the methods for the trade use case.
type TradeUsecase interface {
	CreateTrade(ctx context.Context, trade *models.Trade) error
	DeleteTrade(ctx context.Context, tradeID string) error
	UpdateTrade(ctx context.Context, trade *models.Trade) error
	GetTrade(ctx context.Context, tradeID string) (*models.Trade, error)
	GetTradeByAccount(ctx context.Context, accountID string) ([]*models.Trade, error)
}
