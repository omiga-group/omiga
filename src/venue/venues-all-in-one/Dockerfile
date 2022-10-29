# syntax = docker/dockerfile:1.4.1
ARG FINAL_IMAGE=node:18
ARG BUILD_IMAGE=golang:1.19

FROM $FINAL_IMAGE as base
LABEL maintainer="morteza@omiga.com.au"

FROM $BUILD_IMAGE as build

RUN mkdir -p /src/shared/enterprise \
    mkdir -p /src/shared/clients \
    mkdir -p /src/venue/shared \
    mkdir -p /src/venue/venues-all-in-one \
    mkdir -p /src/venue/binance-processor \
    mkdir -p /src/venue/bitmart-processor \
    mkdir -p /src/venue/bittrex-processor \
    mkdir -p /src/venue/bybit-processor \
    mkdir -p /src/venue/coinbase-processor \
    mkdir -p /src/venue/crypto-processor \
    mkdir -p /src/venue/dextrade-processor \
    mkdir -p /src/venue/ftx-processor \
    mkdir -p /src/venue/gemini-processor \
    mkdir -p /src/venue/huobi-processor \
    mkdir -p /src/venue/kraken-processor \
    mkdir -p /src/venue/kucoin-processor \
    mkdir -p /src/venue/mexc-processor \
    mkdir -p /src/venue/rain-processor \
    mkdir -p /src/venue/xt-processor

COPY go.mod /src/
COPY go.sum /src/
COPY shared/enterprise /src/shared/enterprise
COPY shared/clients /src/shared/clients
COPY venue/shared /src/venue/shared
COPY venue/venues-all-in-one /src/venue/venues-all-in-one
COPY venue/binance-processor /src/venue/binance-processor
COPY venue/bitmart-processor /src/venue/bitmart-processor
COPY venue/bittrex-processor /src/venue/bittrex-processor
COPY venue/bybit-processor /src/venue/bybit-processor
COPY venue/coinbase-processor /src/venue/coinbase-processor
COPY venue/crypto-processor /src/venue/crypto-processor
COPY venue/dextrade-processor /src/venue/dextrade-processor
COPY venue/ftx-processor /src/venue/ftx-processor
COPY venue/gemini-processor /src/venue/gemini-processor
COPY venue/huobi-processor /src/venue/huobi-processor
COPY venue/kraken-processor /src/venue/kraken-processor
COPY venue/kucoin-processor /src/venue/kucoin-processor
COPY venue/mexc-processor /src/venue/mexc-processor
COPY venue/rain-processor /src/venue/rain-processor
COPY venue/xt-processor /src/venue/xt-processor

WORKDIR /src/venue/venues-all-in-one

RUN --mount=type=cache,target=/go CGO_ENABLED=0 go build -v -o /bin/venues-all-in-one main.go

FROM build AS test

RUN mkdir -p /coverage/reports
RUN --mount=type=cache,target=/go CGO_ENABLED=0 go test -coverpkg=./... -v -covermode=count -coverprofile="/coverage/reports/coverage.out" ./...

FROM base AS final

RUN npx playwright install-deps

WORKDIR /omiga

COPY --from=build /bin/venues-all-in-one /omiga
COPY --from=build /src/venue/binance-processor/config.yaml /omiga/binance-processor-config.yaml
COPY --from=build /src/venue/bitmart-processor/config.yaml /omiga/bitmart-processor-config.yaml
COPY --from=build /src/venue/bittrex-processor/config.yaml /omiga/bittrex-processor-config.yaml
COPY --from=build /src/venue/bybit-processor/config.yaml /omiga/bybit-processor-config.yaml
COPY --from=build /src/venue/coinbase-processor/config.yaml /omiga/coinbase-processor-config.yaml
COPY --from=build /src/venue/crypto-processor/config.yaml /omiga/crypto-processor-config.yaml
COPY --from=build /src/venue/dextrade-processor/config.yaml /omiga/dextrade-processor-config.yaml
COPY --from=build /src/venue/ftx-processor/config.yaml /omiga/ftx-processor-config.yaml
COPY --from=build /src/venue/gemini-processor/config.yaml /omiga/gemini-processor-config.yaml
COPY --from=build /src/venue/huobi-processor/config.yaml /omiga/huobi-processor-config.yaml
COPY --from=build /src/venue/kraken-processor/config.yaml /omiga/kraken-processor-config.yaml
COPY --from=build /src/venue/kucoin-processor/config.yaml /omiga/kucoin-processor-config.yaml
COPY --from=build /src/venue/mexc-processor/config.yaml /omiga/mexc-processor-config.yaml
COPY --from=build /src/venue/rain-processor/config.yaml /omiga/rain-processor-config.yaml
COPY --from=build /src/venue/xt-processor/config.yaml /omiga/xt-processor-config.yaml

RUN /omiga/venues-all-in-one install-dependecies

CMD ["/omiga/venues-all-in-one", "start"]
