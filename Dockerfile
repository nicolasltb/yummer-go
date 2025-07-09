# Dockerfile for the Go backend

FROM golang:1.24-alpine

# Install build tools for CGO (required by go-sqlite3)
RUN apk add --no-cache gcc musl-dev

WORKDIR /app

# Copy go.mod and go.sum and download dependencies
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the application (CGO is enabled by default here)
RUN go build -o yummer-go .

# Expose the port the application listens on
EXPOSE 8080

# Command to run the application
CMD ["./yummer-go"]