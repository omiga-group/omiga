# syntax=docker/dockerfile:1.4.1
ARG BUILD_IMAGE=golang:1.19

FROM $BUILD_IMAGE as test
LABEL maintainer="morteza@omiga.com.au"

RUN mkdir -p /src/shared/enterprise \
    mkdir -p /src/shared/clients

COPY go.mod /src/
COPY go.sum /src/
COPY shared/enterprise /src/shared/enterprise
COPY shared/clients /src/shared/clients

WORKDIR /src/shared

RUN mkdir -p /coverage/reports
RUN --mount=type=cache,target=/go CGO_ENABLED=0 go test -coverpkg=./... -v -covermode=count -coverprofile="/coverage/reports/coverage.out" ./...
