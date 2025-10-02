package main

import (
	"log"
	"pos_api_go/config"
	"pos_api_go/internal/models"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load environment variables
	config.LoadEnv()

	// Connect to database
	config.ConnectDatabase()

	// Auto migrate models (contoh sederhana)
	db := config.GetDB()
	err := db.AutoMigrate(&models.User{}, &models.Product{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
	log.Println("Database migration completed!")

	// Inisialisasi Gin
	router := gin.Default()

	// Membuat route dengan method GET
	router.GET("/", func(c *gin.Context) {
		// Return response JSON
		c.JSON(200, gin.H{
			"message": "Hello World! Database connected successfully!",
		})
	})

	// Health check route untuk database
	router.GET("/health", func(c *gin.Context) {
		db := config.GetDB()
		sqlDB, err := db.DB()
		if err != nil {
			c.JSON(500, gin.H{
				"status":  "error",
				"message": "Database connection error",
			})
			return
		}

		if err := sqlDB.Ping(); err != nil {
			c.JSON(500, gin.H{
				"status":  "error",
				"message": "Database ping failed",
			})
			return
		}

		c.JSON(200, gin.H{
			"status":  "ok",
			"message": "Database is healthy",
		})
	})

	// Mulai server
	router.Run(":" + config.GetEnv("APP_PORT", "3000"))
}
