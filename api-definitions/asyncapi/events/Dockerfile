FROM node:16-alpine
LABEL maintainer="morteza@omiga.com.au"

RUN npm install -g @asyncapi/generator

COPY src/utils/go-omiga-event-template /go-omiga-event-template
WORKDIR /go-omiga-event-template
RUN --mount=type=cache,target=~/.yarn YARN_CACHE_FOLDER=~/.yarn yarn install --frozen-lockfile

COPY src/utils/go-omiga-event-jsonschema-template /go-omiga-event-jsonschema-template
WORKDIR /go-omiga-event-jsonschema-template
RUN --mount=type=cache,target=~/.yarn YARN_CACHE_FOLDER=~/.yarn yarn install --frozen-lockfile

COPY api-definitions/asyncapi/events /asyncapi

RUN ag /asyncapi/omiga/Order_V1.yaml /go-omiga-event-template \
    --output /generated-src/omiga/order/v1 \
    --force-write \
    -p packageName=orderv1

RUN ag /asyncapi/omiga/SyntheticOrder_V1.yaml /go-omiga-event-template \
    --output /generated-src/omiga/synthetic-order/v1 \
    --force-write \
    -p packageName=syntheticorderv1

RUN ag /asyncapi/omiga/OrderBook_V1.yaml /go-omiga-event-template \
    --output /generated-src/omiga/order-book/v1 \
    --force-write \
    -p packageName=orderbookv1

# Generate metadata including topic name and json schema

RUN ag /asyncapi/omiga/Order_V1.yaml /go-omiga-event-jsonschema-template \
    --output /generated-src/omiga/order/v1/schema \
    --force-write \
    -p packageName=orderv1

RUN ag /asyncapi/omiga/SyntheticOrder_V1.yaml /go-omiga-event-jsonschema-template \
    --output /generated-src/omiga/synthetic-order/v1/schema \
    --force-write \
    -p packageName=syntheticorderv1

RUN ag /asyncapi/omiga/OrderBook_V1.yaml /go-omiga-event-jsonschema-template \
    --output /generated-src/omiga/order-book/v1/schema \
    --force-write \
    -p packageName=orderbookv1
