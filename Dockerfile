FROM golang:1.21 AS builder

WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o app ./cmd/server

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/app .
EXPOSE 9000
CMD ["./app"]
