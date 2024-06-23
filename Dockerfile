# Builder stage
FROM --platform=linux/arm64 golang:alpine as builder
WORKDIR /src
COPY . .
# Add necessary build tools
RUN apk add --no-cache build-base
# Compile for ARM64
ENV GOOS=linux 
ENV GOARCH=arm64
# Disable CGO to avoid any libc dependencies that might conflict
RUN CGO_ENABLED=0 go build -o main .

# Final stage
FROM --platform=linux/arm64 alpine:latest as base
# Create user and group
RUN addgroup -S api && adduser -S -G api api
# Set the working directory and copy the binary
WORKDIR /app
COPY --from=builder /src/main /app/product-api
COPY conf.json /app/
# Ensure the binary is executable
RUN chmod +x /app/product-api
# Set entrypoint to the binary
ENTRYPOINT ["/app/product-api"]