# Build stage
FROM golang:1.24.5-alpine AS builder
WORKDIR /app

# Install air for hot-reloading
RUN go install github.com/air-verse/air@latest

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Development stage (with air for hot-reload)
FROM builder AS dev
CMD ["air"]

# Production stage
FROM builder AS prod
RUN go build -o /app/main .
CMD ["/app/main"]
