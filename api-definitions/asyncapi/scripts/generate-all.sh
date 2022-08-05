#!/usr/bin/env sh

set -e
set -x

cd "$(dirname "${0}")/.."

cleanup() {
   docker rm extract-asyncapi-generator || true
}
trap cleanup EXIT

# Delete previously generated golang files
find ../../code/shared/events -name "*_gen.go" -type f -delete || true

# AsyncApi
docker build --progress=plain -f Dockerfile -t asyncapi-generator ../../
docker create --name extract-asyncapi-generator asyncapi-generator
docker cp extract-asyncapi-generator:/generated-code/. "../../code/shared/clients"
