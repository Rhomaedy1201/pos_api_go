package inventory

import (
	"pos_api_go/internal/models/core"
	"time"

	"github.com/google/uuid"
)

type Suppliers struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	OutletID  uuid.UUID `gorm:"type:uuid;not null" json:"outlet_id"`
	Name      string    `gorm:"type:varchar(150);not null" json:"name"`
	Phone     string    `gorm:"type:varchar(20)" json:"phone"`
	Email     string    `gorm:"type:varchar(100)" json:"email"`
	Address   string    `gorm:"type:text" json:"address"`
	Status    string    `gorm:"type:supplier_status;default:'active'" json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// Relasi
	Outlet core.Outlets `gorm:"foreignKey:OutletID" json:"outlet,omitempty"`
}
