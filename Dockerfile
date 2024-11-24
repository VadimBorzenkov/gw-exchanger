FROM golang:1.22-alpine3.19 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main ./cmd

FROM alpine:3.19
WORKDIR /app
COPY --from=builder /app/main .
CMD ["./main"]
