# Use an official Golang runtime as a parent image
FROM golang:1.21.3 AS build

# Set the working directory inside the container
WORKDIR /go/src/service-gateway

# Copy the Go application code into the container
COPY . .

# Build the Go application
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /go/bin/service-gateway

# Start a new stage for a smaller runtime image
FROM alpine:3.14

# Set the working directory inside the container
WORKDIR /app

# Copy the built Go application from the previous stage
COPY --from=build /go/bin/service-gateway .

# Expose a port that the application will run on
EXPOSE 8000

# Run the Go application
CMD ["./service-gateway"]