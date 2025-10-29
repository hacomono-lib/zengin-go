# syntax=docker/dockerfile:1

# Build stage
FROM golang:1.23-bookworm AS builder

# Install development tools
RUN apt-get update && apt-get install -y \
    git \
    make \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /workspace

# Copy go mod files
COPY go.mod go.sum* ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Development stage
FROM golang:1.23-bookworm AS development

# Install development and debugging tools
RUN apt-get update && apt-get install -y \
    git \
    make \
    vim \
    curl \
    ca-certificates \
    && rm -rf /var/lib/apt/lists/*

# Install Go tools
RUN go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest && \
    go install github.com/securego/gosec/v2/cmd/gosec@latest

WORKDIR /workspace

# Copy go mod files
COPY go.mod go.sum* ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Default command
CMD ["bash"]

# Test stage
FROM development AS test

# Run tests
RUN make test

# Production stage (for building examples or tools)
FROM golang:1.23-alpine AS production

WORKDIR /app

# Copy go mod files
COPY go.mod go.sum* ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build example
RUN cd example && go build -o /app/example main.go

# Final minimal image
FROM alpine:latest AS final

RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=production /app/example .

CMD ["./example"]

