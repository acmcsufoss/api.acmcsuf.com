ARG GO_VERSION=1
FROM golang:${GO_VERSION}-bookworm AS builder

WORKDIR /usr/src/app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN go build -v -o /run-app ./cmd/acmcsuf-api


FROM debian:bookworm

COPY --from=builder /run-app /usr/local/bin/
ENV GIN_MODE=release
ENV ENV=development
ENV PORT=80
# Staging server
ENV GUILD_ID=1446729071777021966
EXPOSE 80
CMD ["run-app"]
