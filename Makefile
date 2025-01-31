.DEFAULT_GOAL := run

GENERATE_DEPS := $(wildcard sql/*.sql) sqlc.yaml
GENERATE_MARKER := .generate.marker

.PHONY:fmt vet run build clean test

fmt:
	@go fmt ./...

$(GENERATE_MARKER): $(GENERATE_DEPS)
	go generate ./...
	@touch $@

generate: fmt $(GENERATE_MARKER)

vet: fmt
	go vet ./...

run: vet generate
	go run cmd/api/main.go

build: vet generate
	go build cmd/api/main.go

test: vet
	go test ./...

clean:
	go clean
	rm -f $(GENERATE_MARKER)
