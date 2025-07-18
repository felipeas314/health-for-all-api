# Etapa de build
FROM golang:1.24 AS builder

WORKDIR /app

# Copia arquivos do Go
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copia o projeto inteiro
COPY . .

# Compila o binário principal
RUN CGO_ENABLED=0 GOOS=linux go build -o app ./cmd/api/main.go

# Etapa de deploy (imagem final leve)
FROM alpine:latest

WORKDIR /root/
COPY --from=builder /app/app .

# Porta usada no App Runner (ajuste se necessário)
EXPOSE 8080

# Comando de entrada
CMD ["./app"]
