package seeders

import (
	"pos_api_go/internal/models/auth"

	"gorm.io/gorm"
)

func SeedDefaultRoles(db *gorm.DB) error {
	roles := []auth.Roles{
		{Name: "Super Admin"},
		{Name: "Business Owner"},
		{Name: "Manager"},
		{Name: "Cashier"},
		{Name: "Staff"},
	}

	for _, role := range roles {
		var existingRole auth.Roles
		if err := db.Where("name = ?", role.Name).First(&existingRole).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				if err := db.Create(&role).Error; err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func SeedAllDefaultData(db *gorm.DB) error {
	if err := SeedDefaultRoles(db); err != nil {
		return err
	}

	// TODO: Add more seeders here
	// SeedDefaultCategories(db)
	// SeedDefaultBusinessSettings(db)

	return nil
}
