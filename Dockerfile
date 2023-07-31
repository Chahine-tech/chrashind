# Use the official Golang image as the base image
FROM golang:1.20-alpine AS builder

# Set the working directory inside the container
WORKDIR /go/src/app

# Copy the Go module files to the working directory
COPY go.mod go.sum ./

# Download and cache Go modules
RUN go mod download

# Copy the rest of the application source code to the working directory
COPY . .

# Generate the Prisma client
RUN go run github.com/steebchen/prisma-client-go generate

# Build the Go application
RUN go build -o main .

# Expose the port that the GraphQL server will run on
EXPOSE 8080

# Set the command to run the application
CMD ["./main"]
