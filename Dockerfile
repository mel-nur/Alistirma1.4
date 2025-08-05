# Build stage
FROM golang:1.24-alpine AS builder
WORKDIR /app

COPY . .

RUN go build -o /my-golang-app

# Final stage
FROM alpine:latest

WORKDIR /

COPY --from=builder /my-golang-app .
COPY ./static ./static

EXPOSE 8080

CMD ["./my-golang-app"]
