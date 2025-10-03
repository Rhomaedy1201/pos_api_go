package auth

import (
	"time"

	"github.com/google/uuid"
)

type Users struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Name        string    `gorm:"type:varchar(100);not null" json:"name"`
	Email       string    `gorm:"type:varchar(100);uniqueIndex;not null" json:"email"`
	Password    string    `gorm:"type:varchar(255);not null" json:"-"`
	PhoneNumber string    `gorm:"type:varchar(130);uniqueIndex;not null" json:"phone_number"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`

	// Relasi
	UserOutletRoles []UserOutletRoles `gorm:"foreignKey:UserID" json:"user_outlet_roles,omitempty"`
}
