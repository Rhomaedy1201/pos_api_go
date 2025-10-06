package migrations

import (
	"pos_api_go/internal/models/auth"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func seedDefaultRolesMigration() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "202410030003_seed_default_roles",
		Migrate: func(tx *gorm.DB) error {
			roles := []auth.Roles{
				{Name: "Super Admin"},
				{Name: "Business Owner"},
				{Name: "Manager"},
				{Name: "Cashier"},
				{Name: "Staff"},
			}

			for _, role := range roles {
				var existingRole auth.Roles
				if err := tx.Where("name = ?", role.Name).First(&existingRole).Error; err != nil {
					if err == gorm.ErrRecordNotFound {
						if err := tx.Create(&role).Error; err != nil {
							return err
						}
					}
				}
			}
			return nil
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Where("name IN ?", []string{
				"Super Admin", "Business Owner", "Manager", "Cashier", "Staff",
			}).Delete(&auth.Roles{}).Error
		},
	}
}
