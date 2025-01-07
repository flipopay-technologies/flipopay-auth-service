# Use the official Go image for building
FROM golang:1.21 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum to leverage caching
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the application
RUN go build -o main ./cmd

# Use a lightweight base image for the final build (Alpine for production)
FROM alpine:latest

# Install certificates (for HTTPS)
RUN apk --no-cache add ca-certificates

# Set the working directory inside the container
WORKDIR /root/

# Copy the built binary from the builder
COPY --from=builder /app/main .

# Expose the port your application runs on (e.g., 8080)
EXPOSE 8080

# Command to run the binary
CMD ["./main"]
