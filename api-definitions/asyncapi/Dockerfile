FROM node:latest
LABEL maintainer="morteza@omiga.com.au"

RUN npm install -g @asyncapi/generator

COPY src/utils/go-omiga-template /go-omiga-template
WORKDIR /go-omiga-template
RUN --mount=type=cache,target=~/.yarn YARN_CACHE_FOLDER=~/.yarn yarn install --frozen-lockfile

COPY api-definitions/asyncapi /asyncapi
RUN ag /asyncapi/omiga/Order_V1.yaml /go-omiga-template \
    --output /generated-src/events/omiga/order/v1 \
    --force-write \
    -p packageName=orderv1
