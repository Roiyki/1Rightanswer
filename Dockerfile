# Use the official Golang image as a base
FROM golang:latest AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY main.go .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o main .

# Start a new stage from scratch
FROM alpine:latest  

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .

# Copy the HTML file into the container
COPY templates/index.html .

# Expose port 8080 to the outside world
EXPOSE 3000:3000

# Command to run the executable
CMD ["./main"]
