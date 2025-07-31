# Stage 1: Build
FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /routify ./cmd/api

# Stage 2: Runtime
FROM alpine:latest

COPY --from=builder /routify /routify

EXPOSE 50051

ENTRYPOINT ["/routify"]
