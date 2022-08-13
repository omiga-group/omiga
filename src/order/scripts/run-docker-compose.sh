#!/usr/bin/env sh

set -e
set -x

project="order"
dcpath="$project/docker-compose.yml"
command=${@:-up -d --build}

cd "$(dirname "${0}")/../.."

docker compose -p $project -f docker-compose.yml -f $dcpath $command
