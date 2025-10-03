package auth

import (
	"time"

	"github.com/google/uuid"
)

type Roles struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar(100);not null;uniqueIndex" json:"name"`
	CreatedAt time.Time `json:"created_at"`

	// Relasi
	UserOutletRoles []UserOutletRoles `gorm:"foreignKey:RoleID" json:"user_outlet_roles,omitempty"`
}
