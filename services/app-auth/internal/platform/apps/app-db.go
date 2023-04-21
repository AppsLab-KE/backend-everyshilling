package apps

import (
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/config"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/internal/core/adapters"
	"github.com/AppsLab-KE/be-go-gen-grpc/db"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var _ adapters.DBStorage = (db.DbServiceClient)(nil)

func NewDBServiceClient(config config.DatabaseService) (db.DbServiceClient, error) {
	conn, err := grpc.Dial(config.Host+":"+config.Port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	client := db.NewDbServiceClient(conn)
	return client, nil
}
