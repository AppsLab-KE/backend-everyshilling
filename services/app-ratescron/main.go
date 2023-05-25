package main

import (
	"github.com/AppsLab-KE/backend-everyshilling/services/app-ratescron/config"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-ratescron/internal/core/repository"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-ratescron/internal/core/storage"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-ratescron/internal/usecase"
	log "github.com/sirupsen/logrus"
	"os"
)

func main() {
	cfg, err := config.LoadFromEnv()

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	dbStorage, err := storage.NewDbStorage(cfg.DB)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	ratesRepository := repository.NewRatesAPIRepository(cfg.API)
	dbRepository := repository.NewDBRepository(dbStorage)

	uc := usecase.NewUsecase(dbRepository, ratesRepository)

	err = uc.FetchAndStoreRates()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
