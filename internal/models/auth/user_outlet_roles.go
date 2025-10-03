package auth

import (
	"time"

	"github.com/google/uuid"
)

type UserOutletRoles struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	UserID    uuid.UUID `gorm:"type:uuid"`
	OutletID  uuid.UUID `gorm:"type:uuid"`
	RoleID    uuid.UUID `gorm:"type:uuid"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
