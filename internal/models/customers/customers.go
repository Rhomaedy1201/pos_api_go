package customers

import (
	"pos_api_go/internal/models/core"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Customers struct {
	ID            uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	BusinessID    uuid.UUID      `gorm:"type:uuid;not null" json:"business_id"`
	Name          string         `gorm:"type:varchar(100);not null" json:"name"`
	PhoneNumber   string         `gorm:"type:varchar(15);uniqueIndex;not null" json:"phone_number"`
	Email         string         `gorm:"type:varchar(100);uniqueIndex" json:"email"`
	Address       string         `gorm:"type:text" json:"address"`
	LoyaltyPoints int            `gorm:"not null;default:0" json:"loyalty_points"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`

	// Relasi
	Business            core.Business         `gorm:"foreignKey:BusinessID" json:"business,omitempty"`
	LoyaltyTransactions []LoyaltyTransactions `gorm:"foreignKey:CustomerID" json:"loyalty_transactions,omitempty"`
}
