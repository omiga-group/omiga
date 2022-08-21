FROM node:16-alpine
LABEL maintainer="morteza@omiga.com.au"

RUN mkdir -p /src/gateway

COPY gateway /src/gateway

WORKDIR /src/gateway
RUN --mount=type=cache,target=~/.yarn YARN_CACHE_FOLDER=~/.yarn yarn install --frozen-lockfile

CMD ["node", "index.js"]
