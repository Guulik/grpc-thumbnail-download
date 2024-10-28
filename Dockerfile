FROM golang:latest AS builder
WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o grpc-server ./cmd/thumbnail
RUN go build -o cli ./cmd/cli

# Финальный контейнер
FROM alpine:latest

COPY --from=builder grpc-server grpc-server
COPY --from=builder cli cli

CMD ["/usr/local/bin/server"]