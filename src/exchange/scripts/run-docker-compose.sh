#!/usr/bin/env sh

set -e
set -x

command=${@:-up -d --build}

cd "$(dirname "${0}")/../.."

docker compose -p "omiga" \
    --profile exchange \
    -f docker-compose.yml \
    -f ./exchange/docker-compose.yml \
    $command
