# syntax=docker/dockerfile:1.4.1
ARG FINAL_IMAGE=alpine/git:latest
ARG BUILD_IMAGE=golang:1.18

FROM $FINAL_IMAGE as base
LABEL maintainer="morteza@omiga.com.au"

FROM $BUILD_IMAGE as build

ENV CGO_ENABLED = 0

RUN mkdir -p /src/exchange/ftx-processor

COPY go.mod /src/
COPY go.sum /src/
COPY exchange/ftx-processor /src/exchange/ftx-processor

WORKDIR /src/exchange/ftx-processor

RUN --mount=type=cache,target=/go go mod tidy
RUN --mount=type=cache,target=/go go get -v -t ./...
RUN --mount=type=cache,target=/go go build -v -o /bin/ftx-processor main.go

FROM build AS local
CMD ["/bin/ftx-processor", "start"]

FROM build AS test
RUN --mount=type=cache,target=/go go mod tidy
RUN --mount=type=cache,target=/go go test ./...

FROM base AS final

WORKDIR /
COPY --from=build /bin/ftx-processor .
CMD ["/bin/ftx-processor", "start"]
