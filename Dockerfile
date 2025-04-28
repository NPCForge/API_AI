# Étape 1 : Utiliser une image Go officielle
FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o main ./cmd/app/main.go

# Étape 2 : Image finale minimale
FROM debian:bookworm-slim

# Installer ca-certificates et pg_isready (postgresql-client)
RUN apt-get update && apt-get install -y \
    ca-certificates \
    postgresql-client \
 && rm -rf /var/lib/apt/lists/*

WORKDIR /app

# Copier l'exécutable et les assets depuis le builder
COPY --from=builder /app/main .
COPY --from=builder /app/config/asset config/asset
COPY --from=builder /app/prompts prompts

# Copier les fichiers de config (env local)
COPY ./.env .env

# Ajouter le script d'entrypoint
COPY ./entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

EXPOSE 3000

ENTRYPOINT ["/entrypoint.sh"]
CMD ["./main"]
