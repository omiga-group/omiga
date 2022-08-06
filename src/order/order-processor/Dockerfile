# syntax=docker/dockerfile:1.4.1
ARG FINAL_IMAGE=alpine/git:latest
ARG BUILD_IMAGE=golang:1.18

FROM $FINAL_IMAGE as base
LABEL maintainer="morteza@omiga.com.au"

FROM $BUILD_IMAGE as build

ENV CGO_ENABLED = 0

RUN mkdir -p /src/shared/enterprise \
    mkdir -p /src/shared/clients \
    mkdir -p /src/order/shared \
    mkdir -p /src/order/order-processor

COPY go.mod /src/
COPY go.sum /src/
COPY shared/enterprise /src/shared/enterprise
COPY shared/clients /src/shared/clients
COPY order/shared /src/order/shared
COPY order/order-processor /src/order/order-processor

WORKDIR /src/order/order-processor

RUN --mount=type=cache,target=/go go mod tidy
RUN --mount=type=cache,target=/go go get -v -t ./...
RUN --mount=type=cache,target=/go go build -v -o /bin/order-processor main.go

FROM build AS local
CMD ["/bin/order-processor", "start"]

FROM build AS test
RUN --mount=type=cache,target=/go go mod tidy
RUN --mount=type=cache,target=/go go test ./...

FROM base AS final

WORKDIR /
COPY --from=build /bin/order-processor .
CMD ["/bin/order-processor", "start"]
