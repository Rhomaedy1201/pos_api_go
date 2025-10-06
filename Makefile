# Makefile for POS API Go

.PHONY: help build run clean test migrate seed dev docker

# Default target
help:
	@echo "Available commands:"
	@echo "  build     - Build the application"
	@echo "  run       - Run the application"
	@echo "  dev       - Run in development mode with hot reload"
	@echo "  clean     - Clean build artifacts"
	@echo "  test      - Run tests"
	@echo "  migrate   - Run database migrations"
	@echo "  seed      - Run database seeders"
	@echo "  fresh     - Fresh migration (drop all tables and re-run)"
	@echo "  rollback  - Rollback last migration"
	@echo "  status    - Show migration status"
	@echo "  docker    - Build and run with Docker"
	@echo "  deps      - Install dependencies"

# Install dependencies
deps:
	go mod download
	go mod tidy

# Build the application
build:
	@echo "Building application..."
	go build -o bin/pos-api cmd/main.go

# Run the application
run: build
	@echo "Starting server..."
	./bin/pos-api

# Development mode (requires air for hot reload)
dev:
	@echo "Starting development server with hot reload..."
	@if command -v air > /dev/null; then \
		air; \
	else \
		echo "Installing air for hot reload..."; \
		go install github.com/cosmtrek/air@latest; \
		air; \
	fi

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	rm -rf bin/
	rm -rf tmp/

# Run tests
test:
	@echo "Running tests..."
	go test -v ./...

# Database migrations
migrate:
	@echo "Running migrations..."
	go run cmd/migrate/main.go -up

# Run seeders
seed:
	@echo "Running seeders..."
	go run cmd/migrate/main.go -seed

# Fresh migration
fresh:
	@echo "Running fresh migration..."
	go run cmd/migrate/main.go -fresh

# Rollback migration
rollback:
	@echo "Rolling back migration..."
	@read -p "Enter migration ID to rollback to: " id; \
	go run cmd/migrate/main.go -down $$id

# Migration status
status:
	@echo "Checking migration status..."
	go run cmd/migrate/main.go -status

# Docker build and run
docker:
	@echo "Building and running with Docker..."
	docker build -t pos-api .
	docker run -p 3000:3000 pos-api

# Format code
fmt:
	@echo "Formatting code..."
	go fmt ./...

# Lint code
lint:
	@echo "Linting code..."
	@if command -v golangci-lint > /dev/null; then \
		golangci-lint run; \
	else \
		echo "Installing golangci-lint..."; \
		go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest; \
		golangci-lint run; \
	fi

# Generate swagger docs (if using swag)
docs:
	@echo "Generating API documentation..."
	@if command -v swag > /dev/null; then \
		swag init -g cmd/main.go; \
	else \
		echo "Installing swag..."; \
		go install github.com/swaggo/swag/cmd/swag@latest; \
		swag init -g cmd/main.go; \
	fi

# Setup development environment
setup: deps migrate seed
	@echo "Development environment setup complete!"