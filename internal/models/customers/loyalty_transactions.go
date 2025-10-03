package customers

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type LoyaltyTransactions struct {
	Id            uuid.UUID       `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	CustomerId    uuid.UUID       `gorm:"type:uuid;not null" json:"customer_id"`
	TransactionId uuid.UUID       `gorm:"type:uuid;not null" json:"transaction_id"`
	PointChange   decimal.Decimal `gorm:"type:numeric(15,2);not null" json:"point_change"`
	Description   string          `gorm:"type:text" json:"description"`
	CreatedAt     string          `gorm:"type:timestamp;not null;default:current_timestamp" json:"created_at"`
}
