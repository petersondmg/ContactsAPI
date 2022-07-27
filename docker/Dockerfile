FROM golang:1.18-alpine AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o api cmd/api/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/api .
CMD ["/app/api"]  