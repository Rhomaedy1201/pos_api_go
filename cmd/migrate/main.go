package main

import (
	"flag"
	"log"
	"pos_api_go/config"
	"pos_api_go/internal/database"
	"pos_api_go/internal/database/seeders"
)

func main() {
	var (
		up     = flag.Bool("up", false, "Run migrations")
		down   = flag.String("down", "", "Rollback to specific migration ID")
		fresh  = flag.Bool("fresh", false, "Drop all tables and re-run migrations")
		seed   = flag.Bool("seed", false, "Run seeders")
		status = flag.Bool("status", false, "Show migration status")
	)
	flag.Parse()

	// Load configuration
	config.LoadEnv()
	config.ConnectDatabase()
	db := config.GetDB()

	switch {
	case *up:
		log.Println("Running migrations...")
		if err := database.RunMigrations(db); err != nil {
			log.Fatal("Migration failed:", err)
		}
		log.Println("✅ Migrations completed successfully!")

	case *down != "":
		log.Printf("Rolling back to migration: %s", *down)
		if err := database.RollbackMigration(db, *down); err != nil {
			log.Fatal("Rollback failed:", err)
		}
		log.Println("✅ Rollback completed successfully!")

	case *fresh:
		log.Println("Running fresh migration (WARNING: This will drop all tables)")
		// TODO: Implement fresh migration
		// This should drop all tables and re-run all migrations
		log.Println("Fresh migration not implemented yet")

	case *seed:
		log.Println("Running seeders...")
		if err := seeders.SeedAllDefaultData(db); err != nil {
			log.Fatal("Seeding failed:", err)
		}
		log.Println("✅ Seeding completed successfully!")

	case *status:
		log.Println("Migration status:")
		// TODO: Implement migration status check
		log.Println("Status check not implemented yet")

	default:
		log.Println("Migration Management Tool")
		log.Println("Usage:")
		log.Println("  -up      Run all pending migrations")
		log.Println("  -down    Rollback to specific migration ID")
		log.Println("  -fresh   Drop all tables and re-run migrations")
		log.Println("  -seed    Run data seeders")
		log.Println("  -status  Show migration status")
		log.Println()
		log.Println("Examples:")
		log.Println("  go run cmd/migrate/main.go -up")
		log.Println("  go run cmd/migrate/main.go -down 202410030001_initial_schema")
		log.Println("  go run cmd/migrate/main.go -seed")
	}
}
