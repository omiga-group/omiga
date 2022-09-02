# syntax = docker/dockerfile:1.4.1
ARG FINAL_IMAGE=alpine:3
ARG BUILD_IMAGE=golang:1.19

FROM $FINAL_IMAGE as base
LABEL maintainer="morteza@omiga.com.au"

FROM $BUILD_IMAGE as build

RUN mkdir -p /src/shared/enterprise \
    mkdir -p /src/shared/clients \
    mkdir -p /src/exchange/shared \
    mkdir -p /src/exchange/kraken-processor

COPY go.mod /src/
COPY go.sum /src/
COPY shared/enterprise /src/shared/enterprise
COPY shared/clients /src/shared/clients
COPY exchange/shared /src/exchange/shared
COPY exchange/kraken-processor /src/exchange/kraken-processor

WORKDIR /src/exchange/kraken-processor

RUN --mount=type=cache,target=/go CGO_ENABLED=0 go build -v -o /bin/kraken-processor main.go

FROM build AS local
CMD ["/bin/kraken-processor", "start"]

FROM build AS test

RUN mkdir -p /coverage/reports
RUN --mount=type=cache,target=/go CGO_ENABLED=0 go test -coverpkg=./... -v -covermode=count -coverprofile="/coverage/reports/coverage.out" ./...

FROM base AS final

WORKDIR /omiga
COPY --from=build /bin/kraken-processor /omiga
COPY --from=build /src/exchange/kraken-processor/config.yaml /omiga
CMD ["/omiga/kraken-processor", "start"]
