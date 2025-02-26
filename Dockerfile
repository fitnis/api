# ------------------------------------------------------------
# 1. Dev Stage (with Air + Delve)
# ------------------------------------------------------------
FROM golang:1.24-alpine AS dev

RUN apk add --no-cache git bash

WORKDIR /app

# Install Air & Delve
RUN go install github.com/air-verse/air@latest
RUN go install github.com/go-delve/delve/cmd/dlv@latest

# Copy go.mod and go.sum to cache dependencies
COPY go.mod go.sum ./
RUN go mod download

# Expose API port & debug port
EXPOSE 3000 2345

CMD ["air", "-c", ".air.toml"]

# ------------------------------------------------------------
# 2. BUILDER Stage (Production Binary)
# ------------------------------------------------------------
FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build the production binary
RUN go build -o /api cmd/api/main.go

# ------------------------------------------------------------
# 3. PRODUCTION Stage (Minimal Final Image)
# ------------------------------------------------------------
FROM alpine:latest AS production

# Create a non-root user
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

WORKDIR /app

# Copy the compiled binary from builder
COPY --from=builder /api ./api

# Expose the API port
EXPOSE 3000

USER appuser

CMD [\"./api\"]
