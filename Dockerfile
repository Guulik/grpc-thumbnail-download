FROM golang:latest AS builder
WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o /app/server ./cmd/thumbnail/main.go

EXPOSE 500

CMD ["./server"]