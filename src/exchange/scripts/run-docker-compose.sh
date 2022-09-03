#!/usr/bin/env sh

set -e
set -x

command=${@:-up -d --build}

cd "$(dirname "${0}")/../.."

docker compose -p "exchange" \
    --profile exchange \
    -f docker-compose.yml \
    -f ./exchange/docker-compose.yml \
    $command
