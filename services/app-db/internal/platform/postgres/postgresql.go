package postgres

import (
	"fmt"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-db/config"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-db/internal/core/models"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewClient(postgresConfig config.Postgres) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s",
		postgresConfig.Host,
		postgresConfig.User,
		postgresConfig.Password,
		postgresConfig.Database,
		postgresConfig.Port,
		postgresConfig.TimeZone,
	)

	log.Info("database setup: dialing")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	log.Info("database setup: migrations")
	// automigrate
	err = db.AutoMigrate(&models.User{}, &models.ConversionRate{})

	if err != nil {
		return nil, err
	}

	log.Info("database setup: success")

	return db, err
}
