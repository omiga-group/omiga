# syntax = docker/dockerfile:1.4.1
ARG FINAL_IMAGE=alpine:3
ARG BUILD_IMAGE=golang:1.19

FROM $FINAL_IMAGE as base
LABEL maintainer="morteza@omiga.com.au"

FROM $BUILD_IMAGE as build

RUN mkdir -p /src/shared/enterprise \
    mkdir -p /src/shared/clients \
    mkdir -p /src/venue/shared \
    mkdir -p /src/venue/kucoin-processor

COPY go.mod /src/
COPY go.sum /src/
COPY shared/enterprise /src/shared/enterprise
COPY shared/clients /src/shared/clients
COPY venue/shared /src/venue/shared
COPY venue/kucoin-processor /src/venue/kucoin-processor

WORKDIR /src/venue/kucoin-processor

RUN --mount=type=cache,target=/go CGO_ENABLED=0 go build -v -o /bin/kucoin-processor main.go

FROM build AS local
CMD ["/bin/kucoin-processor", "start"]

FROM build AS test

RUN mkdir -p /coverage/reports
RUN --mount=type=cache,target=/go CGO_ENABLED=0 go test -coverpkg=./... -v -covermode=count -coverprofile="/coverage/reports/coverage.out" ./...

FROM base AS final

WORKDIR /omiga
COPY --from=build /bin/kucoin-processor /omiga
COPY --from=build /src/venue/kucoin-processor/config.yaml /omiga
CMD ["/omiga/kucoin-processor", "start"]