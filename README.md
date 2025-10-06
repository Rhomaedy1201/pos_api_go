# 🛒 Point of Sale (POS) API - Go

[![Go Version](https://img.shields.io/badge/Go-1.24.0-blue.svg)](https://golang.org/)
[![Gin Framework](https://img.shields.io/badge/Gin-1.11.0-green.svg)](https://github.com/gin-gonic/gin)
[![GORM](https://img.shields.io/badge/GORM-1.31.0-red.svg)](https://gorm.io/)
[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-Database-blue.svg)](https://www.postgresql.org/)

## 📋 Deskripsi

Point of Sale (POS) API adalah aplikasi backend untuk sistem kasir yang dibangun menggunakan **Go (Golang)** dengan framework **Gin** dan **GORM** sebagai ORM untuk database **PostgreSQL**. API ini menyediakan endpoint untuk mengelola produk, transaksi, dan data terkait sistem Point of Sale.

## 🚀 Teknologi yang Digunakan

### **Backend Framework & Libraries**
- **[Go](https://golang.org/)** `v1.24.0` - Programming Language
- **[Gin](https://github.com/gin-gonic/gin)** `v1.11.0` - HTTP Web Framework
- **[GORM](https://gorm.io/)** `v1.31.0` - ORM Library
- **[PostgreSQL Driver](https://github.com/go-gorm/postgres)** `v1.6.0` - Database Driver
- **[Gormigrate](https://github.com/go-gormigrate/gormigrate)** `v2.1.5` - Database Migration Tool
- **[GoDotEnv](https://github.com/joho/godotenv)** `v1.5.1` - Environment Variable Loader

### **Database**
- **[PostgreSQL](https://www.postgresql.org/)** - Primary Database

### **Development Tools**
- **[Air](https://github.com/cosmtrek/air)** - Live Reload untuk Development
- **Git** - Version Control System

## 📁 Struktur Proyek

```
pos_api_go/
├── cmd/
│   └── main.go              # Entry point aplikasi
├── config/
│   └── config.go            # Konfigurasi database dan environment
├── internal/
│   ├── database/
│   │   └── migrations.go    # Setup migrasi database
│   └── models/
│       └── models.go        # Model/Schema database
├── migrations/              # File migrasi database
├── pkg/                     # Package yang bisa digunakan eksternal
├── tmp/                     # Temporary files (build artifacts)
├── .air.toml               # Konfigurasi Air live reload
├── .env                    # Environment variables
├── go.mod                  # Go module dependencies
├── go.sum                  # Go module checksums
└── # POS API Go - Point of Sale System

A modern Point of Sale (POS) API built with Go, Gin, and GORM with PostgreSQL.

## Features

- 🏗️ **Clean Architecture** - Organized with handlers, middleware, services, and models
- 🔐 **Authentication & Authorization** - JWT-based auth with role-based access control
- 🗄️ **Database Management** - Structured migrations and seeders
- 📊 **Multi-tenant Support** - Business and outlet management
- 🛒 **Sales Management** - Complete transaction processing
- 📦 **Inventory Control** - Product and stock management
- 👥 **User Management** - Role-based user access
- 🔍 **RESTful API** - Clean and documented endpoints

## Quick Start

### Prerequisites

- Go 1.21 or higher
- PostgreSQL 12 or higher
- Make (optional, for easier commands)

### Installation

1. **Install dependencies**
   ```bash
   make deps
   ```

2. **Setup database**
   ```bash
   make migrate
   make seed
   ```

3. **Start the server**
   ```bash
   make dev  # Development mode with hot reload
   ```

## API Endpoints

### Authentication
- `POST /api/v1/auth/signin` - User sign in
- `POST /api/v1/auth/signup` - User registration
- `POST /api/v1/auth/refresh` - Refresh token

### Users (Protected)
- `GET /api/v1/users` - List users
- `POST /api/v1/users` - Create user
- `GET /api/v1/users/:id` - Get user by ID
- `PUT /api/v1/users/:id` - Update user
- `DELETE /api/v1/users/:id` - Delete user

### Products (Protected)
- `GET /api/v1/products` - List products
- `POST /api/v1/products` - Create product
- `GET /api/v1/products/:id` - Get product by ID
- `PUT /api/v1/products/:id` - Update product
- `DELETE /api/v1/products/:id` - Delete product

### Health Check
- `GET /health` - Database health check

## Development Commands

```bash
make deps      # Install dependencies
make dev       # Development with hot reload
make migrate   # Run migrations
make seed      # Run seeders
make build     # Build application
make run       # Run application
make test      # Run tests
make setup     # Setup everything
```               # Dokumentasi proyek
```

## 🛠️ Instalasi dan Setup

### **Prerequisites**
- Go 1.24.0 atau lebih baru
- PostgreSQL 12+ 
- Git

### **1. Clone Repository**
```bash
git clone https://github.com/Rhomaedy1201/pos_api_go.git
cd pos_api_go
```

### **2. Install Dependencies**
```bash
go mod download
```

### **3. Setup Database**
1. Buat database PostgreSQL:
```sql
CREATE DATABASE pos_go;
```

2. Copy file environment:
```bash
cp .env.example .env  # jika ada, atau buat manual
```

3. Konfigurasi file `.env`:
```env
APP_PORT=3000
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASS=your_password
DB_NAME=pos_go
```

### **4. Install Air (Live Reload) - Optional**
```bash
go install github.com/cosmtrek/air@latest
```

### **5. Jalankan Aplikasi**

#### **Dengan Air (Recommended untuk development):**
```bash
air
```

#### **Tanpa Air:**
```bash
go run cmd/main.go
```

#### **Build dan Run:**
```bash
go build -o tmp/main cmd/main.go
./tmp/main
```

## 🌐 API Endpoints

### **Health Check**
- **GET** `/health` - Cek status database dan aplikasi

### **General**
- **GET** `/` - Welcome message

### **Response Example:**

#### GET `/health`
```json
{
  "status": "ok",
  "message": "Database is healthy"
}
```

#### GET `/`
```json
{
  "message": "Hello World! Database connected successfully!"
}
```

## 📊 Models/Schema Database

### **User Model**
```go
type User struct {
    ID        uint           `json:"id" gorm:"primaryKey"`
    Name      string         `json:"name" gorm:"not null"`
    Email     string         `json:"email" gorm:"uniqueIndex;not null"`
    CreatedAt time.Time      `json:"created_at"`
    UpdatedAt time.Time      `json:"updated_at"`
    DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
```

### **Product Model**
```go
type Product struct {
    ID          uint           `json:"id" gorm:"primaryKey"`
    Name        string         `json:"name" gorm:"not null"`
    Description string         `json:"description"`
    Price       float64        `json:"price" gorm:"not null"`
    Stock       int            `json:"stock" gorm:"default:0"`
    CreatedAt   time.Time      `json:"created_at"`
    UpdatedAt   time.Time      `json:"updated_at"`
    DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
```

## ⚙️ Konfigurasi

### **Environment Variables**
| Variable | Description | Default |
|----------|-------------|---------|
| `APP_PORT` | Port aplikasi | `3000` |
| `DB_HOST` | Database host | `localhost` |
| `DB_PORT` | Database port | `5432` |
| `DB_USER` | Database username | `postgres` |
| `DB_PASS` | Database password | - |
| `DB_NAME` | Database name | `pos_go` |

### **Air Configuration**
File `.air.toml` sudah dikonfigurasi untuk:
- Auto reload saat ada perubahan file `.go`
- Build output ke `./tmp/main`
- Watch semua direktori kecuali `tmp`, `vendor`, `testdata`

## 🚀 Development

### **Menambah Model Baru**
1. Tambahkan struct di `internal/models/models.go`
2. Tambahkan auto migration di `cmd/main.go`:
```go
err := db.AutoMigrate(&models.User{}, &models.Product{}, &models.NewModel{})
```

### **Menambah Migration dengan Gormigrate**
1. Edit `internal/database/migrations.go`
2. Tambahkan migration baru:
```go
{
    ID: "202410020001",
    Migrate: func(tx *gorm.DB) error {
        // Schema changes
        return nil
    },
    Rollback: func(tx *gorm.DB) error {
        // Rollback changes
        return nil
    },
},
```

### **Hot Reload Development**
Gunakan Air untuk development yang lebih produktif:
```bash
air
```
Server akan otomatis restart saat ada perubahan kode.

## 📝 Testing

### **Manual Testing**
```bash
# Test health endpoint
curl http://localhost:3000/health

# Test main endpoint
curl http://localhost:3000/
```

### **Build Test**
```bash
go build -o tmp/main cmd/main.go
```

## 🔧 Troubleshooting

### **Database Connection Issues**
1. Pastikan PostgreSQL sudah running
2. Cek kredensial di file `.env`
3. Pastikan database `pos_go` sudah dibuat

### **Port Already in Use**
Ubah `APP_PORT` di file `.env` ke port lain (misal: 8080)

### **Module Issues**
```bash
go mod tidy
go mod download
```

## 📈 Roadmap

- [ ] Authentication & Authorization (JWT)
- [ ] Product Management CRUD
- [ ] User Management
- [ ] Transaction Management
- [ ] Inventory Management
- [ ] Reporting System
- [ ] Unit Testing
- [ ] API Documentation (Swagger)
- [ ] Docker Support
- [ ] CI/CD Pipeline

## 🤝 Contributing

1. Fork repository
2. Buat feature branch: `git checkout -b feature/amazing-feature`
3. Commit changes: `git commit -m 'Add amazing feature'`
4. Push ke branch: `git push origin feature/amazing-feature`
5. Buat Pull Request

## 📄 License

Distributed under the MIT License. See `LICENSE` file for more information.

## 👨‍💻 Author

**Muhammad Rhomaedi**
- GitHub: [@Rhomaedy1201](https://github.com/Rhomaedy1201)

## 🙏 Acknowledgments

- [Gin Web Framework](https://github.com/gin-gonic/gin)
- [GORM](https://gorm.io/)
- [Air Live Reload](https://github.com/cosmtrek/air)
- [PostgreSQL](https://www.postgresql.org/)

---

⭐ **Jika proyek ini membantu, berikan star di GitHub repository!**