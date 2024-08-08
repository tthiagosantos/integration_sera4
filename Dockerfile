# Etapa base para construir o binário
FROM golang:1.22-alpine3.18 as base

WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server ./cmd/main.go

# Etapa final para criar a imagem executável
FROM alpine:3.18 as binary

# Copia o binário
COPY --from=base /app/server /server

# Copia o arquivo .env
COPY --from=base /app/.env /.env

EXPOSE 8080
CMD ["/server"]
