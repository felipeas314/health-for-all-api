# Etapa de build
FROM golang:1.24.2 AS builder

WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

# Altere aqui se o main.go estiver em outro local
RUN CGO_ENABLED=0 GOOS=linux go build -o app ./cmd/api/main.go

# Etapa de deploy
FROM alpine:latest

WORKDIR /root/
COPY --from=builder /app/app .

EXPOSE 8080

CMD ["./app"]
