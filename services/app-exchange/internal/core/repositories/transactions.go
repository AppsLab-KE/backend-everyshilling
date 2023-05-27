package repositories

import (
	"context"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-exchange/internal/core/ports"
	"github.com/AppsLab-KE/be-go-gen-grpc/db"
	"google.golang.org/protobuf/types/known/emptypb"
)

type transactionRepository struct {
	dbStorage ports.DBStorage
}

func (t transactionRepository) CreateTransaction(ctx context.Context, in *db.CreateTransactionRequest) (*db.CreateTransactionResponse, error) {
	res, err := t.dbStorage.CreateTransaction(ctx, in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (t transactionRepository) DeleteTransaction(ctx context.Context, in *db.DeleteTransactionRequest) (emptypb.Empty, error) {
	res, err := t.dbStorage.DeleteTransaction(ctx, in)
	if err != nil {
		return emptypb.Empty{}, err
	}
	return res, nil
}

func (t transactionRepository) UpdateTransaction(ctx context.Context, in *db.UpdateTransactionRequest) (emptypb.Empty, error) {
	res, err := t.dbStorage.UpdateTransaction(ctx, in)
	if err != nil {
		return emptypb.Empty{}, err
	}
	return res, nil
}

func (t transactionRepository) GetTransaction(ctx context.Context, in *db.GetTransactionRequest) (*db.Transaction, error) {
	res, err := t.dbStorage.GetTransaction(ctx, in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (t transactionRepository) GetTransactionByAccount(ctx context.Context, in *db.GetTransactionByAccountRequest) (*db.GetTransactionByAccountResponse, error) {
	res, err := t.dbStorage.GetTransactionByAccount(ctx, in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func NewTransactionRepository(dbStorage ports.DBStorage) ports.TransactionRepository {
	return &transactionRepository{
		dbStorage: dbStorage,
	}
}
