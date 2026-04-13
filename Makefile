.DEFAULT_GOAL := build

BIN_DIR := bin
API_NAME := acmcsuf-api
CLI_NAME := acmcsuf-cli

GO_SOURCES := $(shell find . -type f -name '*.go' -not -path '*/vendor/*')
GO_DEPS := $(GO_SOURCES) go.mod go.sum

SQLC_DEPS := $(wildcard sql/migrations/*.sql) $(wildcard sql/queries/*.sql)
SQLC_TARGET := internal/api/dbmodels/models.go

VERSION := $(shell git describe --tags --always --dirty 2> /dev/null || echo "dev")

run: ## Build and run the api
	air

all: generate fmt build
build: api cli ## Build the api and cli binaries

api: $(BIN_DIR)/$(API_NAME) ## Build the api binary

$(BIN_DIR)/$(API_NAME): $(GO_DEPS) sqlc
	@mkdir -p $(BIN_DIR)
	go build -x -ldflags "-X main.Version=$(VERSION)" -o $(BIN_DIR)/$(API_NAME) ./cmd/$(API_NAME)

cli: $(BIN_DIR)/$(CLI_NAME) ## Build the cli binary

$(BIN_DIR)/$(CLI_NAME): $(GO_DEPS)
	@mkdir -p $(BIN_DIR)
	go build -x -ldflags "-X cli.Version=$(VERSION)" -o $(BIN_DIR)/$(CLI_NAME) ./cmd/$(CLI_NAME)

generate: swag sqlc ## Generate all necessary files

swag: ## Generate OpenAPI docs
	swag init -d  cmd/acmcsuf-api,internal/api/handlers,internal/api/dbmodels -o internal/api/docs --parseDependency

sqlc: $(SQLC_TARGET) ## Generate dbmodels package with sqlc

$(SQLC_TARGET): $(SQLC_DEPS)
	sqlc generate

fmt: ## Format all go files
	@go fmt ./...

check: ## Run static analysis on all go files
	staticcheck -f stylish ./...

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
	rm -rf $(BIN_DIR) result

help: ## Display this help screen
	@grep -hE '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: help fmt run build all api cli check test check-sql fix-sql clean release generate
