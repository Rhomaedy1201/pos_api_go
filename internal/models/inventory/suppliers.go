package inventory

import (
	"time"

	"github.com/google/uuid"
)

type Suppliers struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	OutletID  uuid.UUID `gorm:"type:uuid"`
	Name      string    `gorm:"type:varchar(150);not null"`
	Phone     string    `gorm:"type:varchar(20)"`
	Email     string    `gorm:"type:varchar(100);nullable"`
	Address   string    `gorm:"type:text"`
	Status    string    `gorm:"type:enum('active','inactive')"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
