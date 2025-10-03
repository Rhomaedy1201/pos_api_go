# 📊 Database Schema & Relasi Tabel

## 🏗️ Struktur Relasi Database POS

### **1. Core Tables (Bisnis & Outlet)**

```
Business (businesses)
├── id (UUID, PK)
├── name (varchar)
├── owner_id (UUID, FK → users.id)
├── status (enum: active, suspended, closed, trial)
├── created_at, updated_at
└── Relasi:
    └── outlets[] (One-to-Many)

Outlets (outlets)
├── id (UUID, PK)
├── business_id (UUID, FK → businesses.id)
├── name (varchar)
├── address (text)
├── status (enum: active, inactive)
├── created_at, updated_at
└── Relasi:
    ├── business (Many-to-One)
    ├── categories[] (One-to-Many)
    ├── products[] (One-to-Many)
    ├── suppliers[] (One-to-Many)
    └── transactions[] (One-to-Many)
```

### **2. Authentication Tables**

```
Users (users)
├── id (UUID, PK)
├── name (varchar)
├── email (varchar, unique)
├── password (varchar)
├── phone_number (varchar, unique)
├── created_at, updated_at
└── Relasi:
    ├── user_outlet_roles[] (One-to-Many)
    ├── transactions[] (One-to-Many)
    └── inventory_movements[] (One-to-Many)

Roles (roles)
├── id (UUID, PK)
├── name (varchar, unique)
├── created_at
└── Relasi:
    └── user_outlet_roles[] (One-to-Many)

UserOutletRoles (user_outlet_roles)
├── id (UUID, PK)
├── user_id (UUID, FK → users.id)
├── outlet_id (UUID, FK → outlets.id)
├── role_id (UUID, FK → roles.id)
├── created_at, updated_at
└── Relasi:
    ├── user (Many-to-One)
    ├── outlet (Many-to-One)
    └── role (Many-to-One)
```

### **3. Customer Tables**

```
Customers (customers)
├── id (UUID, PK)
├── business_id (UUID, FK → businesses.id)
├── name (varchar)
├── phone_number (varchar, unique)
├── email (varchar, unique)
├── address (text)
├── loyalty_points (int, default: 0)
├── created_at, updated_at, deleted_at
└── Relasi:
    ├── business (Many-to-One)
    ├── transactions[] (One-to-Many)
    └── loyalty_transactions[] (One-to-Many)

LoyaltyTransactions (loyalty_transactions)
├── id (UUID, PK)
├── customer_id (UUID, FK → customers.id)
├── transaction_id (UUID, FK → transactions.id)
├── point_change (decimal)
├── description (text)
├── created_at
└── Relasi:
    └── customer (Many-to-One)
```

### **4. Inventory Tables**

```
Categories (categories)
├── id (UUID, PK)
├── outlet_id (UUID, FK → outlets.id)
├── name (varchar)
├── status (enum: active, inactive)
├── created_at, updated_at
└── Relasi:
    ├── outlet (Many-to-One)
    └── products[] (One-to-Many)

Products (products)
├── id (UUID, PK)
├── outlet_id (UUID, FK → outlets.id)
├── category_id (UUID, FK → categories.id)
├── name (varchar)
├── description (text)
├── image_url (text)
├── is_active (boolean, default: true)
├── created_at, updated_at, deleted_at
└── Relasi:
    ├── outlet (Many-to-One)
    ├── category (Many-to-One)
    └── product_variants[] (One-to-Many)

ProductVariants (product_variants)
├── id (UUID, PK)
├── product_id (UUID, FK → products.id)
├── sku (varchar, unique)
├── name (varchar)
├── cost_price (decimal)
├── sell_price (decimal)
├── stock_qty (int, default: 0)
├── created_at, updated_at
└── Relasi:
    ├── product (Many-to-One)
    ├── transaction_items[] (One-to-Many)
    └── inventory_movements[] (One-to-Many)

Suppliers (suppliers)
├── id (UUID, PK)
├── outlet_id (UUID, FK → outlets.id)
├── name (varchar)
├── phone (varchar)
├── email (varchar)
├── address (text)
├── status (enum: active, inactive)
├── created_at, updated_at
└── Relasi:
    └── outlet (Many-to-One)

InventoryMovements (inventory_movements)
├── id (UUID, PK)
├── product_variant_id (UUID, FK → product_variants.id)
├── outlet_id (UUID, FK → outlets.id)
├── user_id (UUID, FK → users.id)
├── type (enum: sale, purchase, return, adjustment, transfer_in, transfer_out)
├── qty_change (decimal)
├── notes (text)
├── reference_id (UUID, nullable)
├── created_at
└── Relasi:
    ├── product_variant (Many-to-One)
    ├── outlet (Many-to-One)
    └── user (Many-to-One)
```

