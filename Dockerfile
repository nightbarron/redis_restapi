# Compile stage
FROM golang:1.19 AS builder

RUN mkdir /app

WORKDIR /app

# Add the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Add the rest of the code
COPY . .

# Build the application
RUN CGO_ENABLED=0 go build -o /app/main .

# Final stage
FROM redis:alpine

RUN mkdir /app

WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/main /app/

EXPOSE 8080
EXPOSE 6379

# Run the binary
CMD ["/app/main"]