# Use an official Golang runtime as a parent image
FROM golang:1.19.6-alpine AS builder

# Set the working directory to /go/src/app
WORKDIR /usr/local/go/src/app

# Copy the current directory contents into the container at /go/src/app
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy all source codes into the image
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -o /app


FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app ./

# Set the default command to run when the container starts
CMD ["./app"]
