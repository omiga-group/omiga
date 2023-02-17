# syntax = docker/dockerfile:1.4.1
ARG BUILD_IMAGE=node:18-alpine

FROM $BUILD_IMAGE as build
LABEL maintainer="morteza@omiga.com.au"

RUN mkdir -p /src/gateway

COPY gateway /src/gateway

WORKDIR /src/gateway
RUN --mount=type=cache,target=~/.yarn YARN_CACHE_FOLDER=~/.yarn yarn install --frozen-lockfile

# Fix CVE-2022-3996
RUN apk upgrade libssl3 libcrypto3

CMD ["node", "index.js"]
