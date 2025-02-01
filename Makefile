.DEFAULT_GOAL := run

GENERATE_DEPS := $(wildcard internal/db/sqlite/*.sql) $(wildcard internal/db/sqlc.yaml)
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

run: generate
	go run cmd/api/main.go

build: generate
	go build cmd/api/main.go

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
