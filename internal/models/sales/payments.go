package sales

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Payments struct {
	ID              uuid.UUID       `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	TransactionID   uuid.UUID       `gorm:"type:uuid;not null" json:"transaction_id"`
	PaymentMethod   string          `gorm:"type:varchar(50);not null" json:"payment_method"`
	AmountPaid      decimal.Decimal `gorm:"type:numeric(15,2);not null" json:"amount_paid"`
	ReferenceNumber string          `gorm:"type:varchar(100)" json:"reference_number"`
	PaymentTime     time.Time       `gorm:"not null" json:"payment_time"`
	CreatedAt       time.Time       `json:"created_at"`
	UpdatedAt       time.Time       `json:"updated_at"`

	// Relasi
	Transaction Transactions `gorm:"foreignKey:TransactionID" json:"transaction,omitempty"`
}
