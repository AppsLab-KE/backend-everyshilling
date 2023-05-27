package repository

import (
	"context"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-db/internal/core/models"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-db/internal/core/ports"
)

type accountRepo struct {
	accountStorage ports.AccountStorage
}

func (r *accountRepo) GetAccount(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}

func (r *accountRepo) GetAccounts(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}

func (r *accountRepo) DeleteAccount(ctx context.Context, accountID string) error {
	err := r.accountStorage.DeleteAccount(ctx, accountID)
	if err != nil {

		return err
	}

	return nil
}

func (r *accountRepo) UpdateAccount(ctx context.Context, account *models.Account) (*models.Account, error) {
	updatedAccount, err := r.accountStorage.UpdateAccount(ctx, account)
	if err != nil {

		return nil, err
	}

	return updatedAccount, nil
}

func (r *accountRepo) SearchAccount(ctx context.Context, query string) ([]*models.Account, error) {
	accounts, err := r.accountStorage.SearchAccount(ctx, query)
	if err != nil {

		return nil, err
	}

	return accounts, nil
}

func (r *accountRepo) CreateAccount(ctx context.Context, account *models.Account) (*models.Account, error) {
	createdAccount, err := r.accountStorage.CreateAccount(ctx, account)
	if err != nil {

		return nil, err
	}

	return createdAccount, nil
}

func NewAccountRepo(accountStorage ports.AccountStorage) ports.AccountRepo {
	return &accountRepo{accountStorage: accountStorage}
}
