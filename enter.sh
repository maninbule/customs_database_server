#!/bin/bash
CGO_ENABLED=0 GOOS=linux go build -o /customs_database_server /app
/customs_database_server
exec "$@"