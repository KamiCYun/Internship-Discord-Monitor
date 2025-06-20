# === Stage 1: Python Dependencies ===
FROM python:3.11-slim AS python-deps

WORKDIR /app

# Create venv and install requirements
COPY requirements.txt .
RUN python3 -m venv /app/venv && \
    /app/venv/bin/pip install --no-cache-dir -r requirements.txt

# === Stage 2: Build Go Binary ===
FROM golang:1.24 AS go-build

WORKDIR /app

# Copy Go mod files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy rest of the source
COPY . .

# Build Go binary
RUN go build -o main ./cmd/main.go

# === Stage 3: Final Image ===
FROM debian:bullseye-slim

# Install Python and copy venv from Stage 1
RUN apt-get update && apt-get install -y python3 && rm -rf /var/lib/apt/lists/*
COPY --from=python-deps /app/venv /app/venv

# Copy Go binary from Stage 2
COPY --from=go-build /app/main /app/main
COPY --from=go-build /app/internal /app/internal

WORKDIR /app

# Run the Go app
CMD ["./main"]