### **5. Sales Tables**

```
Transactions (transactions)
├── id (UUID, PK)
├── outlet_id (UUID, FK → outlets.id)
├── user_id (UUID, FK → users.id)
├── customer_id (UUID, FK → customers.id)
├── invoice_number (varchar, unique)
├── sub_total (decimal)
├── total_discount (decimal)
├── tax_amount (decimal)
├── grand_total (decimal)
├── status (varchar)
├── transaction_time (timestamp)
├── created_at, updated_at
└── Relasi:
    ├── outlet (Many-to-One)
    ├── user (Many-to-One)
    ├── customer (Many-to-One)
    ├── transaction_items[] (One-to-Many)
    └── payments[] (One-to-Many)

TransactionItems (transaction_items)
├── id (UUID, PK)
├── transaction_id (UUID, FK → transactions.id)
├── product_variant_id (UUID, FK → product_variants.id)
├── quantity (int)
├── price_per_unit (decimal)
├── discount_per_unit (decimal)
├── total_price (decimal)
├── created_at, updated_at
└── Relasi:
    ├── transaction (Many-to-One)
    └── product_variant (Many-to-One)

Payments (payments)
├── id (UUID, PK)
├── transaction_id (UUID, FK → transactions.id)
├── payment_method (varchar)
├── amount_paid (decimal)
├── reference_number (varchar)
├── payment_time (timestamp)
├── created_at, updated_at
└── Relasi:
    └── transaction (Many-to-One)
```

## 🔗 Diagram Relasi Visual

```
┌─────────────┐    ┌─────────────┐    ┌─────────────┐
│   Business  │────│   Outlets   │────│ Categories  │
│             │    │             │    │             │
└─────────────┘    └─────────────┘    └─────────────┘
       │                   │                   │
       │                   │                   │
┌─────────────┐    ┌─────────────┐    ┌─────────────┐
│  Customers  │    │ Suppliers   │    │  Products   │
│             │    │             │    │             │
└─────────────┘    └─────────────┘    └─────────────┘
       │                                       │
       │                                       │
┌─────────────┐                      ┌─────────────┐
│    Users    │                      │   Product   │
│             │                      │  Variants   │
└─────────────┘                      └─────────────┘
       │                                       │
       │                                       │
┌─────────────┐    ┌─────────────┐    ┌─────────────┐
│    Roles    │    │Transactions │────│Transaction  │
│             │    │             │    │   Items     │
└─────────────┘    └─────────────┘    └─────────────┘
       │                   │
       │                   │
┌─────────────┐    ┌─────────────┐
│UserOutlet   │    │  Payments   │
│   Roles     │    │             │
└─────────────┘    └─────────────┘
```

## 🎯 Fitur Relasi GORM yang Digunakan

### **1. Foreign Keys**
- Semua relasi menggunakan UUID sebagai primary key dan foreign key
- Constraint names yang konsisten untuk referential integrity

### **2. Relasi Types**
- **One-to-Many**: Business → Outlets, Product → ProductVariants
- **Many-to-One**: Outlet → Business, Transaction → Customer
- **Many-to-Many**: Users ↔ Outlets (melalui UserOutletRoles)

### **3. Soft Delete**
- Customers, Products menggunakan `gorm.DeletedAt`
- Data tidak benar-benar dihapus, hanya ditandai sebagai deleted

### **4. JSON Tags**
- Semua field memiliki JSON tags untuk API response
- Password field menggunakan `json:"-"` untuk keamanan
- Relasi menggunakan `omitempty` untuk optional loading

### **5. Migration Strategy**
- **Step 1**: Create all tables dengan AUTO MIGRATE
- **Step 2**: Add foreign key constraints secara manual
- Rollback support untuk development

## 🚀 Cara Menggunakan

1. **Enable UUID Extension** (PostgreSQL):
   ```sql
   CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
   ```

2. **Run Migration**:
   ```go
   database.RunMigrations(db)
   ```

3. **Preload Relasi**:
   ```go
   // Load product dengan category dan variants
   var product inventory.Products
   db.Preload("Category").Preload("ProductVariants").First(&product, id)
   
   // Load transaction dengan items dan payments
   var transaction sales.Transactions
   db.Preload("TransactionItems.ProductVariant").
      Preload("Payments").
      First(&transaction, id)
   ```

## 📝 Catatan Penting

- Semua tabel menggunakan UUID untuk primary key
- Timestamps otomatis dihandle oleh GORM
- Enum values harus sesuai dengan yang didefinisikan di schema
- Foreign key constraints menambah integritas data
- Soft delete tersedia untuk tabel yang membutuhkan audit trail