#!/usr/bin/env sh

set -e
set -x

cd "$(dirname "${0}")/.."

cleanup() {
    docker rm generated-image || true
}
trap cleanup EXIT

docker build --file Dockerfile.generate --tag generator:latest ../../
docker create --name generated-image generator:latest
docker cp generated-image:/src/venue/venue-api/. ./
