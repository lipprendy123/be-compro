package model

import (
	"time"

	"gorm.io/gorm"
)

type PorotfolioDetail struct {
	ID                  int64 `gorm:"id,primaryKey"`
	PorotfolioSectionID int64
	Category            string
	ClientName          string
	ProjectDate         time.Time
	ProjectUrl          *string
	Title               string
	Description         string
	CreatedAt           time.Time
	UpdatedAt           *time.Time
	DeletedAt           gorm.DeletedAt `gorm:"index"`
}
