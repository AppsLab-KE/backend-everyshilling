package repositories

import (
	"context"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-exchange/internal/core/ports"
	"github.com/AppsLab-KE/be-go-gen-grpc/db"
)

type accountRepository struct {
	dbStorage ports.DBStorage
}

func (a accountRepository) CreateAccount(ctx context.Context, in *db.CreateAccountRequest) (*db.CreateAccountResponse, error) {
	account, err := a.dbStorage.CreateAccount(ctx, in)
	if err != nil {
		return nil, err
	}
	return account, nil
}

func (a accountRepository) DeleteAccount(ctx context.Context, in *db.DeleteAccountRequest) (*db.DeleteAccountResponse, error) {
	result, err := a.dbStorage.DeleteAccount(ctx, in)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (a accountRepository) UpdateAccount(ctx context.Context, in *db.UpdateAccountRequest) (*db.UpdateAccountResponse, error) {
	res, err := a.dbStorage.UpdateAccount(ctx, in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (a accountRepository) SearchAccount(ctx context.Context, in *db.SearchAccountRequest) (*db.SearchAccountResponse, error) {
	res, err := a.dbStorage.SearchAccount(ctx, in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func NewAccountRepository(dbStorage ports.DBStorage) ports.AccountRepository {
	return &accountRepository{
		dbStorage: dbStorage,
	}
}
