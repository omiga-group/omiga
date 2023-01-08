# syntax = docker/dockerfile:1.4.1
FROM microbusiness/atlas:latest
LABEL maintainer="morteza@omiga.com.au"

WORKDIR /migrations
COPY order/shared/migrations /migrations

CMD /atlas migrate apply --dir file:///migrations/ --url $OMIGA_POSTGRES_CONNECTIONSTRING
