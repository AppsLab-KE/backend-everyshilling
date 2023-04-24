package models

import (
	"database/sql"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Model struct {
	ID          uuid.UUID `gorm:"primaryKey;size:36"`
	ActivatedAt sql.NullTime
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

func (b *Model) BeforeCreate(db *gorm.DB) error {
	b.ID = uuid.New()
	return nil
}
