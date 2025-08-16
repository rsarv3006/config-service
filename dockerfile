# Use the official Golang image as the base image
FROM golang:1.24-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code and .env file into the container
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Use a smaller base image for the final stage
FROM alpine:latest

# Set the working directory
WORKDIR /root/

# Copy the binary and .env file from the builder stage
COPY --from=builder /app/main .
# COPY --from=builder /app/.env .

# Expose port 3000
EXPOSE 3000

# Command to run the application
CMD ["./main"]
