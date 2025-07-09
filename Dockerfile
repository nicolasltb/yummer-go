# Dockerfile for the Go backend

# Stage 1: Build the Go application
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum and download dependencies
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the application
RUN go build -o yummer-go .

# Stage 2: Create the final lightweight image
FROM alpine:latest

WORKDIR /app

# Copy the compiled application from the builder stage
COPY --from=builder /app/yummer-go .

# Expose the port the application listens on
EXPOSE 8080

# Command to run the application
CMD ["./yummer-go"]
