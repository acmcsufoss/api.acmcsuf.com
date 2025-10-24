.DEFAULT_GOAL := build

BIN_DIR := bin
API_NAME := acmcsuf-api
CLI_NAME := acmcsuf-cli

GENERATE_DEPS := $(wildcard internal/db/sql/schemas/*.sql) $(wildcard internal/db/sql/queries/*.sql) internal/db/sqlc.yaml $(wildcard internal/api/handlers/*.go)
GENERATE_MARKER := .generate.marker

.PHONY:fmt run build vet check test check-sql fix-sql clean

fmt:
	@go fmt ./...

VERSION := $(shell git describe --tags --always --dirty 2> /dev/null || echo "dev")

$(GENERATE_MARKER): $(GENERATE_DEPS)
	go generate ./...
	@touch $@

generate: fmt $(GENERATE_MARKER)

run: build
	./$(BIN_DIR)/$(API_NAME)

build: generate
	@mkdir -p $(BIN_DIR)
	go build -ldflags "-X main.Version=$(VERSION)" -o $(BIN_DIR)/$(API_NAME) ./cmd/$(API_NAME)
	go build -ldflags "-X main.Version=$(VERSION)" -o $(BIN_DIR)/$(CLI_NAME) ./cmd/$(CLI_NAME)

vet:
	go vet ./...

check:
	staticcheck ./...

test: check
	go test ./...

check-sql:
	sqlfluff lint --dialect sqlite

fix-sql:
	sqlfluff fix --dialect sqlite

release:
	@echo "Current version: $(VERSION)"
	@read -p "Enter new version (e.g., v0.2.0): " version; \
	git tag -a $$version -m "Release $$version"; \
	git push origin $$version; \
	echo "Tagged $$version."

clean:
	go clean
	rm -f $(GENERATE_MARKER)
	rm -rf $(BIN_DIR) result
