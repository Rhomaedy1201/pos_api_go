package core

import (
	"time"

	"github.com/google/uuid"
)

type Business struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar(150);not null" json:"name"`
	OwnerID   uuid.UUID `gorm:"type:uuid;not null" json:"owner_id"`
	Status    string    `gorm:"type:enum('active','suspended','closed','trial');default:'trial'" json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// Relasi
	Outlets []Outlets `gorm:"foreignKey:BusinessID" json:"outlets,omitempty"`
}
