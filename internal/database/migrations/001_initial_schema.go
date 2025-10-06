package migrations

import (
	"pos_api_go/internal/models/auth"
	"pos_api_go/internal/models/core"
	"pos_api_go/internal/models/customers"
	"pos_api_go/internal/models/inventory"
	"pos_api_go/internal/models/sales"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func initialSchemaMigration() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "202410030001_initial_schema",
		Migrate: func(tx *gorm.DB) error {
			// Enable extensions
			if err := enableExtensions(tx); err != nil {
				return err
			}

			// Create custom types
			if err := createEnumTypes(tx); err != nil {
				return err
			}

			// Create tables using AutoMigrate
			return createTables(tx)
		},
		Rollback: func(tx *gorm.DB) error {
			return rollbackInitialSchema(tx)
		},
	}
}

func enableExtensions(tx *gorm.DB) error {
	extensions := []string{
		"CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"",
		"CREATE EXTENSION IF NOT EXISTS \"pgcrypto\"", // for gen_random_uuid()
	}

	for _, ext := range extensions {
		if err := tx.Exec(ext).Error; err != nil {
			return err
		}
	}
	return nil
}

func createEnumTypes(tx *gorm.DB) error {
	enums := []struct {
		name   string
		values string
	}{
		{"business_status", "('active', 'suspended', 'closed', 'trial')"},
		{"outlet_status", "('active', 'inactive')"},
		{"category_status", "('active', 'inactive')"},
		{"supplier_status", "('active', 'inactive')"},
		{"inventory_movement_type", "('sale', 'purchase', 'return', 'adjustment', 'transfer_in', 'transfer_out')"},
		{"transaction_status", "('pending', 'completed', 'cancelled', 'refunded')"},
		{"payment_method", "('cash', 'card', 'digital_wallet', 'bank_transfer')"},
		{"user_status", "('active', 'inactive', 'suspended')"},
	}

	for _, enum := range enums {
		// Check if type exists first
		var exists bool
		checkQuery := `SELECT EXISTS (SELECT 1 FROM pg_type WHERE typname = ?)`
		if err := tx.Raw(checkQuery, enum.name).Scan(&exists).Error; err != nil {
			return err
		}

		// Create type if it doesn't exist
		if !exists {
			createQuery := `CREATE TYPE ` + enum.name + ` AS ENUM ` + enum.values
			if err := tx.Exec(createQuery).Error; err != nil {
				return err
			}
		}
	}
	return nil
}

func createTables(tx *gorm.DB) error {
	// Create tables in dependency order
	models := []interface{}{
		&auth.Roles{},
		&core.Business{},
		&auth.Users{},
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
	}

	return tx.AutoMigrate(models...)
}

func rollbackInitialSchema(tx *gorm.DB) error {
	tables := []string{
		"loyalty_transactions", "payments", "transaction_items", "transactions",
		"inventory_movements", "product_variants", "products", "suppliers",
		"categories", "customers", "user_outlet_roles", "outlets",
		"businesses", "roles", "users",
	}

	// Drop tables
	for _, table := range tables {
		if err := tx.Migrator().DropTable(table); err != nil {
			return err
		}
	}

	// Drop enums
	enums := []string{
		"DROP TYPE IF EXISTS user_status",
		"DROP TYPE IF EXISTS payment_method",
		"DROP TYPE IF EXISTS transaction_status",
		"DROP TYPE IF EXISTS inventory_movement_type",
		"DROP TYPE IF EXISTS supplier_status",
		"DROP TYPE IF EXISTS category_status",
		"DROP TYPE IF EXISTS outlet_status",
		"DROP TYPE IF EXISTS business_status",
	}

	for _, enum := range enums {
		tx.Exec(enum)
	}

	return nil
}
