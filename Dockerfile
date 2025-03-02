# Stage 1: Build the Go binary using the official Golang image.
FROM golang:1.18-alpine AS builder

# Set the working directory inside the container.
WORKDIR /app

# Copy go.mod and go.sum files first to take advantage of Docker cache.
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire project into the container.
COPY . .

# Build the Go project. CGO is disabled for a statically-linked binary.
RUN CGO_ENABLED=0 go build -o main .

# Stage 2: Create a minimal image with the compiled binary.
FROM alpine:latest

# Install CA certificates (useful if your app makes HTTPS requests).
RUN apk --no-cache add ca-certificates

# Set working directory.
WORKDIR /root/

# Copy the binary from the builder stage.
COPY --from=builder /app/main .

# Expose the port the application listens on.
EXPOSE 8080

# Run the binary.
ENTRYPOINT ["./main"]
