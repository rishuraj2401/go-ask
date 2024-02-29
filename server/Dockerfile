# Use the official Go base image
FROM golang:1.22.0

# Set the working directory inside the container
WORKDIR /app

# Copy the Go application source code into the container
COPY . .

# Download and install any dependencies
RUN go mod download

# Build the Go application
RUN go build -o main

# Expose the port the application runs on
EXPOSE 8089

# Define the command to run the application when the container starts
CMD ["./main"]
