#!/usr/bin/env sh

set -e
set -x

command=${@:-up -d --build}

cd "$(dirname "${0}")/../.."

docker compose -p "omiga" \
    --profile gateway \
    -f docker-compose.yml \
    -f ./gateway/docker-compose.yml \
    $command
