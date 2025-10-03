package inventory

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type ProductVariants struct {
	ID        uuid.UUID       `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	ProductID uuid.UUID       `gorm:"type:uuid;not null" json:"product_id"`
	Sku       string          `gorm:"type:varchar(100);unique;not null" json:"sku"`
	Name      string          `gorm:"type:varchar(150);not null" json:"name"`
	CostPrice decimal.Decimal `gorm:"type:decimal(15,2);default:0" json:"cost_price"`
	SellPrice decimal.Decimal `gorm:"type:decimal(15,2);default:0" json:"sell_price"`
	StockQty  int             `gorm:"default:0" json:"stock_qty"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`

	// Relasi
	Product Products `gorm:"foreignKey:ProductID" json:"product,omitempty"`
}
