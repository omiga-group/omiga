#!/usr/bin/env sh

set -e
set -x

project="exchange"
dcpath="$project/docker-compose.yml"
command=${@:-up -d --build}

cd "$(dirname "${0}")/../.."

docker compose -p $project -f docker-compose.yml -f $dcpath $command
