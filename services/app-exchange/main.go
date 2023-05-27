package main

import (
	"github.com/AppsLab-KE/backend-everyshilling/services/app-exchange/config"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-exchange/internal/routes/server"
	log "github.com/sirupsen/logrus"
)

func main() {
	cfg, err := config.LoadFromEnv()

	if err != nil {
		log.Panic("config error:", err)
	}

	grpcServer := server.NewServer(*cfg)
	grpcServer.Run()
}
