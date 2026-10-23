#!/usr/bin/env bash
# Script to apply Supabase migrations to PostgreSQL database
# Usage: ./scripts/apply-migrations.sh "<POSTGRES_DATABASE_URL>"

set -e

DB_URL="$1"

if [ -z "$DB_URL" ]; then
  if [ -n "$DATABASE_URL" ]; then
    DB_URL="$DATABASE_URL"
  else
    echo "Usage: ./scripts/apply-migrations.sh \"postgres://user:password@host:port/database\""
    echo "Or export DATABASE_URL environment variable."
    exit 1
  fi
fi

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
ROOT_DIR="$(dirname "$SCRIPT_DIR")"
MIGRATIONS_DIR="$ROOT_DIR/supabase/migrations"

echo "==> Checking psql availability..."
if ! command -v psql &> /dev/null; then
  echo "Error: psql is not installed. You can paste the contents of:"
  echo "  1) supabase/migrations/001_initial_schema.sql"
  echo "  2) supabase/migrations/002_seed_initial_data.sql"
  echo "directly into Supabase Dashboard -> SQL Editor."
  exit 1
fi

echo "==> Applying 001_initial_schema.sql..."
psql "$DB_URL" -f "$MIGRATIONS_DIR/001_initial_schema.sql"

echo "==> Applying 002_seed_initial_data.sql..."
psql "$DB_URL" -f "$MIGRATIONS_DIR/002_seed_initial_data.sql"

echo "==> Migrations applied successfully!"
