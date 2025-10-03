package inventory

import (
	"pos_api_go/internal/models/auth"
	"pos_api_go/internal/models/core"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type InventoryMovements struct {
	ID               uuid.UUID       `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	ProductVariantID uuid.UUID       `gorm:"type:uuid;not null" json:"product_variant_id"`
	OutletID         uuid.UUID       `gorm:"type:uuid;not null" json:"outlet_id"`
	UserID           uuid.UUID       `gorm:"type:uuid;not null" json:"user_id"`
	Type             string          `gorm:"type:inventory_movement_type" json:"type"`
	QtyChange        decimal.Decimal `gorm:"type:decimal(15,2)" json:"qty_change"`
	Notes            string          `gorm:"type:text" json:"notes"`
	ReferenceID      *uuid.UUID      `gorm:"type:uuid" json:"reference_id"`
	CreatedAt        time.Time       `json:"created_at"`

	// Relasi
	ProductVariant ProductVariants `gorm:"foreignKey:ProductVariantID" json:"product_variant,omitempty"`
	Outlet         core.Outlets    `gorm:"foreignKey:OutletID" json:"outlet,omitempty"`
	User           auth.Users      `gorm:"foreignKey:UserID" json:"user,omitempty"`
}
