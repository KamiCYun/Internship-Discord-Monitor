# Use official Go image
FROM golang:1.22

# Install Python and pip
RUN apt-get update && apt-get install -y python3 python3-pip

# Set working directory inside the container
WORKDIR /app

# Copy Go files
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of your files
COPY . .

# Install Python dependencies
RUN pip3 install --no-cache-dir -r requirements.txt

# Build Go binary from ./cmd/main.go
RUN go build -o app ./cmd/main.go

# Expose port (match your app)
EXPOSE 8080

# Run the Go binary
CMD ["./app"]
