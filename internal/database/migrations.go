package database

import (
	"log"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

// RunMigrations menjalankan semua migrasi database
func RunMigrations(db *gorm.DB) {
	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		// Contoh migrasi - Anda bisa menambahkan migrasi sesuai kebutuhan
		// {
		//     ID: "202410020001",
		//     Migrate: func(tx *gorm.DB) error {
		//         // Definisi tabel atau perubahan schema di sini
		//         return nil
		//     },
		//     Rollback: func(tx *gorm.DB) error {
		//         // Rollback perubahan di sini
		//         return nil
		//     },
		// },
	})

	if err := m.Migrate(); err != nil {
		log.Fatalf("Could not migrate: %v", err)
	}

	log.Println("Migration completed successfully!")
}
