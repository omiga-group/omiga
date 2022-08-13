#!/usr/bin/env sh

set -e
set -x

project="omiga"
dcpath="docker-compose.yml"
command=${@:-up -d --build}

cd "$(dirname "${0}")/.."

docker compose -p $project -f $dcpath \
    -f order/docker-compose.yml \
    -f exchange/docker-compose.yml \
    $command
