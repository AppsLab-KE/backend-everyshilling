package ports

import (
	"context"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-authentication/internal/models"
)

type Storage interface {
	Create(ctx context.Context, data models.User) error
}
