# Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /go/src/wallet

# Copy only the necessary files for fetching dependencies
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download && go mod tidy

# Copy the entire source code
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o /usr/local/bin/server ./cmd/fiber

# Final stage
FROM alpine:latest

RUN apk update && apk add bash

# Install MySQL client and migrate tool
RUN apk --no-cache add mysql-client && \
    wget -O /tmp/migrate.tar.gz https://github.com/golang-migrate/migrate/releases/download/v4.17.0/migrate.linux-amd64.tar.gz && \
    tar -zxvf /tmp/migrate.tar.gz -C /usr/local/bin/ && \
    rm /tmp/migrate.tar.gz && \
    chmod +x /usr/local/bin/migrate

# Run migrations during the build process
# RUN migrate -path=/migrations -database "${MYSQL_DSN}" up

# Set the working directory
WORKDIR /usr/local/bin

# Copy only the built binary and migrations
COPY --from=builder /usr/local/bin/server .
COPY ./database/migrations /migrations

# Expose the necessary port
EXPOSE 80

# Entrypoint command
CMD ["./server"]
