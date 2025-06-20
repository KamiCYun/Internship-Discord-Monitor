# === Stage 1: Build Python venv ===
FROM python:3.11-slim AS python-deps

WORKDIR /app

# Install pip, venv, and the 'python' alias
RUN apt-get update && \
    apt-get install -y python3-venv python-is-python3 && \
    rm -rf /var/lib/apt/lists/*

# Copy requirements and install into virtual environment
COPY requirements.txt .
RUN python3 -m venv /app/venv && \
    /app/venv/bin/pip install --no-cache-dir -r requirements.txt


# === Stage 2: Build Go Binary ===
FROM golang:1.24 AS go-build

WORKDIR /app

# Copy go.mod files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application
COPY . .

# Build static binary (no glibc dependency)
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./cmd/main.go


# === Stage 3: Final minimal image ===
FROM debian:bullseye-slim

WORKDIR /app

# Install ca-certificates and python alias
RUN apt-get update && \
    apt-get install -y ca-certificates python-is-python3 && \
    rm -rf /var/lib/apt/lists/*

# Copy Go binary and venv
COPY --from=go-build /app/main /app/main
COPY --from=go-build /app/internal /app/internal
COPY --from=python-deps /app/venv /app/venv

# Optional: Add venv to PATH
ENV PATH="/app/venv/bin:$PATH"

# Run the Go application
CMD ["./main"]
