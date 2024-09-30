# Use the official Golang image as the base image
FROM golang:1.19-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files first to take advantage of Docker caching
COPY go.mod go.sum ./

# Download Go module dependencies
RUN go mod download

RUN go mod tidy

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN go build -o main CMD/main/main.go

# Expose port 8080 to the host machine
EXPOSE 8080

# Set the binary as the entry point
CMD ["./main"]
