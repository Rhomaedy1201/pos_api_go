package core

import (
	"time"

	"github.com/google/uuid"
)

type Outlets struct {
	ID         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	BusinessID uuid.UUID `gorm:"type:uuid;not null" json:"business_id"`
	Name       string    `gorm:"type:varchar(150);not null" json:"name"`
	Address    string    `gorm:"type:text" json:"address"`
	Status     string    `gorm:"type:outlet_status;default:'active'" json:"status"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`

	// Relasi
	Business Business `gorm:"foreignKey:BusinessID" json:"business,omitempty"`
}
