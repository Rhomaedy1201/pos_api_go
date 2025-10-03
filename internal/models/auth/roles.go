package auth

import (
	"time"

	"github.com/google/uuid"
)

type Roles struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name      string    `gorm:"type:varchar(100);not null;uniqueIndex"`
	CreatedAt time.Time
}
