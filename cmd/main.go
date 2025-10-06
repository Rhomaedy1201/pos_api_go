package main

import (
	"log"
	"pos_api_go/config"
	"pos_api_go/internal/database"
	"pos_api_go/internal/handlers"
	"pos_api_go/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load environment variables
	config.LoadEnv()

	// Connect to database
	config.ConnectDatabase()

	// Run migrations
	db := config.GetDB()
	if err := database.RunMigrations(db); err != nil {
		log.Fatal("Failed to run migrations:", err)
	}

	// Initialize handlers
	handler := handlers.NewHandler()

	// Setup routes
	router := routes.NewRouter(handler)
	engine := router.SetupAPIRoutes()

	// Add basic root endpoint
	engine.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "POS API Go - Server is running",
			"version": "1.0.0",
			"status":  "ok",
		})
	})

	// Start server
	port := config.GetEnv("APP_PORT", "3000")
	log.Printf("🚀 Server starting on port %s", port)
	log.Printf("📋 Health check available at: http://localhost:%s/health", port)
	log.Printf("🔗 API endpoints available at: http://localhost:%s/api/v1", port)

	if err := engine.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
