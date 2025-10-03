package auth

import (
	"pos_api_go/internal/models/core"
	"time"

	"github.com/google/uuid"
)

type UserOutletRoles struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	UserID    uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`
	OutletID  uuid.UUID `gorm:"type:uuid;not null" json:"outlet_id"`
	RoleID    uuid.UUID `gorm:"type:uuid;not null" json:"role_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// Relasi - Disable auto foreign key creation
	User   Users        `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"user,omitempty"`
	Outlet core.Outlets `gorm:"foreignKey:OutletID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"outlet,omitempty"`
	Role   Roles        `gorm:"foreignKey:RoleID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"role,omitempty"`
}
