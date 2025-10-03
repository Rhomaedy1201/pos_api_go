package inventory

import (
	"pos_api_go/internal/models/core"
	"time"

	"github.com/google/uuid"
)

type Categories struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	OutletID  uuid.UUID `gorm:"type:uuid;not null" json:"outlet_id"`
	Name      string    `gorm:"type:varchar(150);not null" json:"name"`
	Status    string    `gorm:"type:category_status;default:'active'" json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// Relasi
	Outlet   core.Outlets `gorm:"foreignKey:OutletID" json:"outlet,omitempty"`
	Products []Products   `gorm:"foreignKey:CategoryID" json:"products,omitempty"`
}
