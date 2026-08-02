ARG GO_VERSION=1
FROM golang:${GO_VERSION}-bookworm AS builder

WORKDIR /usr/src/app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN go build -v -o /run-app ./cmd/acmcsuf-api


FROM debian:bookworm

COPY --from=builder /run-app /usr/local/bin/
RUN mkdir -p /app/data
WORKDIR /app
COPY --from=builder /usr/src/app/sql/migrations ./sql/migrations
ENV DATABASE_URL=file:/app/data/acmcsuf.db?cache=shared&mode=rwc
ENV GIN_MODE=debug
ENV ENV=development
ENV PORT=80
EXPOSE 80
CMD ["run-app"]
