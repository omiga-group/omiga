# syntax = docker/dockerfile:1.4.1
ARG FINAL_IMAGE=alpine:3
ARG BUILD_IMAGE=golang:1.20

FROM $BUILD_IMAGE as build

RUN mkdir -p /src/shared/enterprise \
    mkdir -p /src/venue/shared \
    mkdir -p /src/venue/venuectl

COPY go.mod /src/
COPY go.sum /src/
COPY shared/enterprise /src/shared/enterprise
COPY venue/shared /src/venue/shared
COPY venue/venuectl /src/venue/venuectl

WORKDIR /src/venue/venuectl

RUN --mount=type=cache,target=/go CGO_ENABLED=0 go build -v -o /bin/venuectl main.go

FROM build AS test

RUN mkdir -p /coverage/reports
RUN --mount=type=cache,target=/go CGO_ENABLED=0 go test -coverpkg=./... -v -covermode=count -coverprofile="/coverage/reports/coverage.out" ./...

FROM $FINAL_IMAGE as final
LABEL maintainer="morteza@omiga.com.au"

WORKDIR /omiga
COPY --from=build /bin/venuectl /omiga
COPY --from=build /src/venue/venuectl/config.yaml /omiga

RUN /omiga/venuectl install-dependecies

# Fix CVE-2022-3996
RUN apk upgrade libssl3 libcrypto3
