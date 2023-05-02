/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/AppsLab-KE/backend-everyshilling/services/app-db/config"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-db/internal/routes/server"
	log "github.com/sirupsen/logrus"
	_ "os"
)

func Execute() {
	cfg, err := config.LoadConfig()

	if err != nil {
		log.Panic("config error:", err)
	}

	grpcServer := server.NewServer(*cfg)
	grpcServer.Run()
}
