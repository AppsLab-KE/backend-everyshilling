package repositories

import (
	"context"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-exchange/internal/core/ports"
	"github.com/AppsLab-KE/be-go-gen-grpc/db"
	"google.golang.org/protobuf/types/known/emptypb"
)

type tradeRepository struct {
	dbStorage ports.DBStorage
}

func (t tradeRepository) CreateTrade(ctx context.Context, in *db.CreateTradeRequest) (*db.CreateTradeResponse, error) {
	res, err := t.dbStorage.CreateTrade(ctx, in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (t tradeRepository) DeleteTrade(ctx context.Context, in *db.DeleteTradeRequest) (emptypb.Empty, error) {
	res, err := t.dbStorage.DeleteTrade(ctx, in)
	if err != nil {
		return emptypb.Empty{}, err
	}
	return res, nil
}

func (t tradeRepository) UpdateTrade(ctx context.Context, in *db.UpdateTradeRequest) (emptypb.Empty, error) {
	res, err := t.dbStorage.UpdateTrade(ctx, in)
	if err != nil {
		return emptypb.Empty{}, err
	}
	return res, nil
}

func (t tradeRepository) GetTrade(ctx context.Context, in *db.GetTradeRequest) (*db.Trade, error) {
	res, err := t.dbStorage.GetTrade(ctx, in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (t tradeRepository) GetTradeByAccount(ctx context.Context, in *db.GetTradeByAccountRequest) (*db.GetTradeByAccountResponse, error) {
	res, err := t.dbStorage.GetTradeByAccount(ctx, in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func NewTradeRepository(dbStorage ports.DBStorage) ports.TradeRepository {
	return &tradeRepository{
		dbStorage: dbStorage,
	}
}
