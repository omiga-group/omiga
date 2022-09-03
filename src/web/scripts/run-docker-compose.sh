#!/usr/bin/env sh

set -e
set -x

command=${@:-up -d --build}

cd "$(dirname "${0}")/../.."

docker compose -p "omiga" \
    --profile web \
    -f docker-compose.yml \
    -f ./web/docker-compose.yml \
    $command
