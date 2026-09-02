DB_URL=postgres://postgres:postgres@localhost:5432/gochess?sslmode=disable
MIGRATIONS_DIR=internal/database/migrations

.PHONY: migrate-up migrate-down migrate-create generate

# Crear una nueva migración: make migrate-create NAME=add_users
migrate-create:
	goose -dir $(MIGRATIONS_DIR) -s create $(NAME) sql

# Aplicar migraciones
migrate-up:
	goose -dir $(MIGRATIONS_DIR) postgres "$(DB_URL)" up

# Revertir última migración
migrate-down:
	goose -dir $(MIGRATIONS_DIR) postgres "$(DB_URL)" down

# Generar código de sqlc
generate:
	sqlc generate

# Pipeline de desarrollo completo
dev-db: migrate-up generate