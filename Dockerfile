# Use the official Golang image as the base image
FROM golang:1.18

# Set the working directory inside the container
WORKDIR /app

# Copy the source code from the host to the container
COPY . .

# Download the project dependencies
RUN go mod download

# Build the Go application
RUN go build -o simple-auth-v2 .

# Set environment variables for MongoDB connection
ENV DB_ADDRESS mongodb://mongodb-service:27017
ENV MONGO_PORT 27017
ENV MONGO_DB auth_db

# Expose the port the application listens on
EXPOSE 8080

# Run the Go application
CMD ["./simple-auth-v2"]
