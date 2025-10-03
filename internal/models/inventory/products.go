package inventory

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Products struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	OutletID    uuid.UUID `gorm:"type:uuid"`
	CategoryID  uuid.UUID `gorm:"type:uuid"`
	Name        string    `gorm:"type:varchar(150);not null"`
	Description string    `gorm:"type:text"`
	ImageURL    string    `gorm:"type:text"`
	IsActive    bool      `gorm:"default:true"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
