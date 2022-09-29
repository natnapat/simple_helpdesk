#!/bin/sh

set -e

echo "run db migration"
/app/migrate -path /app/migration -database "postgresql://root:mysecret@postgres:5432/simple_helpdesk?sslmode=disable" -verbose up

echo "start the app"
exec "$@"