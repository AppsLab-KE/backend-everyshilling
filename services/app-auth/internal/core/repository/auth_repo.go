package repository

import "github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/internal/core/storage"

type DefaultAuthAppDbRepo struct {
	client storage.Db
}

func NewDefaultAuthAppDbRepo(db storage.Db) DefaultAuthAppDbRepo {
}
