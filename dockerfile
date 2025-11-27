# Dockerfile
FROM golang:1.19-alpine AS builder

# Instalar dependências de compilação
RUN apk add --no-cache git

# Definir diretório de trabalho
WORKDIR /app

# Copiar arquivos de dependências
COPY go.mod go.sum ./
RUN go mod download

# Copiar código fonte
COPY . .

# Compilar aplicação
RUN go build -o text-stats main.go

# Imagem final mínima
FROM alpine:latest

# Instalar dependências de runtime
RUN apk add --no-cache ca-certificates

# Criar usuário não-root
RUN adduser -D -s /bin/sh appuser
USER appuser

# Copiar binário da etapa de build
COPY --from=builder /app/text-stats /usr/local/bin/text-stats

# Definir ponto de entrada
ENTRYPOINT ["text-stats"]
CMD ["--help"]