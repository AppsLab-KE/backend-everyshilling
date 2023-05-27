package repository

import (
	"context"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-ratescron/internal/core/ports"
	"github.com/AppsLab-KE/be-go-gen-grpc/db"
)

type dbRepo struct {
	storage ports.DBStorage
}

func (d dbRepo) CreateConversionRate(ctx context.Context, in *db.CreateConversionRateRequest) (*db.CreateConversionRateResponse, error) {
	return d.storage.CreateConversionRate(ctx, in)
}

func (d dbRepo) ReadConversionRate(ctx context.Context, in *db.ReadConversionRateRequest) (*db.ReadConversionRateResponse, error) {
	return d.storage.ReadConversionRate(ctx, in)
}

func (d dbRepo) UpdateConversionRate(ctx context.Context, in *db.UpdateConversionRateRequest) (*db.UpdateConversionRateResponse, error) {
	return d.storage.UpdateConversionRate(ctx, in)
}

func (d dbRepo) DeleteConversionRate(ctx context.Context, in *db.DeleteConversionRateRequest) (*db.DeleteConversionRateResponse, error) {
	return d.storage.DeleteConversionRate(ctx, in)
}

func NewDBRepository(storage ports.DBStorage) ports.DBRepository {
	return &dbRepo{storage: storage}
}
