FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o main ./cmd/app/main.go

FROM debian:bookworm-slim
RUN apt-get update && apt-get install -y \
    ca-certificates \
    postgresql-client \
    bash \
 && rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY --from=builder /app/main /app/main
COPY --from=builder /app/config/asset /app/config/asset
COPY --from=builder /app/prompts /app/prompts
COPY ./.env /app/.env

COPY --chmod=755 entrypoint.sh /entrypoint.sh
RUN sed -i 's/\r$//' /entrypoint.sh

EXPOSE 3000

ENTRYPOINT ["/bin/bash", "/entrypoint.sh"]
CMD ["/app/main"]