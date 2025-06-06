# Build stage
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o task-scheduler ./cmd/main.go

# Final stage
FROM alpine:latest

WORKDIR /app

# Copy the binary from builder
COPY --from=builder /app/task-scheduler .

# Expose the gRPC and metrics ports
EXPOSE 50053
EXPOSE 2113

# Run the application
CMD ["./task-scheduler"]
