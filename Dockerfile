# Stage 1: Build the Go app
FROM golang:1.23 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o main .

# Stage 2: Create a smaller image for the final build
FROM debian:bookworm-slim

# Set the working directory inside the container
WORKDIR /app

# Copy the pre-built binary file from the previous stage
COPY --from=builder /app/main .

# Expose port 80 to the outside world
EXPOSE 80

# Command to run the executable
CMD ["./main"]
