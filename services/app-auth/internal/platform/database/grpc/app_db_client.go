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

func NewClient() *Client {
	return &Client{}
}
