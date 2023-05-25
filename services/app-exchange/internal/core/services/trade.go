package services

import (
	"context"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-exchange/internal/core/ports"
	"github.com/AppsLab-KE/be-go-gen-grpc/db"
	"google.golang.org/protobuf/types/known/emptypb"
)

type tradeService struct {
	tradeRepository ports.TradeRepository
}

func (t tradeService) CreateTrade(ctx context.Context, in *db.CreateTradeRequest) (*db.CreateTradeResponse, error) {

}

func (t tradeService) DeleteTrade(ctx context.Context, in *db.DeleteTradeRequest) (emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (t tradeService) UpdateTrade(ctx context.Context, in *db.UpdateTradeRequest) (emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (t tradeService) GetTrade(ctx context.Context, in *db.GetTradeRequest) (*db.Trade, error) {
	//TODO implement me
	panic("implement me")
}

func (t tradeService) GetTradeByAccount(ctx context.Context, in *db.GetTradeByAccountRequest) (*db.GetTradeByAccountResponse, error) {
	//TODO implement me
	panic("implement me")
}

func NewTradeService(tradeRepository ports.TradeRepository) ports.TradeService {
	return &tradeService{
		tradeRepository: tradeRepository,
	}
}
