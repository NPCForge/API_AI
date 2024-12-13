# Étape 1 : Utiliser une image Go officielle
FROM golang:1.23 AS builder

# Définir le répertoire de travail dans le conteneur
WORKDIR /app

# Copier les fichiers nécessaires dans le conteneur
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Compiler l'application
RUN go build -o main ./cmd/app/main.go

# Étape 2 : Utiliser une image minimale avec une version récente de glibc
FROM debian:bookworm-slim

# Installer les dépendances minimales
RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

# Définir le répertoire de travail
WORKDIR /app

# Copier l'exécutable compilé depuis l'étape précédente
COPY --from=builder /app/main .

# Copier les fichiers de configuration nécessaires (si requis)
COPY ./.env.local .env.local

# Exposer le port sur lequel l'application écoute (ex. : 8080)
EXPOSE 3000

# Commande pour démarrer l'application
CMD ["./main"]
