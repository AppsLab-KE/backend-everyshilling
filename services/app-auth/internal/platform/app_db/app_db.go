package appdb

import "github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/internal/core/storage"

type Client struct {
}

func (c *Client) CreateOtp(data *interface{}) interface{} {
	//TODO implement me
	panic("implement me")
}

var _ storage.Db = (*Client)(nil)

func NewClient(cfg) Client {

}
