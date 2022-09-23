#!/usr/bin/env sh

set -e
set -x

command=${@:-up -d --build}

cd "$(dirname "${0}")/.."

docker compose -p "omiga" \
    --profile core \
    -f docker-compose.yml \
    -f shared/docker-compose.yml \
    -f order/docker-compose.yml \
    -f exchange/docker-compose.yml \
    -f gateway/docker-compose.yml \
    -f web/docker-compose.yml \
    $command
