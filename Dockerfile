# Use the official Golang image as the base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the Go server code into the container
COPY . .

# Build the Go server
RUN go build -o server .

# Expose the port that the server will listen on
EXPOSE 8080

# Set the command to run the server when the container starts
CMD ["./server"]