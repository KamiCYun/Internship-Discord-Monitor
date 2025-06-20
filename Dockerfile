# === Stage 1: Build Python venv ===
FROM python:3.11-slim AS python-deps

WORKDIR /app

# Install pip and venv tools
RUN apt-get update && apt-get install -y python3-venv && rm -rf /var/lib/apt/lists/*

# Copy requirements and install into venv
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

# Build a static binary (no glibc dependency)
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./cmd/main.go


# === Stage 3: Final minimal image ===
FROM debian:bullseye-slim

WORKDIR /app

# Copy built Go binary and Python venv
COPY --from=go-build /app/main /app/main
COPY --from=python-deps /app/venv /app/venv
COPY --from=go-build /app/internal /app/internal

# # Set environment for Python if needed (optional)
# ENV PATH="/app/venv/bin:$PATH"

# Run the Go application
CMD ["./main"]
