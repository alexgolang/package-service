# Build stage
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o package-service ./cmd/main.go

# Final stage
FROM alpine:latest

# Add necessary packages
RUN apk update && \
    apk add --no-cache \
    ca-certificates \
    tzdata \
    curl

RUN adduser -D appuser
USER appuser

WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/package-service .

# Expose port
EXPOSE 8080

# Set environment variables (can be overridden at runtime)
ENV HTTP_SERVER_PORT=8080
# 23,31,53
# 250,500,1000,2000,5000
ENV PACKAGE_SIZES=250,500,1000,2000,5000

# Run the application
CMD ["./package-service"]