# syntax = docker/dockerfile:1.4.1
ARG FINAL_IMAGE=alpine:3
ARG BUILD_IMAGE=golang:1.19

FROM $BUILD_IMAGE as build

RUN mkdir -p /src/shared/enterprise \
    mkdir -p /src/order/shared \
    mkdir -p /src/order/orderctl

COPY go.mod /src/
COPY go.sum /src/
COPY shared/enterprise /src/shared/enterprise
COPY order/shared /src/order/shared
COPY order/orderctl /src/order/orderctl

WORKDIR /src/order/orderctl

RUN --mount=type=cache,target=/go CGO_ENABLED=0 go build -v -o /bin/orderctl main.go

FROM build AS test

RUN mkdir -p /coverage/reports
RUN --mount=type=cache,target=/go CGO_ENABLED=0 go test -coverpkg=./... -v -covermode=count -coverprofile="/coverage/reports/coverage.out" ./...

FROM $FINAL_IMAGE as final
LABEL maintainer="morteza@omiga.com.au"

WORKDIR /omiga
COPY --from=build /bin/orderctl /omiga
COPY --from=build /src/order/orderctl/config.yaml /omiga
COPY --from=build /src/order/orderctl/post_migration_script.sql /omiga

RUN /omiga/orderctl install-dependecies

# Fix CVE-2022-3996
RUN apk upgrade libssl3 libcrypto3

CMD ["/omiga/orderctl", "migrate", "--post-migration-script-path", "post_migration_script.sql"]
