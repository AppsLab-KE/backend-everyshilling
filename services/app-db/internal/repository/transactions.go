package repository

import (
	"context"
	"errors"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-db/internal/core/models"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-db/internal/core/ports"
)

type transactionRepo struct {
	transactionStorage ports.TransactionStorage
}

func NewTransactionRepo(transactionStorage ports.TransactionStorage) ports.TransactionRepo {
	return &transactionRepo{
		transactionStorage: transactionStorage,
	}
}

func (r *transactionRepo) CreateTransaction(ctx context.Context, transaction *models.Transaction) error {
	err := r.transactionStorage.CreateTransaction(ctx, transaction)
	if err != nil {
		return errors.New("failed to create transaction")
	}
	return nil
}

func (r *transactionRepo) DeleteTransaction(ctx context.Context, transactionID string) error {
	err := r.transactionStorage.DeleteTransaction(ctx, transactionID)
	if err != nil {
		return errors.New("failed to delete transaction")
	}
	return nil
}

func (r *transactionRepo) UpdateTransaction(ctx context.Context, transaction *models.Transaction) error {
	err := r.transactionStorage.UpdateTransaction(ctx, transaction)
	if err != nil {
		return errors.New("failed to update transaction")
	}
	return nil
}

func (r *transactionRepo) GetTransaction(ctx context.Context, transactionID string) (*models.Transaction, error) {
	transaction, err := r.transactionStorage.GetTransaction(ctx, transactionID)
	if err != nil {
		return nil, errors.New("failed to retrieve transaction")
	}
	return transaction, nil
}

func (r *transactionRepo) GetTransactionByAccount(ctx context.Context, accountID string) ([]*models.Transaction, error) {
	transactions, err := r.transactionStorage.GetTransactionsByAccount(ctx, accountID)
	if err != nil {
		return nil, errors.New("failed to retrieve transactions by account")
	}
	return transactions, nil
}
