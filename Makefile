.DEFAULT_GOAL := build

BIN_DIR := bin
API_NAME := acmcsuf-api
CLI_NAME := acmcsuf-cli

GO_SOURCES := $(shell find . -type f -name '*.go' -not -path '*/vendor/*')
GO_DEPS := $(GO_SOURCES) go.mod go.sum

MIGRATE_DIR := sql/migrations
DB_URL := sqlite3://dev.db

GENERATE_DOCS_DEPS := $(wildcard internal/api/handlers/*.go)
SQLC_DEPS := $(wildcard sql/migrations/*.sql) $(wildcard sql/queries/*.sql)
SQLC_TARGET := internal/api/dbmodels

VERSION := $(shell git describe --tags --always --dirty 2> /dev/null || echo "dev")

run: ## Build and run the api
	air

all: build
build: generate fmt api cli ## Build the api and cli binaries

api: $(BIN_DIR)/$(API_NAME) ## Build the api binary

$(BIN_DIR)/$(API_NAME): $(GO_DEPS)
	@mkdir -p $(BIN_DIR)
	go build -ldflags "-X main.Version=$(VERSION)" -o $(BIN_DIR)/$(API_NAME) ./cmd/$(API_NAME)

cli: $(BIN_DIR)/$(CLI_NAME) ## Build the cli binary

$(BIN_DIR)/$(CLI_NAME): $(GO_DEPS)
	@mkdir -p $(BIN_DIR)
	go build -ldflags "-X cli.Version=$(VERSION)" -o $(BIN_DIR)/$(CLI_NAME) ./cmd/$(CLI_NAME)

# generate: generate-docs generate-sqlc ## Generate all necessary files with go generate
$(SQLC_TARGET): $(SQLC_DEPS)
	sqlc generate

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
