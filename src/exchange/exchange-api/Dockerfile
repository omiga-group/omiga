# syntax=docker/dockerfile:1.4.1
ARG FINAL_IMAGE=alpine:3
ARG BUILD_IMAGE=golang:1.19

FROM $FINAL_IMAGE as base
LABEL maintainer="morteza@omiga.com.au"

FROM $BUILD_IMAGE as build

RUN mkdir -p /src/shared/enterprise \
    mkdir -p /src/shared/clients \
    mkdir -p /src/exchange/shared \
    mkdir -p /src/exchange/exchange-api

COPY go.mod /src/
COPY go.sum /src/
COPY shared/enterprise /src/shared/enterprise
COPY shared/clients /src/shared/clients
COPY exchange/shared /src/exchange/shared
COPY exchange/exchange-api /src/exchange/exchange-api

WORKDIR /src/exchange/exchange-api

RUN --mount=type=cache,target=/go CGO_ENABLED=0 go build -v -o /bin/exchange-api main.go

FROM build AS local
CMD ["/bin/exchange-api", "start"]

FROM build AS test

RUN mkdir -p /coverage/reports
RUN --mount=type=cache,target=/go CGO_ENABLED=0 go test -v -covermode=count -coverprofile="/coverage/reports/coverage.out" ./...

FROM base AS final

WORKDIR /omiga
COPY --from=test /bin/exchange-api /omiga
COPY --from=test /src/exchange/exchange-api/config.yaml /omiga
CMD ["/omiga/exchange-api", "start"]
