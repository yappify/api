#!/bin/sh

# Load environment variables from .env file
set -a
. ./.env
set +a

# Run the goose command with the provided argument and migration directory
goose -dir ./sql/migrations postgres "${DB_URL}" "$1"