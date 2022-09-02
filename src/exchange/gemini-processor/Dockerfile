# syntax = docker/dockerfile:1.4.1
ARG FINAL_IMAGE=alpine:3
ARG BUILD_IMAGE=golang:1.19

FROM $FINAL_IMAGE as base
LABEL maintainer="abdollah.alizadeh.dahanesari@omiga.com.au"

FROM $BUILD_IMAGE as build

RUN mkdir -p /src/shared/enterprise \
    mkdir -p /src/shared/clients \
    mkdir -p /src/exchange/shared \
    mkdir -p /src/exchange/gemini-processor

COPY go.mod /src/
COPY go.sum /src/
COPY shared/enterprise /src/shared/enterprise
COPY shared/clients /src/shared/clients
COPY exchange/shared /src/exchange/shared
COPY exchange/gemini-processor /src/exchange/gemini-processor

WORKDIR /src/exchange/gemini-processor

RUN --mount=type=cache,target=/go CGO_ENABLED=0 go build -v -o /bin/gemini-processor main.go

FROM build AS local
CMD ["/bin/gemini-processor", "start"]

FROM build AS test

RUN mkdir -p /coverage/reports
RUN --mount=type=cache,target=/go CGO_ENABLED=0 go test -coverpkg=./... -v -covermode=count -coverprofile="/coverage/reports/coverage.out" ./...

FROM base AS final

WORKDIR /omiga
COPY --from=build /bin/gemini-processor /omiga
COPY --from=build /src/exchange/gemini-processor/config.yaml /omiga
CMD ["/omiga/gemini-processor", "start"]
