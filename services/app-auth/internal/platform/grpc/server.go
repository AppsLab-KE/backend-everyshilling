package appdb

import (
	"context"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/internal/core/entity"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/internal/core/ports"
)

type Client struct {
	client *string
}

func (c Client) Create(ctx context.Context, data entity.User) error {
	return nil
}

var _ ports.Storage = (*Client)(nil)

func NewClient() *Client {
	return &Client{}
}
