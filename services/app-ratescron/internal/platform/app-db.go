package platform

import (
	"fmt"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-ratescron/config"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-ratescron/internal/core/ports"
	"github.com/AppsLab-KE/be-go-gen-grpc/db"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var _ ports.DBStorage = (db.DbServiceClient)(nil)

func NewDBServiceClient(config config.DB) (db.DbServiceClient, error) {
	uri := fmt.Sprintf("%s:%d", config.Host, config.Port)
	conn, err := grpc.Dial(uri, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	log.Info("connected to app-db")
	client := db.NewDbServiceClient(conn)
	return client, nil
}
