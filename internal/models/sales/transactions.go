package sales

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Transactions struct {
	Id              uuid.UUID       `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	OutletId        uuid.UUID       `gorm:"type:uuid;not null"`
	UserId          uuid.UUID       `gorm:"type:uuid;not null"`
	CustomerId      uuid.UUID       `gorm:"type:uuid;not null"`
	InvoiceNumber   string          `gorm:"type:varchar(100);uniqueIndex;not null" json:"invoice_number"`
	SubTotal        decimal.Decimal `gorm:"type:numeric(15,2);not null" json:"sub_total"`
	TotalDiscount   decimal.Decimal `gorm:"type:numeric(15,2);not null" json:"total_discount"`
	TaxAmount       decimal.Decimal `gorm:"type:numeric(15,2);not null" json:"tax_amount"`
	GrandTotal      decimal.Decimal `gorm:"type:numeric(15,2);not null" json:"grand_total"`
	Status          string          `gorm:"type:varchar(50);not null" json:"status"`
	TransactionTime string          `gorm:"type:timestamp;not null" json:"transaction_time"`
}
