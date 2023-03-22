package appdb

import (
	"context"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/internal/models"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/internal/ports"
)

type Client struct {
	client *string
}

func (c Client) Create(ctx context.Context, data models.User) error {
	return nil
}

var _ ports.Storage = (*Client)(nil)

func NewClient() *Client {
	return &Client{}
}
