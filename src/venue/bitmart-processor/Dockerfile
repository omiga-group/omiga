# syntax = docker/dockerfile:1.4.1
ARG FINAL_IMAGE=alpine:3
ARG BUILD_IMAGE=golang:1.20

FROM $BUILD_IMAGE as build

RUN mkdir -p /src/shared/enterprise \
    mkdir -p /src/shared/clients \
    mkdir -p /src/venue/shared \
    mkdir -p /src/venue/bitmart-processor

COPY go.mod /src/
COPY go.sum /src/
COPY shared/enterprise /src/shared/enterprise
COPY shared/clients /src/shared/clients
COPY venue/shared /src/venue/shared
COPY venue/bitmart-processor /src/venue/bitmart-processor

WORKDIR /src/venue/bitmart-processor

RUN --mount=type=cache,target=/go CGO_ENABLED=0 go build -v -o /bin/bitmart-processor main.go

FROM build AS test

RUN mkdir -p /coverage/reports
RUN --mount=type=cache,target=/go CGO_ENABLED=0 go test -coverpkg=./... -v -covermode=count -coverprofile="/coverage/reports/coverage.out" ./...

FROM $FINAL_IMAGE as final
LABEL maintainer="morteza@omiga.com.au"

WORKDIR /omiga
COPY --from=build /bin/bitmart-processor /omiga
COPY --from=build /src/venue/bitmart-processor/config.yaml /omiga

RUN /omiga/bitmart-processor install-dependecies

# Fix CVE-2022-3996
RUN apk upgrade libssl3 libcrypto3

CMD ["/omiga/bitmart-processor", "start"]
