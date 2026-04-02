# Spread Management Service

Small HTTP service written in Go for managing spreads for trading instruments.

## Requirements

- Go 1.21+

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

## Endpoints

- `GET /health`
- `GET /symbols`
- `GET /spreads/{symbol}`
- `PATCH /spreads/{symbol}`
