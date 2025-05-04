ARG GO_VERSION=1.23
FROM golang:${GO_VERSION}-alpine AS builder

ENV CGO_ENABLED=0

WORKDIR /usr/src/app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN go build -o /run-app ./cmd/api

FROM alpine:3.21

COPY --from=builder /run-app /usr/local/bin/
CMD ["run-app"]
