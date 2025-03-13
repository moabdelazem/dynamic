# Description: Dockerfile for building the go application

# Use the official Golang image to create a build artifact.
# Using Go 1.23 alpine image
FROM golang:1.23-alpine3.21 AS builder        

# Create and change to the app directory.
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the application
RUN go build -o main ./cmd/main.go  

# Use a minimal alpine image
FROM alpine:3.14

# Copy the binary to the production image from the builder stage
COPY --from=builder /app/main /app/main

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./app/main"]