package ports

import (
	"context"
	"github.com/AppsLab-KE/be-go-gen-grpc/db"
)

type AccountRepository interface {
	// ACCOUNT
	CreateAccount(ctx context.Context, in *db.CreateAccountRequest) (*db.CreateAccountResponse, error)
	DeleteAccount(ctx context.Context, in *db.DeleteAccountRequest) (*db.DeleteAccountResponse, error)
	UpdateAccount(ctx context.Context, in *db.UpdateAccountRequest) (*db.UpdateAccountResponse, error)
	SearchAccount(ctx context.Context, in *db.SearchAccountRequest) (*db.SearchAccountResponse, error)
}

type AccountsService interface {
	// ACCOUNT
	CreateAccount(ctx context.Context, in *db.CreateAccountRequest) (*db.CreateAccountResponse, error)
	DeleteAccount(ctx context.Context, in *db.DeleteAccountRequest) (*db.DeleteAccountResponse, error)
	UpdateAccount(ctx context.Context, in *db.UpdateAccountRequest) (*db.UpdateAccountResponse, error)
	SearchAccount(ctx context.Context, in *db.SearchAccountRequest) (*db.SearchAccountResponse, error)
}
