ARG GO_VERSION=1
FROM golang:${GO_VERSION}-bookworm AS builder

WORKDIR /usr/src/app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN go build -v -o /run-app ./cmd/acmcsuf-api


FROM debian:bookworm

COPY --from=builder /run-app /usr/local/bin/
RUN apt-get update \
    && apt-get install -y --no-install-recommends ca-certificates \
    && rm -rf /var/lib/apt/lists/* \
    && mkdir -p /app/data
WORKDIR /app
COPY --from=builder /usr/src/app/sql/migrations ./sql/migrations
EXPOSE 8080
CMD ["run-app"]
