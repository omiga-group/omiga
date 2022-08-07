# syntax=docker/dockerfile:1.4.1
ARG FINAL_IMAGE=alpine/git:latest
ARG BUILD_IMAGE=golang:1.18

FROM $FINAL_IMAGE as base
LABEL maintainer="morteza@omiga.com.au"

FROM $BUILD_IMAGE as build

ENV CGO_ENABLED = 0

RUN mkdir -p /src/shared/enterprise \
    mkdir -p /src/shared/clients \
    mkdir -p /src/exchange/shared \
    mkdir -p /src/exchange/binance-processor

COPY go.mod /src/
COPY go.sum /src/
COPY shared/enterprise /src/shared/enterprise
COPY shared/clients /src/shared/clients
COPY exchange/shared /src/exchange/shared
COPY exchange/binance-processor /src/exchange/binance-processor

WORKDIR /src/exchange/binance-processor

RUN --mount=type=cache,target=/go go mod tidy
RUN --mount=type=cache,target=/go go build -v -o /bin/binance-processor main.go

FROM build AS local
CMD ["/bin/binance-processor", "start"]

FROM build AS test
RUN --mount=type=cache,target=/go go mod tidy
RUN --mount=type=cache,target=/go go test ./...

FROM base AS final

WORKDIR /
COPY --from=build /bin/binance-processor .
CMD ["/bin/binance-processor", "start"]
