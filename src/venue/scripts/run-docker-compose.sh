#!/usr/bin/env sh

set -e
set -x

command=${@:-up -d --build}

cd "$(dirname "${0}")/../.."

docker compose -p "omiga" \
    --profile venue \
    -f docker-compose.yml \
    -f ./venue/docker-compose.yml \
    $command
