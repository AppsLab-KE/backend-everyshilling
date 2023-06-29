package ports

import (
	"context"
	"github.com/AppsLab-KE/be-go-gen-grpc/db"
	"google.golang.org/protobuf/types/known/emptypb"
)

type TradeRepository interface {
	CreateTrade(ctx context.Context, in *db.CreateTradeRequest) (*db.CreateTradeResponse, error)
	DeleteTrade(ctx context.Context, in *db.DeleteTradeRequest) (emptypb.Empty, error)
	UpdateTrade(ctx context.Context, in *db.UpdateTradeRequest) (emptypb.Empty, error)
	GetTrade(ctx context.Context, in *db.GetTradeRequest) (*db.Trade, error)
	GetTradeByAccount(ctx context.Context, in *db.GetTradeByAccountRequest) (*db.GetTradeByAccountResponse, error)
}

type TradeService interface {
	CreateTrade(ctx context.Context, in *db.CreateTradeRequest) (*db.CreateTradeResponse, error)
	DeleteTrade(ctx context.Context, in *db.DeleteTradeRequest) (emptypb.Empty, error)
	UpdateTrade(ctx context.Context, in *db.UpdateTradeRequest) (emptypb.Empty, error)
	GetTrade(ctx context.Context, in *db.GetTradeRequest) (*db.Trade, error)
	GetTradeByAccount(ctx context.Context, in *db.GetTradeByAccountRequest) (*db.GetTradeByAccountResponse, error)
}
