package customers

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type LoyaltyTransactions struct {
	ID            uuid.UUID       `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	CustomerID    uuid.UUID       `gorm:"type:uuid;not null" json:"customer_id"`
	TransactionID uuid.UUID       `gorm:"type:uuid;not null" json:"transaction_id"`
	PointChange   decimal.Decimal `gorm:"type:numeric(15,2);not null" json:"point_change"`
	Description   string          `gorm:"type:text" json:"description"`
	CreatedAt     time.Time       `json:"created_at"`

	// Relasi
	Customer Customers `gorm:"foreignKey:CustomerID" json:"customer,omitempty"`
}
