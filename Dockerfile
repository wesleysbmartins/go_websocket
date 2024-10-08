FROM golang:latest AS builder

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o main ./cmd/main.go

EXPOSE 8080

CMD ["/app/main"]
