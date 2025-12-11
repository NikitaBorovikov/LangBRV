include .env
.PHONY: up first-up down build rebuild migrate-up migrate-down 

DB_DSN = postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable
MIGRATION_PATH = /app/migrations

up:
	docker-compose up
first-up:
	docker-compose --profile migrate up
down:
	docker-compose down 
rebuild:
	docker-compose build --no-cache

migrate-up:
	docker-compose exec app migrate -path $(MIGRATION_PATH) -database "$(DB_DSN)" up
migrate-down:
	docker-compose exec app migrate -path $(MIGRATION_PATH) -database "$(DB_DSN)" down 1