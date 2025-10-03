package customers

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type customers struct {
	Id            uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	BusinessId    uuid.UUID      `gorm:"type:uuid;not null" json:"business_id"`
	Name          string         `gorm:"type:varchar(100);not null" json:"name"`
	PhoneNumber   string         `gorm:"type:varchar(15);uniqueIndex;not null" json:"phone_number"`
	Email         string         `gorm:"type:varchar(100);uniqueIndex" json:"email"`
	Address       string         `gorm:"type:text" json:"address"`
	LoyaltyPoints int            `gorm:"not null;default:0" json:"loyalty_points"`
	CreatedAt     string         `gorm:"type:timestamp;not null;default:current_timestamp" json:"created_at"`
	UpdatedAt     string         `gorm:"type:timestamp;not null;default:current_timestamp" json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
}
