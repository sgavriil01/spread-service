FROM golang:1.22-alpine

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .
RUN go build -o spread-service ./cmd/server

EXPOSE 8080

CMD ["./spread-service"]
