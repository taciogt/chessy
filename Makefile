.PHONY: help test test-v cover build run fmt vet tidy clean

help: ## Show this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-12s\033[0m %s\n", $$1, $$2}'

test: ## Run all tests
	go test ./...

test-v: ## Run all tests with verbose output
	go test -v ./...

cover: ## Run tests with coverage report
	go test -cover ./...

build: ## Build the application
	go build ./...

run: ## Run the application
	go run ./...

fmt: ## Format the code
	go fmt ./...

vet: ## Run go vet
	go vet ./...

tidy: ## Tidy go modules
	go mod tidy

clean: ## Remove build artifacts
	go clean ./...
