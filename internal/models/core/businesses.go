package core

import (
	"time"

	"github.com/google/uuid"
)

type Business struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name      string    `gorm:"type:varchar(150);not null"`
	OwnerID   uuid.UUID `gorm:"type:uuid"`
	Status    string    `gorm:"type:enum('active','suspended','closed','trial')"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
