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

func RunMigrations(db *gorm.DB) {
	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: "202410030001_initial_schema",
			Migrate: func(tx *gorm.DB) error {
				// Enable UUID extension
				tx.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")

				// Create ENUM types
				enums := []string{
					"CREATE TYPE IF NOT EXISTS business_status AS ENUM ('active', 'suspended', 'closed', 'trial')",
					"CREATE TYPE IF NOT EXISTS outlet_status AS ENUM ('active', 'inactive')",
					"CREATE TYPE IF NOT EXISTS category_status AS ENUM ('active', 'inactive')",
					"CREATE TYPE IF NOT EXISTS supplier_status AS ENUM ('active', 'inactive')",
					"CREATE TYPE IF NOT EXISTS inventory_movement_type AS ENUM ('sale', 'purchase', 'return', 'adjustment', 'transfer_in', 'transfer_out')",
				}

				for _, enum := range enums {
					tx.Exec(enum)
				}

				// Auto migrate all models at once
				// GORM will figure out the dependency order
				return tx.AutoMigrate(
					&auth.Users{},
					&auth.Roles{},
					&core.Business{},
					&core.Outlets{},
					&auth.UserOutletRoles{},
					&customers.Customers{},
					&inventory.Categories{},
					&inventory.Suppliers{},
					&inventory.Products{},
					&inventory.ProductVariants{},
					&inventory.InventoryMovements{},
					&sales.Transactions{},
					&sales.TransactionItems{},
					&sales.Payments{},
					&customers.LoyaltyTransactions{},
				)
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable(
					"loyalty_transactions", "payments", "transaction_items", "transactions",
					"inventory_movements", "product_variants", "products", "suppliers", "categories",
					"customers", "user_outlet_roles", "outlets", "businesses", "roles", "users",
				)
			},
		},
	})

	if err := m.Migrate(); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}
	log.Println("✅ Migration completed!")
}
