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
			ID: "202410030001",
			Migrate: func(tx *gorm.DB) error {
				// Enable UUID extension
				if err := tx.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"").Error; err != nil {
					return err
				}

				// Core tables - Business first
				if err := tx.AutoMigrate(&core.Business{}); err != nil {
					return err
				}

				// Outlets depend on Business
				if err := tx.AutoMigrate(&core.Outlets{}); err != nil {
					return err
				}

				// Auth tables
				if err := tx.AutoMigrate(&auth.Users{}); err != nil {
					return err
				}

				if err := tx.AutoMigrate(&auth.Roles{}); err != nil {
					return err
				}

				if err := tx.AutoMigrate(&auth.UserOutletRoles{}); err != nil {
					return err
				}

				// Customer tables
				if err := tx.AutoMigrate(&customers.Customers{}); err != nil {
					return err
				}

				// Inventory tables
				if err := tx.AutoMigrate(&inventory.Categories{}); err != nil {
					return err
				}

				if err := tx.AutoMigrate(&inventory.Products{}); err != nil {
					return err
				}

				if err := tx.AutoMigrate(&inventory.ProductVariants{}); err != nil {
					return err
				}

				if err := tx.AutoMigrate(&inventory.Suppliers{}); err != nil {
					return err
				}

				if err := tx.AutoMigrate(&inventory.InventoryMovements{}); err != nil {
					return err
				}

				// Sales tables
				if err := tx.AutoMigrate(&sales.Transactions{}); err != nil {
					return err
				}

				if err := tx.AutoMigrate(&sales.TransactionItems{}); err != nil {
					return err
				}

				if err := tx.AutoMigrate(&sales.Payments{}); err != nil {
					return err
				}

				if err := tx.AutoMigrate(&customers.LoyaltyTransactions{}); err != nil {
					return err
				}

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
					"suppliers",
					"product_variants",
					"products",
					"categories",
					"customers",
					"user_outlet_roles",
					"roles",
					"users",
					"outlets",
					"businesses",
				}

				for _, table := range tables {
					if err := tx.Migrator().DropTable(table); err != nil {
						return err
					}
				}

				return nil
			},
		},
		{
			ID: "202410030002",
			Migrate: func(tx *gorm.DB) error {
				// Add foreign key constraints

				// Business -> Users (owner)
				if err := tx.Exec("ALTER TABLE businesses ADD CONSTRAINT fk_businesses_owner FOREIGN KEY (owner_id) REFERENCES users(id)").Error; err != nil {
					log.Printf("Warning: Could not add FK constraint fk_businesses_owner: %v", err)
				}

				// Outlets -> Business
				if err := tx.Exec("ALTER TABLE outlets ADD CONSTRAINT fk_outlets_business FOREIGN KEY (business_id) REFERENCES businesses(id)").Error; err != nil {
					log.Printf("Warning: Could not add FK constraint fk_outlets_business: %v", err)
				}

				// UserOutletRoles -> Users, Outlets, Roles
				if err := tx.Exec("ALTER TABLE user_outlet_roles ADD CONSTRAINT fk_user_outlet_roles_user FOREIGN KEY (user_id) REFERENCES users(id)").Error; err != nil {
					log.Printf("Warning: Could not add FK constraint fk_user_outlet_roles_user: %v", err)
				}

				if err := tx.Exec("ALTER TABLE user_outlet_roles ADD CONSTRAINT fk_user_outlet_roles_outlet FOREIGN KEY (outlet_id) REFERENCES outlets(id)").Error; err != nil {
					log.Printf("Warning: Could not add FK constraint fk_user_outlet_roles_outlet: %v", err)
				}

				if err := tx.Exec("ALTER TABLE user_outlet_roles ADD CONSTRAINT fk_user_outlet_roles_role FOREIGN KEY (role_id) REFERENCES roles(id)").Error; err != nil {
					log.Printf("Warning: Could not add FK constraint fk_user_outlet_roles_role: %v", err)
				}

				// Customers -> Business
				if err := tx.Exec("ALTER TABLE customers ADD CONSTRAINT fk_customers_business FOREIGN KEY (business_id) REFERENCES businesses(id)").Error; err != nil {
					log.Printf("Warning: Could not add FK constraint fk_customers_business: %v", err)
				}

				// Categories -> Outlets
				if err := tx.Exec("ALTER TABLE categories ADD CONSTRAINT fk_categories_outlet FOREIGN KEY (outlet_id) REFERENCES outlets(id)").Error; err != nil {
					log.Printf("Warning: Could not add FK constraint fk_categories_outlet: %v", err)
				}

				// Products -> Outlets, Categories
				if err := tx.Exec("ALTER TABLE products ADD CONSTRAINT fk_products_outlet FOREIGN KEY (outlet_id) REFERENCES outlets(id)").Error; err != nil {
					log.Printf("Warning: Could not add FK constraint fk_products_outlet: %v", err)
				}

				if err := tx.Exec("ALTER TABLE products ADD CONSTRAINT fk_products_category FOREIGN KEY (category_id) REFERENCES categories(id)").Error; err != nil {
					log.Printf("Warning: Could not add FK constraint fk_products_category: %v", err)
				}

				// ProductVariants -> Products
				if err := tx.Exec("ALTER TABLE product_variants ADD CONSTRAINT fk_product_variants_product FOREIGN KEY (product_id) REFERENCES products(id)").Error; err != nil {
					log.Printf("Warning: Could not add FK constraint fk_product_variants_product: %v", err)
				}

				// Suppliers -> Outlets
				if err := tx.Exec("ALTER TABLE suppliers ADD CONSTRAINT fk_suppliers_outlet FOREIGN KEY (outlet_id) REFERENCES outlets(id)").Error; err != nil {
					log.Printf("Warning: Could not add FK constraint fk_suppliers_outlet: %v", err)
				}

				// InventoryMovements -> ProductVariants, Outlets, Users
				if err := tx.Exec("ALTER TABLE inventory_movements ADD CONSTRAINT fk_inventory_movements_product_variant FOREIGN KEY (product_variant_id) REFERENCES product_variants(id)").Error; err != nil {
					log.Printf("Warning: Could not add FK constraint fk_inventory_movements_product_variant: %v", err)
				}

				if err := tx.Exec("ALTER TABLE inventory_movements ADD CONSTRAINT fk_inventory_movements_outlet FOREIGN KEY (outlet_id) REFERENCES outlets(id)").Error; err != nil {
					log.Printf("Warning: Could not add FK constraint fk_inventory_movements_outlet: %v", err)
				}

				if err := tx.Exec("ALTER TABLE inventory_movements ADD CONSTRAINT fk_inventory_movements_user FOREIGN KEY (user_id) REFERENCES users(id)").Error; err != nil {
					log.Printf("Warning: Could not add FK constraint fk_inventory_movements_user: %v", err)
				}

				// Transactions -> Outlets, Users, Customers
				if err := tx.Exec("ALTER TABLE transactions ADD CONSTRAINT fk_transactions_outlet FOREIGN KEY (outlet_id) REFERENCES outlets(id)").Error; err != nil {
					log.Printf("Warning: Could not add FK constraint fk_transactions_outlet: %v", err)
				}

				if err := tx.Exec("ALTER TABLE transactions ADD CONSTRAINT fk_transactions_user FOREIGN KEY (user_id) REFERENCES users(id)").Error; err != nil {
					log.Printf("Warning: Could not add FK constraint fk_transactions_user: %v", err)
				}

				if err := tx.Exec("ALTER TABLE transactions ADD CONSTRAINT fk_transactions_customer FOREIGN KEY (customer_id) REFERENCES customers(id)").Error; err != nil {
					log.Printf("Warning: Could not add FK constraint fk_transactions_customer: %v", err)
				}

				// TransactionItems -> Transactions, ProductVariants
				if err := tx.Exec("ALTER TABLE transaction_items ADD CONSTRAINT fk_transaction_items_transaction FOREIGN KEY (transaction_id) REFERENCES transactions(id)").Error; err != nil {
					log.Printf("Warning: Could not add FK constraint fk_transaction_items_transaction: %v", err)
				}

				if err := tx.Exec("ALTER TABLE transaction_items ADD CONSTRAINT fk_transaction_items_product_variant FOREIGN KEY (product_variant_id) REFERENCES product_variants(id)").Error; err != nil {
					log.Printf("Warning: Could not add FK constraint fk_transaction_items_product_variant: %v", err)
				}

				// Payments -> Transactions
				if err := tx.Exec("ALTER TABLE payments ADD CONSTRAINT fk_payments_transaction FOREIGN KEY (transaction_id) REFERENCES transactions(id)").Error; err != nil {
					log.Printf("Warning: Could not add FK constraint fk_payments_transaction: %v", err)
				}

				// LoyaltyTransactions -> Customers
				if err := tx.Exec("ALTER TABLE loyalty_transactions ADD CONSTRAINT fk_loyalty_transactions_customer FOREIGN KEY (customer_id) REFERENCES customers(id)").Error; err != nil {
					log.Printf("Warning: Could not add FK constraint fk_loyalty_transactions_customer: %v", err)
				}

				return nil
			},
			Rollback: func(tx *gorm.DB) error {
				// Drop foreign key constraints
				constraints := []string{
					"fk_loyalty_transactions_customer",
					"fk_payments_transaction",
					"fk_transaction_items_product_variant",
					"fk_transaction_items_transaction",
					"fk_transactions_customer",
					"fk_transactions_user",
					"fk_transactions_outlet",
					"fk_inventory_movements_user",
					"fk_inventory_movements_outlet",
					"fk_inventory_movements_product_variant",
					"fk_suppliers_outlet",
					"fk_product_variants_product",
					"fk_products_category",
					"fk_products_outlet",
					"fk_categories_outlet",
					"fk_customers_business",
					"fk_user_outlet_roles_role",
					"fk_user_outlet_roles_outlet",
					"fk_user_outlet_roles_user",
					"fk_outlets_business",
					"fk_businesses_owner",
				}

				for _, constraint := range constraints {
					// Drop constraint if exists (PostgreSQL syntax)
					tx.Exec("ALTER TABLE IF EXISTS " + constraint + " DROP CONSTRAINT IF EXISTS " + constraint)
				}

				return nil
			},
		},
	})

	if err := m.Migrate(); err != nil {
		log.Fatalf("Could not migrate: %v", err)
	}

	log.Println("Migration completed successfully!")
}
