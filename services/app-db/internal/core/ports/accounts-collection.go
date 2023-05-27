package ports

import (
	"context"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-db/internal/core/models"
)

type AccountRepo interface {
	CreateAccount(ctx context.Context, account *models.Account) (*models.Account, error)
	DeleteAccount(ctx context.Context, accountID string) error
	UpdateAccount(ctx context.Context, account *models.Account) (*models.Account, error)
	SearchAccount(ctx context.Context, query string) ([]*models.Account, error)
	GetAccount(ctx context.Context, id string) error
	GetAccounts(ctx context.Context, id string) error
}

type AccountStorage interface {
	CreateAccount(ctx context.Context, account *models.Account) (*models.Account, error)
	DeleteAccount(ctx context.Context, accountID string) error
	UpdateAccount(ctx context.Context, account *models.Account) (*models.Account, error)
	SearchAccount(ctx context.Context, query string) ([]*models.Account, error)
}
