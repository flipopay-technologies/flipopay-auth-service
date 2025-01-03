# Stage 1: Build the Go application
FROM golang:1.20-alpine AS builder

# Set the working directory inside the container
WORKDIR /auth-service

# Copy go mod and sum files to leverage Docker cache
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go application
RUN go build -o app .

# Stage 2: Create a lightweight image for production
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /root/

# Copy the built application from the builder stage
COPY --from=builder /auth-service/app .

# Expose the port your application runs on
EXPOSE 8080

# Run the built application
CMD ["./app"]
