.PHONY: help build test run clean docker-build docker-up docker-down docker-logs

# Default target
.DEFAULT_GOAL := help

# Application name
APP_NAME := app

# Build directory
BUILD_DIR := bin

help: ## Display this help message
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@awk '/^[a-zA-Z\-_0-9]+:/ { \
		helpMessage = match(lastLine, /^## (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 3, RLENGTH); \
			printf "  %-20s %s\n", helpCommand, helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

build: ## Build the application
	@echo "Building..."
	@mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(APP_NAME) cmd/app/main.go

test: deps ## Run tests with updated dependencies
	@echo "Running tests..."
	go test -v ./...

run: build ## Build and run the application
	@echo "Running..."
	./$(BUILD_DIR)/$(APP_NAME)

clean: ## Clean build artifacts
	@echo "Cleaning..."
	rm -rf $(BUILD_DIR)
	go clean

deps: ## Install dependencies
	@echo "Installing dependencies..."
	go mod tidy
	go mod verify

lint: ## Run linters
	@echo "Running linters..."
	golangci-lint run

# Docker commands
docker-build: ## Build Docker image
	@echo "Building Docker image..."
	docker build -t $(APP_NAME) .

docker-up: ## Start Docker containers
	@echo "Starting Docker containers..."
	docker-compose up -d

docker-down: ## Stop Docker containers
	@echo "Stopping Docker containers..."
	docker-compose down

docker-logs: ## View Docker logs
	@echo "Viewing Docker logs..."
	docker-compose logs -f

.DEFAULT_GOAL := help 