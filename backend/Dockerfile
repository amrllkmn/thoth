FROM golang:1.24-alpine as builder

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o ./bin/server .

FROM alpine:latest

WORKDIR /app

RUN apk --no-cache add ca-certificates

COPY --from=builder /app/bin/server .
COPY thoth.db ./thoth.db

EXPOSE 8080

CMD [ "./server" ]