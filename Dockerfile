# Start from the latest golang base image.
FROM golang:latest as builder

# Add Maintainer Info.
LABEL maintainer="Nikolajus Karpovas <mwsoftofficial@gmail.com>"

# Create build directory.
RUN mkdir build

# Copy project to build directory.
COPY . /build

# Set build as working directory.
WORKDIR /build/cmd/media

# Fetch dependencies.
RUN go mod download

# Build the Go app.
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o main

# Create unprivelleged user.
RUN adduser --disabled-login appuser

# Start a new stage from scratch.
FROM alpine:latest

RUN apk update && apk add --no-cache git ca-certificates && update-ca-certificates

# Import the user and group files from the builder.
COPY --from=builder /etc/passwd /etc/passwd

# Create build directory.
RUN mkdir app

# Copy the pre-built binary file from the previous stage.
COPY --from=builder /build/cmd/media/main /app/

# Set working directory in current stage.
WORKDIR /app

# Use an unprivileged user.
USER appuser

# Expose port 7100.
EXPOSE 7100 7100

# Command to run the executable.
CMD ["./main"]