# Use the official Go image for building
FROM golang:1.21 AS builder

# Set environment variables for cross-compilation
ENV GOOS=linux
ENV GOARCH=amd64

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum to leverage caching
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the application file
RUN go build -o main ./cmd/main.go

# Use a lightweight base image for the final build
FROM alpine:latest

# Install certificates (for HTTPS)
RUN apk --no-cache add ca-certificates

# Set the working directory inside the container
WORKDIR /root/

# Copy the built binary from the builder
COPY --from=builder /app/main .

# Verify the binary exists and has execution permissions
RUN chmod +x ./main

# Expose the port your application runs on
EXPOSE 8080

# Command to run the binary
CMD ["./main"]
