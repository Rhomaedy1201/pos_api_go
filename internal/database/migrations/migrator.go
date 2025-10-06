package migrations

import (
	"log"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

type Migrator struct {
	db *gorm.DB
}

func NewMigrator(db *gorm.DB) *Migrator {
	return &Migrator{db: db}
}

func (m *Migrator) Run() error {
	migrator := gormigrate.New(m.db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		initialSchemaMigration(),
		addIndexesMigration(),
		seedDefaultRolesMigration(),
	})

	if err := migrator.Migrate(); err != nil {
		log.Printf("Migration failed: %v", err)
		return err
	}

	log.Println("✅ Migration completed successfully!")
	return nil
}

func (m *Migrator) Rollback(migrationID string) error {
	migrator := gormigrate.New(m.db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		initialSchemaMigration(),
		addIndexesMigration(),
		seedDefaultRolesMigration(),
	})

	if err := migrator.RollbackTo(migrationID); err != nil {
		log.Printf("Rollback failed: %v", err)
		return err
	}

	log.Printf("✅ Rollback to %s completed successfully!", migrationID)
	return nil
}
