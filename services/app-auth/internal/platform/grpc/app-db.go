package appdb

import (
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/config"
	"github.com/AppsLab-KE/be-go-gen-grpc/db"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewDBServiceClient(config config.DatabaseService) (*db.DbServiceClient, error) {
	conn, err := grpc.Dial(config.Host+":"+config.Port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	client := db.NewDbServiceClient(conn)
	return &client, nil
}
