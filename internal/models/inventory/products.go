package inventory

import (
	"pos_api_go/internal/models/core"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Products struct {
	ID          uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	OutletID    uuid.UUID      `gorm:"type:uuid;not null" json:"outlet_id"`
	CategoryID  uuid.UUID      `gorm:"type:uuid;not null" json:"category_id"`
	Name        string         `gorm:"type:varchar(150);not null" json:"name"`
	Description string         `gorm:"type:text" json:"description"`
	ImageURL    string         `gorm:"type:text" json:"image_url"`
	IsActive    bool           `gorm:"default:true" json:"is_active"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`

	// Relasi
	Outlet          core.Outlets      `gorm:"foreignKey:OutletID" json:"outlet,omitempty"`
	Category        Categories        `gorm:"foreignKey:CategoryID" json:"category,omitempty"`
	ProductVariants []ProductVariants `gorm:"foreignKey:ProductID" json:"product_variants,omitempty"`
}
