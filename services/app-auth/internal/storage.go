package internal

import (
	appdb "github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/internal/platforms/app_db"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/internal/ports"
)

func New() (ports.Storage, error) {
	return appdb.NewClient(), nil
}
