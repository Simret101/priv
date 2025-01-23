# Use the official Golang image as a base image
FROM golang:1.23-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod tidy

# Copy the source code into the container
COPY . .

# Build the Go application
RUN go build -o userservice .

# Expose the port your application will run on
EXPOSE 8083

# Command to run your app
CMD ["./userservice"]
