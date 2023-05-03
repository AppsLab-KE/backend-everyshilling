package storage

import (
	"github.com/AppsLab-KE/backend-everyshilling/services/app-db/internal/core/ports"
)

type userCacheStorage struct {
}

func NewUserCacheStorage() ports.UserCache {
	return &userCacheStorage{}
}
