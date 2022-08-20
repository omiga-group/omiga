# syntax=docker/dockerfile:1.4.1
ARG FINAL_IMAGE=alpine:3
ARG BUILD_IMAGE=golang:1.18

FROM $FINAL_IMAGE as base
LABEL maintainer="morteza@omiga.com.au"

FROM $BUILD_IMAGE as build

RUN mkdir -p /src/shared/enterprise \
    mkdir -p /src/exchange/shared \
    mkdir -p /src/exchange/shared/repositories/cli

COPY go.mod /src/
COPY go.sum /src/
COPY shared/enterprise /src/shared/enterprise
COPY exchange/shared /src/exchange/shared
COPY exchange/shared/repositories/cli /src/exchange/shared/repositories/cli

WORKDIR /src/exchange/shared/repositories/cli

RUN --mount=type=cache,target=/go CGO_ENABLED=0 go build -v -o /bin/exchange-database-migration main.go

FROM build AS local
CMD ["/bin/exchange-database-migration", "migrate"]

FROM build AS test

RUN mkdir -p /coverage/reports
RUN --mount=type=cache,target=/go CGO_ENABLED=0 go test -v -covermode=count -coverprofile="/coverage/reports/coverage.out" ./...

FROM base AS final

WORKDIR /omiga
COPY --from=test /bin/exchange-database-migration /omiga
COPY --from=test /src/exchange/shared/repositories/cli/config.yaml /omiga
CMD ["/omiga/exchange-database-migration", "migrate"]
