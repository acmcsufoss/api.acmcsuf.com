.DEFAULT_GOAL := build

BIN_DIR := bin
API_NAME := acmcsuf-api
CLI_NAME := acmcsuf-cli

MIGRATE_DIR := sql/migrations
DB_URL := sqlite3://dev.db

GENERATE_DEPS := $(wildcard internal/api/handlers/*.go)
GENERATE_MARKER := .generate.marker

VERSION := $(shell git describe --tags --always --dirty 2> /dev/null || echo "dev")

run: ## Build and run the api
	air

all: build
build: generate fmt api cli ## Build the api and cli binaries

api: ## Build the api binary
	@mkdir -p $(BIN_DIR)
	go build -ldflags "-X main.Version=$(VERSION)" -o $(BIN_DIR)/$(API_NAME) ./cmd/$(API_NAME)

cli: ## Build the cli binary
	@mkdir -p $(BIN_DIR)
	go build -ldflags "-X cli.Version=$(VERSION)" -o $(BIN_DIR)/$(CLI_NAME) ./cmd/$(CLI_NAME)

$(GENERATE_MARKER): $(GENERATE_DEPS)
	go generate ./...
	@touch $@

generate: $(GENERATE_MARKER) ## Generate all necessary files with go generate

fmt: ## Format all go files
	@go fmt ./...

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

clean: ## Clean up binaries and build artifacts
	go clean
	rm -f $(GENERATE_MARKER)
	rm -rf $(BIN_DIR) result

migrate-up: ## Perform database migration up
	migrate -database $(DB_URL) -path $(MIGRATE_DIR) up

migrate-down: ## Perform database migration down
	migrate -database $(DB_URL) -path $(MIGRATE_DIR) down 1

help: ## Display this help screen
	@grep -hE '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: help fmt run build all api cli vet check test check-sql fix-sql clean release generate migrate-up migrate-down
