package inventory

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type ProductVariants struct {
	ID        uuid.UUID       `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	ProductID uuid.UUID       `gorm:"type:uuid"`
	Sku       string          `gorm:"type:varchar(100);unique;not null"`
	Name      string          `gorm:"type:varchar(150);not null"`
	CostPrice decimal.Decimal `gorm:"type:decimal(15,2);default:0"`
	SellPrice decimal.Decimal `gorm:"type:decimal(15,2);default:0"`
	StockQty  int             `gorm:"default:0"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
