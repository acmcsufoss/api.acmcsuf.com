ARG GO_VERSION=1.23-alpine
FROM golang:${GO_VERSION} AS builder

ENV CGO_ENABLED=0
ENV GOOS=linux
RUN mkdir -p /app
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -ldflags="-w -s" -o /app/api .

FROM alpine:3.21
WORKDIR /app
COPY --from=builder /app/api .
EXPOSE 8080
CMD ["/app/api"]
