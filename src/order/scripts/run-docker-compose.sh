#!/usr/bin/env sh

set -e
set -x

command=${@:-up -d --build}

cd "$(dirname "${0}")/../.."

docker compose -p "omiga" \
    --profile order \
    -f docker-compose.yml \
    -f ./order/docker-compose.yml \
    $command
