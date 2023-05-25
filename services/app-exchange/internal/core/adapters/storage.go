package adapters

import (
	"context"
	"google.golang.org/grpc"
)

type CurrencyStorage interface {
	HealthCheck(ctx context.Context, req *currencyproto.DefaultRequest, opts ...grpc.CallOption) (*currencyproto.HealthResponse, error)
	//CreateCurrency(ctx context.Context, req *currencyproto.CreateCurrencyReq, opts ...grpc.CallOption) (*currencyproto.CreateCurrencyRes, error)
	//UpdateCurrency(ctx context.Context, req *currencyproto.UpdateCurrencyReq, opts ...grpc.CallOption) (*currencyproto.UpdateCurrencyRes, error)
	//GetCurrencyByField(ctx context.Context, req *currencyproto.GetByFieldReq, opts ...grpc.CallOption) (*currencyproto.GetByFieldRes, error)
	//GetPagedCurrencies(ctx context.Context, req *currencyproto.GetPagedCurrenciesReq, opts ...grpc.CallOption) (*currencyproto.GetPagedCurrenciesRes, error)
}
