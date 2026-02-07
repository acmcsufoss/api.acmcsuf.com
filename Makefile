.DEFAULT_GOAL := build

include .env

.PHONY: help
help: ## Display this help screen
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

BIN_DIR := bin
API_NAME := acmcsuf-api
CLI_NAME := acmcsuf-cli

MIGRATE_DIR := internal/db/sql/migrations
DB_URL := sqlite3://dev.db

GENERATE_DEPS := $(wildcard internal/api/handlers/*.go)
GENERATE_MARKER := .generate.marker

.PHONY:fmt run build vet check test check-sql fix-sql clean release generate migrate-up migrate-down

fmt: ## Format all go files
	@go fmt ./...

VERSION := $(shell git describe --tags --always --dirty 2> /dev/null || echo "dev")


$(GENERATE_MARKER): $(GENERATE_DEPS)
	go generate ./...
	@touch $@

generate: fmt $(GENERATE_MARKER) ## Generate all necessary files

run: build ## Build and run the api
	./$(BIN_DIR)/$(API_NAME)

build: generate ## Build the api and cli binaries
	@mkdir -p $(BIN_DIR)
	go build -ldflags "-X main.Version=$(VERSION)" -o $(BIN_DIR)/$(API_NAME) ./cmd/$(API_NAME)
	go build -ldflags "-X cli.Version=$(VERSION)" -o $(BIN_DIR)/$(CLI_NAME) ./cmd/$(CLI_NAME)

vet: ## Vet all go files
	go vet ./...

check: ## Run static analysis on all go files
	staticcheck ./...

test: check ## Run all tests
	go test ./...

check-sql: ## Lint all sql files
	sqlfluff lint --dialect sqlite

fix-sql: ## Fix all sql files
	sqlfluff fix --dialect sqlite

release: ## Create a new release tag
	@echo "Current version: $(VERSION)"
	@read -p "Enter new version (e.g., v0.2.0): " version; \
	git tag -a $$version -m "Release $$version"; \
	git push origin $$version; \
	echo "Tagged $$version."

clean: ## Clean up all generated files and binaries
	go clean
	rm -f $(GENERATE_MARKER)
	rm -rf $(BIN_DIR) result

migrate-up:
	migrate -database $(DB_URL) -path $(MIGRATE_DIR) up

migrate-down:
	migrate -database $(DB_URL) -path $(MIGRATE_DIR) down 1
