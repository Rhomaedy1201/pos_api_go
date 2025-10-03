package sales

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type TransactionItems struct {
	Id               uuid.UUID       `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	TransactionId    uuid.UUID       `gorm:"type:uuid;not null"`
	ProductVariantId uuid.UUID       `gorm:"type:uuid;not null"`
	Quantity         int             `gorm:"not null" json:"quantity"`
	PricePerUnit     decimal.Decimal `gorm:"type:numeric(15,2);not null" json:"price_per_unit"`
	DiscountPerUnit  decimal.Decimal `gorm:"type:numeric(15,2);not null" json:"discount_per_unit"`
	TotalPrice       decimal.Decimal `gorm:"type:numeric(15,2);not null" json:"total_price"`
}
