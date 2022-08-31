# syntax = docker/dockerfile:1.4.1
ARG FINAL_IMAGE=alpine:3
ARG BUILD_IMAGE=golang:1.19

FROM $FINAL_IMAGE as base
LABEL maintainer="morteza@omiga.com.au"

FROM $BUILD_IMAGE as build

RUN mkdir -p /src/shared/enterprise \
    mkdir -p /src/shared/clients \
    mkdir -p /src/order/shared \
    mkdir -p /src/order/order-api

COPY go.mod /src/
COPY go.sum /src/
COPY shared/enterprise /src/shared/enterprise
COPY shared/clients /src/shared/clients
COPY order/shared /src/order/shared
COPY order/order-api /src/order/order-api

WORKDIR /src/order/order-api

RUN --mount=type=cache,target=/go CGO_ENABLED=0 go build -v -o /bin/order-api main.go

FROM build AS local
CMD ["/bin/order-api", "start"]

FROM build AS test

RUN mkdir -p /coverage/reports
RUN --mount=type=cache,target=/go CGO_ENABLED=0 go test -coverpkg=./... -v -covermode=count -coverprofile="/coverage/reports/coverage.out" ./...

FROM base AS final

WORKDIR /omiga
COPY --from=build /bin/order-api /omiga
COPY --from=build /src/order/order-api/config.yaml /omiga
CMD ["/omiga/order-api", "start"]
