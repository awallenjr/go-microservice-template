# Start from the official Golang base image
FROM golang:1.22-alpine

# Install curl
RUN apk update && apk add --no-cache curl

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy everything from the current directory to the PWD(Present Working Directory) inside the container
COPY . .

# Download all the dependencies
RUN go mod download

# Build the Go app
RUN go build -o go-microservice-template .

# This container exposes port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./go-microservice-template"]
