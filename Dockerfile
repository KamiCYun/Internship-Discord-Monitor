# === Stage 1: Build Python venv ===
FROM python:3.11-slim AS python-deps

WORKDIR /app

# Install venv tools and python alias
RUN apt-get update && \
    apt-get install -y python3-venv python-is-python3 && \
    rm -rf /var/lib/apt/lists/*

# Install Python packages into virtualenv
COPY requirements.txt .
RUN python -m venv /app/venv && \
    /app/venv/bin/pip install --no-cache-dir -r requirements.txt


# === Stage 2: Build Go binary ===
FROM golang:1.24 AS go-build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./cmd/main.go


# === Stage 3: Final runtime image ===
FROM debian:bullseye-slim

WORKDIR /app

# Install certs and symlink tools
RUN apt-get update && \
    apt-get install -y ca-certificates && \
    rm -rf /var/lib/apt/lists/*

# Copy Go binary and Python virtualenv
COPY --from=go-build /app/main /app/main
COPY --from=go-build /app/internal /app/internal
COPY --from=python-deps /app/venv /app/venv

# ✅ Set venv's Python as default
RUN ln -sf /app/venv/bin/python /usr/local/bin/python && \
    ln -sf /app/venv/bin/pip /usr/local/bin/pip

# ✅ Also prepend venv to PATH just in case
ENV PATH="/app/venv/bin:$PATH"

# Run Go binary
CMD ["./main"]
