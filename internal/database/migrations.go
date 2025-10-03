package database

import (
	"log"
	"pos_api_go/internal/models/auth"
	"pos_api_go/internal/models/core"
	"pos_api_go/internal/models/customers"
	"pos_api_go/internal/models/inventory"
	"pos_api_go/internal/models/sales"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

// RunMigrations menjalankan semua migrasi database
func RunMigrations(db *gorm.DB) {
	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: "202410030001_create_enums",
			Migrate: func(tx *gorm.DB) error {
				// Enable UUID extension
				if err := tx.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"").Error; err != nil {
					return err
				}

				// Create ENUM types untuk PostgreSQL
				enums := []string{
					"DO $$ BEGIN CREATE TYPE business_status AS ENUM ('active', 'suspended', 'closed', 'trial'); EXCEPTION WHEN duplicate_object THEN null; END $$;",
					"DO $$ BEGIN CREATE TYPE outlet_status AS ENUM ('active', 'inactive'); EXCEPTION WHEN duplicate_object THEN null; END $$;",
					"DO $$ BEGIN CREATE TYPE category_status AS ENUM ('active', 'inactive'); EXCEPTION WHEN duplicate_object THEN null; END $$;",
					"DO $$ BEGIN CREATE TYPE supplier_status AS ENUM ('active', 'inactive'); EXCEPTION WHEN duplicate_object THEN null; END $$;",
					"DO $$ BEGIN CREATE TYPE inventory_movement_type AS ENUM ('sale', 'purchase', 'return', 'adjustment', 'transfer_in', 'transfer_out'); EXCEPTION WHEN duplicate_object THEN null; END $$;",
				}

				for _, enum := range enums {
					if err := tx.Exec(enum).Error; err != nil {
						log.Printf("Warning: Could not create enum: %v", err)
					}
				}

				log.Println("✅ ENUM types created successfully")
				return nil
			},
			Rollback: func(tx *gorm.DB) error {
				// Drop ENUM types
				enums := []string{
					"DROP TYPE IF EXISTS business_status CASCADE;",
					"DROP TYPE IF EXISTS outlet_status CASCADE;",
					"DROP TYPE IF EXISTS category_status CASCADE;",
					"DROP TYPE IF EXISTS supplier_status CASCADE;",
					"DROP TYPE IF EXISTS inventory_movement_type CASCADE;",
				}

				for _, enum := range enums {
					if err := tx.Exec(enum).Error; err != nil {
						log.Printf("Warning: Could not drop enum: %v", err)
					}
				}

				// Drop extension
				if err := tx.Exec("DROP EXTENSION IF EXISTS \"uuid-ossp\"").Error; err != nil {
					log.Printf("Warning: Could not drop UUID extension: %v", err)
				}

				return nil
			},
		},
		{
			ID: "202410030002_create_tables",
			Migrate: func(tx *gorm.DB) error {
				// Create tables in phases to avoid FK constraint issues

				// Phase 1: Independent tables (no foreign keys)
				phase1 := []interface{}{
					&auth.Users{},
					&auth.Roles{},
				}

				for _, model := range phase1 {
					if err := tx.AutoMigrate(model); err != nil {
						return err
					}
				}
				log.Println("✅ Phase 1: Users and Roles tables created")

				// Phase 2: Business and Outlets (depends on Users)
				phase2 := []interface{}{
					&core.Business{},
					&core.Outlets{},
				}

				for _, model := range phase2 {
					if err := tx.AutoMigrate(model); err != nil {
						return err
					}
				}
				log.Println("✅ Phase 2: Business and Outlets tables created")

				// Phase 3: Junction and Customer tables
				phase3 := []interface{}{
					&auth.UserOutletRoles{},
					&customers.Customers{},
				}

				for _, model := range phase3 {
					if err := tx.AutoMigrate(model); err != nil {
						return err
					}
				}
				log.Println("✅ Phase 3: Junction and Customer tables created")

				// Phase 4: Inventory tables
				phase4 := []interface{}{
					&inventory.Categories{},
					&inventory.Suppliers{},
					&inventory.Products{},
					&inventory.ProductVariants{},
					&inventory.InventoryMovements{},
				}

				for _, model := range phase4 {
					if err := tx.AutoMigrate(model); err != nil {
						return err
					}
				}
				log.Println("✅ Phase 4: Inventory tables created")

				// Phase 5: Sales tables
				phase5 := []interface{}{
					&sales.Transactions{},
					&sales.TransactionItems{},
					&sales.Payments{},
					&customers.LoyaltyTransactions{},
				}

				for _, model := range phase5 {
					if err := tx.AutoMigrate(model); err != nil {
						return err
					}
				}
				log.Println("✅ Phase 5: Sales and Loyalty tables created")
				log.Println("🎉 All tables created successfully with relationships")
				return nil
			},
			Rollback: func(tx *gorm.DB) error {
				// Drop tables in reverse order
				tables := []string{
					"loyalty_transactions",
					"payments",
					"transaction_items",
					"transactions",
					"inventory_movements",
					"product_variants",
					"products",
					"suppliers",
					"categories",
					"customers",
					"user_outlet_roles",
					"outlets",
					"businesses",
					"roles",
					"users",
				}

				for _, table := range tables {
					if err := tx.Migrator().DropTable(table); err != nil {
						log.Printf("Warning: Could not drop table %s: %v", table, err)
					}
				}

				// Drop extension
				if err := tx.Exec("DROP EXTENSION IF EXISTS \"uuid-ossp\"").Error; err != nil {
					log.Printf("Warning: Could not drop UUID extension: %v", err)
				}

				log.Println("✅ All tables dropped successfully")
				return nil
			},
		},
	})

	if err := m.Migrate(); err != nil {
		log.Fatalf("❌ Migration failed: %v", err)
	}

	log.Println("🎉 Migration completed successfully!")
}
