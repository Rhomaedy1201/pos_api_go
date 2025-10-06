package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func addIndexesMigration() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "202410030002_add_indexes",
		Migrate: func(tx *gorm.DB) error {
			indexes := []string{
				// Users indexes
				"CREATE INDEX IF NOT EXISTS idx_users_email ON users(email)",
				"CREATE INDEX IF NOT EXISTS idx_users_created_at ON users(created_at)",
				"CREATE INDEX IF NOT EXISTS idx_users_status ON users(status)",

				// Products indexes
				"CREATE INDEX IF NOT EXISTS idx_products_business_id ON products(business_id)",
				"CREATE INDEX IF NOT EXISTS idx_products_category_id ON products(category_id)",
				"CREATE INDEX IF NOT EXISTS idx_products_sku ON products(sku)",
				"CREATE INDEX IF NOT EXISTS idx_products_barcode ON products(barcode)",
				"CREATE INDEX IF NOT EXISTS idx_products_status ON products(status)",

				// Transactions indexes
				"CREATE INDEX IF NOT EXISTS idx_transactions_outlet_id ON transactions(outlet_id)",
				"CREATE INDEX IF NOT EXISTS idx_transactions_created_at ON transactions(created_at)",
				"CREATE INDEX IF NOT EXISTS idx_transactions_status ON transactions(status)",
				"CREATE INDEX IF NOT EXISTS idx_transactions_customer_id ON transactions(customer_id)",

				// Transaction items indexes
				"CREATE INDEX IF NOT EXISTS idx_transaction_items_transaction_id ON transaction_items(transaction_id)",
				"CREATE INDEX IF NOT EXISTS idx_transaction_items_product_variant_id ON transaction_items(product_variant_id)",

				// Inventory movements indexes
				"CREATE INDEX IF NOT EXISTS idx_inventory_movements_product_variant_id ON inventory_movements(product_variant_id)",
				"CREATE INDEX IF NOT EXISTS idx_inventory_movements_outlet_id ON inventory_movements(outlet_id)",
				"CREATE INDEX IF NOT EXISTS idx_inventory_movements_created_at ON inventory_movements(created_at)",
				"CREATE INDEX IF NOT EXISTS idx_inventory_movements_type ON inventory_movements(type)",

				// Payments indexes
				"CREATE INDEX IF NOT EXISTS idx_payments_transaction_id ON payments(transaction_id)",
				"CREATE INDEX IF NOT EXISTS idx_payments_method ON payments(payment_method)",
				"CREATE INDEX IF NOT EXISTS idx_payments_created_at ON payments(created_at)",

				// Customers indexes
				"CREATE INDEX IF NOT EXISTS idx_customers_business_id ON customers(business_id)",
				"CREATE INDEX IF NOT EXISTS idx_customers_email ON customers(email)",
				"CREATE INDEX IF NOT EXISTS idx_customers_phone_number ON customers(phone_number)",

				// Categories indexes
				"CREATE INDEX IF NOT EXISTS idx_categories_business_id ON categories(business_id)",
				"CREATE INDEX IF NOT EXISTS idx_categories_status ON categories(status)",

				// Suppliers indexes
				"CREATE INDEX IF NOT EXISTS idx_suppliers_business_id ON suppliers(business_id)",
				"CREATE INDEX IF NOT EXISTS idx_suppliers_status ON suppliers(status)",

				// Outlets indexes
				"CREATE INDEX IF NOT EXISTS idx_outlets_business_id ON outlets(business_id)",
				"CREATE INDEX IF NOT EXISTS idx_outlets_status ON outlets(status)",

				// User outlet roles indexes
				"CREATE INDEX IF NOT EXISTS idx_user_outlet_roles_user_id ON user_outlet_roles(user_id)",
				"CREATE INDEX IF NOT EXISTS idx_user_outlet_roles_outlet_id ON user_outlet_roles(outlet_id)",
				"CREATE INDEX IF NOT EXISTS idx_user_outlet_roles_role_id ON user_outlet_roles(role_id)",
			}

			for _, index := range indexes {
				if err := tx.Exec(index).Error; err != nil {
					return err
				}
			}
			return nil
		},
		Rollback: func(tx *gorm.DB) error {
			// Drop indexes
			dropIndexes := []string{
				"DROP INDEX IF EXISTS idx_users_email",
				"DROP INDEX IF EXISTS idx_users_created_at",
				"DROP INDEX IF EXISTS idx_users_status",
				"DROP INDEX IF EXISTS idx_products_business_id",
				"DROP INDEX IF EXISTS idx_products_category_id",
				"DROP INDEX IF EXISTS idx_products_sku",
				"DROP INDEX IF EXISTS idx_products_barcode",
				"DROP INDEX IF EXISTS idx_products_status",
				"DROP INDEX IF EXISTS idx_transactions_outlet_id",
				"DROP INDEX IF EXISTS idx_transactions_created_at",
				"DROP INDEX IF EXISTS idx_transactions_status",
				"DROP INDEX IF EXISTS idx_transactions_customer_id",
				"DROP INDEX IF EXISTS idx_transaction_items_transaction_id",
				"DROP INDEX IF EXISTS idx_transaction_items_product_variant_id",
				"DROP INDEX IF EXISTS idx_inventory_movements_product_variant_id",
				"DROP INDEX IF EXISTS idx_inventory_movements_outlet_id",
				"DROP INDEX IF EXISTS idx_inventory_movements_created_at",
				"DROP INDEX IF EXISTS idx_inventory_movements_type",
				"DROP INDEX IF EXISTS idx_payments_transaction_id",
				"DROP INDEX IF EXISTS idx_payments_method",
				"DROP INDEX IF EXISTS idx_payments_created_at",
				"DROP INDEX IF EXISTS idx_customers_business_id",
				"DROP INDEX IF EXISTS idx_customers_email",
				"DROP INDEX IF EXISTS idx_customers_phone_number",
				"DROP INDEX IF EXISTS idx_categories_business_id",
				"DROP INDEX IF EXISTS idx_categories_status",
				"DROP INDEX IF EXISTS idx_suppliers_business_id",
				"DROP INDEX IF EXISTS idx_suppliers_status",
				"DROP INDEX IF EXISTS idx_outlets_business_id",
				"DROP INDEX IF EXISTS idx_outlets_status",
				"DROP INDEX IF EXISTS idx_user_outlet_roles_user_id",
				"DROP INDEX IF EXISTS idx_user_outlet_roles_outlet_id",
				"DROP INDEX IF EXISTS idx_user_outlet_roles_role_id",
			}

			for _, drop := range dropIndexes {
				tx.Exec(drop)
			}
			return nil
		},
	}
}
