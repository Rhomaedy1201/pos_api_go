package sales

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Payemts struct {
	Id             uuid.UUID       `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	TransactionId  uuid.UUID       `gorm:"type:uuid;not null"`
	PaymentMethod  string          `gorm:"type:varchar(50);not null" json:"payment_method"`
	AmountPaid     decimal.Decimal `gorm:"type:numeric(15,2);not null" json:"amount_paid"`
	RefrenceNumber string          `gorm:"type:varchar(100)" json:"reference_number" note:"nomor referensi kartu kredit/debit, jika ada"`
	PaymentTime    string          `gorm:"type:timestamp;not null" json:"payment_time"`
}
