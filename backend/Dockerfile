FROM golang:1.24-bookworm as builder

WORKDIR /app

# Install system dependencies for CGO and SQLite
RUN apt-get update && apt-get install -y gcc libc6-dev

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Important: CGO must be enabled for go-sqlite3
ENV CGO_ENABLED=1

RUN go build -o server .

# Final stage: lightweight image with just the binary and DB
FROM debian:bookworm-slim

WORKDIR /app

RUN apt-get update && apt-get install -y libc6 ca-certificates && rm -rf /var/lib/apt/lists/*

COPY --from=builder /app/server .
COPY data/thoth.db ./data/thoth.db
VOLUME [ "/data" ]

EXPOSE 8080

CMD ["./server"]
