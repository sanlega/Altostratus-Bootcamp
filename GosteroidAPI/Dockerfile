# Build
FROM golang:1.22.3-alpine3.20 as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Run
FROM alpine:latest

WORKDIR /root/

RUN apk add --no-cache ca-certificates tzdata

COPY --from=builder /app/main .
COPY --from=builder /app/.env .

EXPOSE 8080

CMD ["./main"]
