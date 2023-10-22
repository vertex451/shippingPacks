# Use an official Golang runtime as a parent image
FROM golang:1.21

# Set the working directory to /app
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . /app

# Build the Go app
RUN go build -o main cmd/main.go

# Command to run the executable
CMD ["./main"]
