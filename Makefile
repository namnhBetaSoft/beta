# Simple Makefile for a Go project

# Build the application
all: build test

build:
	@echo "Building..."
	@cd backend && go build -o main cmd/api/main.go

# Run the application
run:
	@cd backend && go run cmd/api/main.go &
	@npm install --prefer-offline --no-fund --prefix ./frontend
	@npm run dev --prefix ./frontend
# Create DB container
docker-run:
	@if docker compose up --build 2>/dev/null; then \
		: ; \
	else \
		echo "Falling back to Docker Compose V1"; \
		docker-compose up --build; \
	fi

# Shutdown DB container
docker-down:
	@if docker compose down 2>/dev/null; then \
		: ; \
	else \
		echo "Falling back to Docker Compose V1"; \
		docker-compose down; \
	fi

# Test the application
test:
	@echo "Testing..."
	@cd backend && go test ./... -v
# Integrations Tests for the application
itest:
	@echo "Running integration tests..."
	@cd backend && go test ./internal/database -v

# Clean the binary
clean:
	@echo "Cleaning..."
	@rm -f backend/main

# Live Reload
watch:
	@if command -v air > /dev/null; then \
            cd backend && air; \
            echo "Watching...";\
        else \
            read -p "Go's 'air' is not installed on your machine. Do you want to install it? [Y/n] " choice; \
            if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
                go install github.com/air-verse/air@latest; \
                cd backend && air; \
                echo "Watching...";\
            else \
                echo "You chose not to install air. Exiting..."; \
                exit 1; \
            fi; \
        fi

.PHONY: all build run test clean watch docker-run docker-down itest
