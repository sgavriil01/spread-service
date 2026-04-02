# Spread Management Service

Small HTTP service written in Go for managing spreads for trading instruments.

## Requirements

- Go 1.21+
- Docker (optional, for containerized run)

## Build

Build the service:

```bash
go build ./cmd/server
```

## Run

Run the service:

```bash
go run ./cmd/server
```

The server starts on `http://localhost:8080`.

## Docker

Build the image:

```bash
docker build -t spread-service .
```

Run the container:

```bash
docker run --rm -p 8080:8080 spread-service
```

## Tests

Run unit tests:

```bash
go test ./...
```

## Endpoints

- `GET /health`
- `GET /symbols`
- `GET /spreads/{symbol}`
- `PATCH /spreads/{symbol}`
