#!/usr/bin/env sh

set -e
set -x

cd "$(dirname "${0}")/../../api-definitions/graphql/omiga"

if ! [ -x "$(command -v rover)" ]; then
  curl -sSL https://rover.apollo.dev/nix/latest | sh
fi

rover graph introspect http://localhost:4343 > schema.graphql
