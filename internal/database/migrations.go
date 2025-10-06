package database

import (
	"pos_api_go/internal/database/migrations"

	"gorm.io/gorm"
)

// RunMigrations runs all database migrations
func RunMigrations(db *gorm.DB) error {
	migrator := migrations.NewMigrator(db)
	return migrator.Run()
}

// RollbackMigration rolls back to a specific migration
func RollbackMigration(db *gorm.DB, migrationID string) error {
	migrator := migrations.NewMigrator(db)
	return migrator.Rollback(migrationID)
}
