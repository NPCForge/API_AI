#!/bin/bash

echo "⏳ Waiting for PostgreSQL..."

until pg_isready -h postgres -p 5432 -U "API"; do
  sleep 2
done

echo "✅ PostgreSQL is ready. Launching app..."

exec "$@"
