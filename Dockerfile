FROM golang:1.26 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main cmd/main.go

FROM ubuntu:latest
WORKDIR /app
COPY --from=builder /app/main .
CMD ["./main"]