package inventory

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type ProductVariants struct {
	ID              uuid.UUID       `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	ProductID       uuid.UUID       `gorm:"type:uuid;not null" json:"product_id"`
	OutletID        uuid.UUID       `gorm:"type:uuid;not null" json:"outlet_id"`
	VariantName     string          `gorm:"type:varchar(200)" json:"variant_name"`
	VariantValue    string          `gorm:"type:varchar(200)" json:"variant_value"`
	AdditionalPrice decimal.Decimal `gorm:"type:decimal(10,2);default:0" json:"additional_price"`
	StockQuantity   int             `gorm:"default:0" json:"stock_quantity"`
	MinimumStock    int             `gorm:"default:0" json:"minimum_stock"`
	CreatedAt       time.Time       `json:"created_at"`
	UpdatedAt       time.Time       `json:"updated_at"`

	// Relasi
	Product Products `gorm:"foreignKey:ProductID" json:"product,omitempty"`
}
