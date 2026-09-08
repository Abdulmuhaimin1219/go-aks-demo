# ---------- Build stage ----------
FROM golang:1.25-alpine AS builder

WORKDIR /app

# Download dependencies first for better Docker layer caching
COPY go.mod ./
RUN go mod download

# Copy application source
COPY . .

# Build a static Linux binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags="-s -w" -o /server ./cmd/server


# ---------- Runtime stage ----------
FROM gcr.io/distroless/static-debian12:nonroot

WORKDIR /

COPY --from=builder /server /server

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/server"]