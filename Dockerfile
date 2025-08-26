FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -o wallet-api main.go

FROM alpine:3.20

WORKDIR /app

COPY --from=builder /app/wallet-api .

EXPOSE 8080

CMD ["./wallet-api"]