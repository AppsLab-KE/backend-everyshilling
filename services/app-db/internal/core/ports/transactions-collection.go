package ports

import (
	"context"
	"errors"

	"github.com/AppsLab-KE/backend-everyshilling/services/app-db/internal/core/models"
)

var (
	ErrNotFound = errors.New("not found")
)

type TransactionRepo interface {
	CreateTransaction(ctx context.Context, transaction *models.Transaction) error
	DeleteTransaction(ctx context.Context, transactionID string) error
	UpdateTransaction(ctx context.Context, transaction *models.Transaction) error
	GetTransaction(ctx context.Context, transactionID string) (*models.Transaction, error)
	GetTransactionByAccount(ctx context.Context, accountID string) ([]*models.Transaction, error)
}

type TransactionStorage interface {
	CreateTransaction(ctx context.Context, transaction *models.Transaction) error
	DeleteTransaction(ctx context.Context, transactionID string) error
	UpdateTransaction(ctx context.Context, transaction *models.Transaction) error
	GetTransaction(ctx context.Context, transactionID string) (*models.Transaction, error)
	GetTransactionsByAccount(ctx context.Context, accountID string) ([]*models.Transaction, error)
}
