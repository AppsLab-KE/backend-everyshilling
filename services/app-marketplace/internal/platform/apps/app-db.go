package apps

import (
	"github.com/AppsLab-KE/backend-everyshilling/services/app-marketplace/config"
	"github.com/AppsLab-KE/be-go-gen-grpc/db"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewDBServiceClient(config config.DatabaseService) (db.DbServiceClient, error) {
	conn, err := grpc.Dial(config.Host+":"+config.Port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	log.Info("connected to app-db")
	client := db.NewDbServiceClient(conn)
	return client, nil
}
