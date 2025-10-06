package sales

import (
	"pos_api_go/internal/models/auth"
	"pos_api_go/internal/models/core"
	"pos_api_go/internal/models/customers"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Transactions struct {
	ID              uuid.UUID       `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	OutletID        uuid.UUID       `gorm:"type:uuid;not null" json:"outlet_id"`
	UserID          uuid.UUID       `gorm:"type:uuid;not null" json:"user_id"`
	CustomerID      uuid.UUID       `gorm:"type:uuid" json:"customer_id"`
	InvoiceNumber   string          `gorm:"type:varchar(100);uniqueIndex;not null" json:"invoice_number"`
	SubTotal        decimal.Decimal `gorm:"type:numeric(15,2);not null" json:"sub_total"`
	TotalDiscount   decimal.Decimal `gorm:"type:numeric(15,2);not null" json:"total_discount"`
	TaxAmount       decimal.Decimal `gorm:"type:numeric(15,2);not null" json:"tax_amount"`
	GrandTotal      decimal.Decimal `gorm:"type:numeric(15,2);not null" json:"grand_total"`
	Status          string          `gorm:"type:transaction_status;default:'pending'" json:"status"`
	TransactionTime time.Time       `gorm:"not null" json:"transaction_time"`
	CreatedAt       time.Time       `json:"created_at"`
	UpdatedAt       time.Time       `json:"updated_at"`

	// Relasi
	Outlet           core.Outlets        `gorm:"foreignKey:OutletID" json:"outlet,omitempty"`
	User             auth.Users          `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Customer         customers.Customers `gorm:"foreignKey:CustomerID" json:"customer,omitempty"`
	TransactionItems []TransactionItems  `gorm:"foreignKey:TransactionID" json:"transaction_items,omitempty"`
	Payments         []Payments          `gorm:"foreignKey:TransactionID" json:"payments,omitempty"`
}
