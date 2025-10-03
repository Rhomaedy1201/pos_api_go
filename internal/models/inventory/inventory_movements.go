package inventory

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type InventoryMovements struct {
	ID               uuid.UUID       `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	ProductVariantID uuid.UUID       `gorm:"type:uuid"`
	OutletID         uuid.UUID       `gorm:"type:uuid"`
	UserID           uuid.UUID       `gorm:"type:uuid"`
	Type             string          `gorm:"type:enum('sale','purchase','return','adjustment','transfer_in','transfer_out')"`
	QtyChange        decimal.Decimal `gorm:"type:decimal(15,2)"`
	Notes            string          `gorm:"type:text"`
	RefrenceID       uuid.UUID       `gorm:"type:uuid;nullable"`
	CreatedAt        time.Time
}
