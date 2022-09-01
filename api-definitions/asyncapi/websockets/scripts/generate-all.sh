#!/usr/bin/env sh

set -e
set -x

cd "$(dirname "${0}")/.."

cleanup() {
   docker rm extract-asyncapi-websockets-generator || true
}
trap cleanup EXIT

# Delete previously generated golang files
find "../../../src/shared/websockets" -name "*_gen.go" -type f -delete || true

# AsyncApi
docker build --progress=plain -f Dockerfile -t asyncapi-websockets-generator ../../../
docker create --name extract-asyncapi-websockets-generator asyncapi-websockets-generator
mkdir -p "../../../src/shared/clients/websockets"
docker cp extract-asyncapi-websockets-generator:/generated-src/. "../../../src/shared/clients/websockets/"
