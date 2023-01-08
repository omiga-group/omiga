FROM node:16-alpine as build
LABEL maintainer="morteza@omiga.com.au"

RUN mkdir -p /src/web

COPY src/web /src/web
COPY api-definitions/graphql/omiga/schema.graphql /src/web

WORKDIR /src/web

RUN --mount=type=cache,target=~/.yarn YARN_CACHE_FOLDER=~/.yarn yarn install --frozen-lockfile
RUN sed -i 's/..\/..\/api-definitions\/graphql\/omiga\/schema.graphql/schema.graphql/g' relay.config.js
RUN yarn run relay
RUN yarn run build

FROM nginx:1-alpine

# Nginx config
RUN rm -rf /etc/nginx/conf.d
COPY --from=build /src/web/conf /etc/nginx/

# Static build
COPY --from=build /src/web/build /usr/share/nginx/html

# Initialize environment variables into filesystem
WORKDIR /usr/share/nginx/html
COPY --from=build /src/web/scripts/env.sh .
COPY --from=build /src/web/.env .

# Add bash
RUN apk add --no-cache bash

# Run script which initializes env vars to fs
RUN chmod +x env.sh

# Fix CVE-2022-3996, CVE-2022-43551
RUN apk upgrade libssl3 libcrypto3 libcurl curl

# Start Nginx server
CMD ["/bin/bash", "-c", "/usr/share/nginx/html/env.sh && nginx -g \"daemon off;\""]
