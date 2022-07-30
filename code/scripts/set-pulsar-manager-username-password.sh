#!/usr/bin/env sh

set -e
set -x

CSRF_TOKEN=$(curl http://dashboard.localhost:7750/pulsar-manager/csrf-token)
curl \
    -H "X-XSRF-TOKEN: $CSRF_TOKEN" \
    -H "Cookie: XSRF-TOKEN=$CSRF_TOKEN;" \
    -H 'Content-Type: application/json' \
    -X PUT http://dashboard.localhost:7750/pulsar-manager/users/superuser \
    -d '{"name": "admin", "password": "123456", "description": "admin", "email": "morteza@omiga.com.au"}'
