# syntax=docker/dockerfile:1.4.1
ARG FINAL_IMAGE=alpine/git:latest
ARG BUILD_IMAGE=golang:1.18

FROM $FINAL_IMAGE as base
LABEL maintainer="morteza@omiga.com.au"

FROM $BUILD_IMAGE as build

ENV CGO_ENABLED = 0

RUN mkdir -p /src/shared/enterprise \
    mkdir -p /src/order/shared \
    mkdir -p /src/order/shared/repositories/cli

COPY go.mod /src/
COPY go.sum /src/
COPY shared/enterprise /src/shared/enterprise
COPY order/shared /src/order/shared
COPY order/shared/repositories/cli /src/order/shared/repositories/cli

WORKDIR /src/order/shared/repositories/cli

RUN --mount=type=cache,target=/go go build -v -o /bin/order-database-migration main.go

FROM build AS local
CMD ["/bin/order-database-migration", "migrate"]

FROM build AS test
RUN --mount=type=cache,target=/go go test ./...

FROM base AS final

WORKDIR /
COPY --from=build /bin/order-database-migration .
CMD ["/bin/order-database-migration", "migrate"]
