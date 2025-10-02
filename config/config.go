package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// LoadEnv membaca file .env
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: No .env file found, using system environment variables")
	}
}

// GetEnv mengambil nilai dari environment variable
func GetEnv(key string, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}

// ConnectDatabase menginisialisasi koneksi database
func ConnectDatabase() {
	// Membuat DSN (Data Source Name) untuk PostgreSQL
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		GetEnv("DB_HOST", "localhost"),
		GetEnv("DB_USER", "postgres"),
		GetEnv("DB_PASS", ""),
		GetEnv("DB_NAME", "pos_go"),
		GetEnv("DB_PORT", "5432"),
	)

	// Membuka koneksi ke database
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	DB = database
	log.Println("Database connected successfully!")
}

// GetDB mengembalikan instance database
func GetDB() *gorm.DB {
	return DB
}
