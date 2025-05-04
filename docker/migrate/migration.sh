#!/usr/bin/env sh

migrate \
  -path "/migrations" \
  -database "postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable" \
  "$@"
