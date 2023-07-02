package platform

import (
	"fmt"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-exchange/config"
	"github.com/AppsLab-KE/be-go-gen-grpc/db"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func NewDBServiceClient(config config.DB) (db.DbServiceClient, error) {
	uri := fmt.Sprintf("%s:%d", config.Host, config.Port)
	conn, err := grpc.Dial(uri, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	log.Println("connected to app-db")
	client := db.NewDbServiceClient(conn)
	return client, nil
}
