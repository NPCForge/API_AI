#!/usr/bin/env bash
set -euo pipefail

PGHOST="${PGHOST:-postgres}"
PGPORT="${PGPORT:-5432}"
PGUSER="${PGUSER:-API}"

echo "⏳ Waiting for PostgreSQL at ${PGHOST}:${PGPORT} (user: ${PGUSER})..."

until pg_isready -h "$PGHOST" -p "$PGPORT" -U "$PGUSER" >/dev/null 2>&1; do
    echo "  ...still waiting"
    sleep 2
done

echo "✅ PostgreSQL is ready. Launching app..."

exec "$@"
