# Variables
DB_URL=postgres://postgres:postgres@localhost:5432/auth-service-db?sslmode=disable

MIGRATIONS_DIR=database/migrations

# Default target
.DEFAULT_GOAL := help

# Help: Display available commands
help:
	@echo "Makefile commands for managing Goose migrations:"
	@echo ""
	@echo "  make up        Apply all pending migrations"
	@echo "  make down      Roll back the last migration"
	@echo "  make redo      Roll back the last migration and reapply it"
	@echo "  make status    Show the status of all migrations"
	@echo "  make create    Create a new migration file"
	@echo ""

# Apply all pending migrations
up:
	goose -dir $(MIGRATIONS_DIR) postgres "$(DB_URL)" up

# Roll back the last migration
down:
	goose -dir $(MIGRATIONS_DIR) postgres "$(DB_URL)" down

# Roll back the last migration and reapply it
redo:
	goose -dir $(MIGRATIONS_DIR) postgres "$(DB_URL)" redo


#Apply Migrations: make up
#Roll Back the Last Migration: make down
#Reapply the Last Migration: make redo


# RUN command goose -dir databse/migrations postgres "postgres://postgres:postgres@localhost:5432/auth-service-db?sslmode=disable" up --> to up goose migration / MAKE UP