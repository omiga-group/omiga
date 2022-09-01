FROM node:16-alpine
LABEL maintainer="morteza@omiga.com.au"

RUN npm install -g @asyncapi/generator

COPY src/utils/go-omiga-websocket-template /go-omiga-websocket-template
WORKDIR /go-omiga-websocket-template
RUN --mount=type=cache,target=~/.yarn YARN_CACHE_FOLDER=~/.yarn yarn install --frozen-lockfile

COPY api-definitions/asyncapi/websockets /asyncapi

RUN ag /asyncapi/kraken/Kraken_V1_8_0.yaml @asyncapi/nodejs-ws-template --force-write -p server=swamp --output /generated-src/kraken1/kraken/v1_8_0

# RUN ag /asyncapi/kraken/Kraken_V1_8_0.yaml /go-omiga-websocket-template \
#     --output /generated-src/kraken/kraken/v1_8_0 \
#     --force-write \
#     -p packageName=krakenv180
