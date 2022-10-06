# syntax = docker/dockerfile:1.4.1
ARG FINAL_IMAGE=alpine:3
ARG BUILD_IMAGE=golang:1.19

FROM $FINAL_IMAGE as base
LABEL maintainer="morteza@omiga.com.au"

FROM $BUILD_IMAGE as build

RUN mkdir -p /src/shared/enterprise \
    mkdir -p /src/order/shared \
    mkdir -p /src/order/shared/entities/cli

COPY go.mod /src/
COPY go.sum /src/
COPY shared/enterprise /src/shared/enterprise
COPY order/shared /src/order/shared
COPY order/shared/entities/cli /src/order/shared/entities/cli

WORKDIR /src/order/shared/entities/cli

RUN --mount=type=cache,target=/go CGO_ENABLED=0 go build -v -o /bin/order-database-migration main.go

FROM build AS local
CMD ["/bin/order-database-migration", "migrate", "--post-migration-script-path", "post-migration-script.sql"]

FROM build AS test

RUN mkdir -p /coverage/reports
RUN --mount=type=cache,target=/go CGO_ENABLED=0 go test -coverpkg=./... -v -covermode=count -coverprofile="/coverage/reports/coverage.out" ./...

FROM base AS final

WORKDIR /omiga
COPY --from=build /bin/order-database-migration /omiga
COPY --from=build /src/order/shared/entities/cli/config.yaml /omiga
COPY --from=build /src/order/shared/entities/cli/post-migration-script.sql /omiga
CMD ["/omiga/order-database-migration", "migrate", "--post-migration-script-path", "post-migration-script.sql"]
