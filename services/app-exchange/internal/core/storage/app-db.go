package storage

import (
	"context"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-exchange/internal/core/ports"
	"github.com/AppsLab-KE/be-go-gen-grpc/db"
	"google.golang.org/protobuf/types/known/emptypb"
)

type dbStorage struct {
	client db.DbServiceClient
}

func (d dbStorage) HealthCheck(ctx context.Context, in *db.DefaultRequest) (*db.HealthResponse, error) {
	res, err := d.client.HealthCheck(ctx, in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (d dbStorage) ReadConversionRate(ctx context.Context, in *db.ReadConversionRateRequest) (*db.ReadConversionRateResponse, error) {
	res, err := d.client.ReadConversionRate(ctx, in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (d dbStorage) CreateAccount(ctx context.Context, in *db.CreateAccountRequest) (*db.CreateAccountResponse, error) {
	res, err := d.client.CreateAccount(ctx, in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (d dbStorage) DeleteAccount(ctx context.Context, in *db.DeleteAccountRequest) (*db.DeleteAccountResponse, error) {
	res, err := d.client.DeleteAccount(ctx, in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (d dbStorage) UpdateAccount(ctx context.Context, in *db.UpdateAccountRequest) (*db.UpdateAccountResponse, error) {
	res, err := d.client.UpdateAccount(ctx, in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (d dbStorage) SearchAccount(ctx context.Context, in *db.SearchAccountRequest) (*db.SearchAccountResponse, error) {
	res, err := d.client.SearchAccount(ctx, in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (d dbStorage) CreateTransaction(ctx context.Context, in *db.CreateTransactionRequest) (*db.CreateTransactionResponse, error) {
	res, err := d.client.CreateTransaction(ctx, in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (d dbStorage) DeleteTransaction(ctx context.Context, in *db.DeleteTransactionRequest) (emptypb.Empty, error) {
	res, err := d.client.DeleteTransaction(ctx, in)
	if err != nil {
		return emptypb.Empty{}, err
	}
	return *res, nil
}

func (d dbStorage) UpdateTransaction(ctx context.Context, in *db.UpdateTransactionRequest) (emptypb.Empty, error) {
	res, err := d.client.UpdateTransaction(ctx, in)
	if err != nil {
		return emptypb.Empty{}, err
	}
	return *res, nil
}

func (d dbStorage) GetTransaction(ctx context.Context, in *db.GetTransactionRequest) (*db.Transaction, error) {
	res, err := d.client.GetTransaction(ctx, in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (d dbStorage) GetTransactionByAccount(ctx context.Context, in *db.GetTransactionByAccountRequest) (*db.GetTransactionByAccountResponse, error) {
	res, err := d.client.GetTransactionByAccount(ctx, in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (d dbStorage) CreateTrade(ctx context.Context, in *db.CreateTradeRequest) (*db.CreateTradeResponse, error) {
	res, err := d.client.CreateTrade(ctx, in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (d dbStorage) DeleteTrade(ctx context.Context, in *db.DeleteTradeRequest) (emptypb.Empty, error) {
	_, err := d.client.DeleteTrade(ctx, in)
	if err != nil {
		return emptypb.Empty{}, err
	}
	return emptypb.Empty{}, nil
}

func (d dbStorage) UpdateTrade(ctx context.Context, in *db.UpdateTradeRequest) (emptypb.Empty, error) {
	_, err := d.client.UpdateTrade(ctx, in)
	if err != nil {
		return emptypb.Empty{}, err
	}
	return emptypb.Empty{}, nil
}

func (d dbStorage) GetTrade(ctx context.Context, in *db.GetTradeRequest) (*db.Trade, error) {
	res, err := d.client.GetTrade(ctx, in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (d dbStorage) GetTradeByAccount(ctx context.Context, in *db.GetTradeByAccountRequest) (*db.GetTradeByAccountResponse, error) {
	res, err := d.client.GetTradeByAccount(ctx, in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func NewDBStorage(client db.DbServiceClient) ports.DBStorage {
	return &dbStorage{
		client: client,
	}
}
