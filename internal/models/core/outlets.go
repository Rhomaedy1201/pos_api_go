package core

import (
	"time"

	"github.com/google/uuid"
)

type Outlets struct {
	ID         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	BusinessID uuid.UUID `gorm:"type:uuid"`
	Name       string    `gorm:"type:varchar(150);not null"`
	Address    string    `gorm:"type:text"`
	Status     string    `gorm:"type:enum('active','inactive')"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
