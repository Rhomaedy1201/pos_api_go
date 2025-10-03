package sales

import (
	"pos_api_go/internal/models/inventory"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type TransactionItems struct {
	ID               uuid.UUID       `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	TransactionID    uuid.UUID       `gorm:"type:uuid;not null" json:"transaction_id"`
	ProductVariantID uuid.UUID       `gorm:"type:uuid;not null" json:"product_variant_id"`
	Quantity         int             `gorm:"not null" json:"quantity"`
	PricePerUnit     decimal.Decimal `gorm:"type:numeric(15,2);not null" json:"price_per_unit"`
	DiscountPerUnit  decimal.Decimal `gorm:"type:numeric(15,2);not null" json:"discount_per_unit"`
	TotalPrice       decimal.Decimal `gorm:"type:numeric(15,2);not null" json:"total_price"`
	CreatedAt        time.Time       `json:"created_at"`
	UpdatedAt        time.Time       `json:"updated_at"`

	// Relasi
	Transaction    Transactions              `gorm:"foreignKey:TransactionID" json:"transaction,omitempty"`
	ProductVariant inventory.ProductVariants `gorm:"foreignKey:ProductVariantID" json:"product_variant,omitempty"`
}
