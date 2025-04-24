.DEFAULT_GOAL := build

BIN_DIR := bin
APP_NAME := api

GENERATE_DEPS := $(wildcard internal/db/sql/schemas/*.sql) $(wildcard internal/db/sql/queries/*.sql) $(wildcard sqlc.yaml)
GENERATE_MARKER := .generate.marker

.PHONY:fmt vet run build check test sql-check sql-fix clean

fmt:
	@go fmt ./...

$(GENERATE_MARKER): $(GENERATE_DEPS)
	go generate ./...
	@touch $@

generate: fmt $(GENERATE_MARKER)

vet: fmt
	go vet ./...

run: build
	./$(BIN_DIR)/$(APP_NAME)

build: generate
	mkdir -p $(BIN_DIR)
	go build -o $(BIN_DIR)/$(APP_NAME) ./cmd/api

check: vet
	nilaway ./...

test: check
	go test ./...

sql-check:
	sqlfluff lint --dialect sqlite

sql-fix:
	sqlfluff format --dialect sqlite
	sqlfluff fix --dialect sqlite

clean:
	go clean
	rm -f $(GENERATE_MARKER)
	rm -rf $(BIN_DIR)
