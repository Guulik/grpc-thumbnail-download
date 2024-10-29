FROM golang:latest AS builder
WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o /app/thumbnail-server ./cmd/thumbnail

EXPOSE 50051

CMD ["/app/thumbnail-server"]
