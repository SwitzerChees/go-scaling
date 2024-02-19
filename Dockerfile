# Use the official Go image to build the binary.
# This is based on Debian and sets the GOPATH to /go.
# https://hub.docker.com/_/golang
FROM golang:1.21.6 AS builder

# Set the working directory inside the container
WORKDIR /build

# Copy local code to the container image.
COPY main.go main.go

# Build the binary.
# -o flag sets the output file name
# CGO_ENABLED=0 disables CGO for a fully static binary
# GOOS=linux sets the target OS to Linux
RUN CGO_ENABLED=0 GOOS=linux go build -o /go-scaling ./main.go

# Use a minimal image to run the application
FROM alpine:3.14.10

WORKDIR /app

# Copy the binary to the production image from the builder stage.
COPY --from=builder /go-scaling /app/go-scaling

# Run the web service on container startup.
CMD ["./go-scaling"]
