package grpc

import (
	"context"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/internal/core/entity"
)

type Client struct {
	client *string
}

func (c Client) Create(ctx context.Context, data entity.User) error {
	return nil
}

func Update(ctx context.Context, user entity.User) error {
	return nil
}

func FindByID(ctx context.Context, uuid string) (*entity.User, error) {
	panic("not implemented")
}

func FindByEmail(ctx context.Context, emailAddress string) (*entity.User, error) {
	panic("not implemented")
}

func NewClient() *Client {
	return &Client{}
}
