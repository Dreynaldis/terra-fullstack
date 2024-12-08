include .env
export $(shell sed 's/=.*//' .env)
.PHONY: build run clean


build:
	@echo "Building the application..."
	@go build -o main cmd/api/main.go

migrate-up :
	migrate -path internal/database/migrations -database "postgres://$(DB_USERNAME):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" up

migrate-down :
	migrate -path internal/database/migrations -database "postgres://$(DB_USERNAME):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" down

run:
	@echo "Running the application..."
	@go run cmd/api/main.go

clean:
	@echo "Cleaning up..."
	@rm -f main