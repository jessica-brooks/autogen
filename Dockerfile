FROM golang:1.19-alpine AS builder

# Move to working directory (/build).
WORKDIR /build

# Copy and download dependency using go mod.
COPY go.mod go.sum ./
RUN go mod download


COPY .env.template ./.env

# Copy the code into the container.
COPY ./main.go .
COPY ./api ./api

# Set-up the build configuration for the go build.
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

# Build the Go app
RUN go build -ldflags="-s -w" -o app .

# Start a new stage from scratch
FROM gcr.io/distroless/static

# Expose port 8080 to the outside world
EXPOSE 8080

# Copy the Pre-built binary file from the previous stage.
COPY --from=builder ["/build/app", "/"]
ENTRYPOINT ["/app"]
