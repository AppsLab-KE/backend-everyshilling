package storage

import (
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/internal/core/ports"
	appdb "github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/internal/platform/app_db"
)

func New() (ports.Storage, error) {
	return appdb.NewClient(), nil
}
