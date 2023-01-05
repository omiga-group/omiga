SHELL = /bin/bash
CURRENT_DIRECTORY = $(shell pwd)

# Go variables
GOFILES = $(shell find . -type f -name '*.go' -not -path "*/mock/*.go" -not -path "*.pb.go" -not -path "*_eventgen.go" -not -path "*_gen.go")

.PHONY: all
all: dep generate ## Runs dep generate

.PHONY: dep
dep: ## Install dependencies
	@go mod tidy

.PHONY: gen-dep
gen-dep: ## Install code generation dependencies
	@go install github.com/google/wire/cmd/wire@latest
	@go install entgo.io/ent/cmd/ent@latest
	@go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest
	@go install github.com/golang/mock/mockgen@latest

.PHONY: generate
generate: ## Generate code
	@../api-definitions/asyncapi/events/scripts/generate-all.sh
	@GOFLAGS=-mod=mod go generate ./...

.PHONY: lint
lint: ## run golanci-lint locally
	@docker pull golangci/golangci-lint:latest-alpine
	@docker run --rm -v $(CURRENT_DIRECTORY):/app -w /app golangci/golangci-lint:latest-alpine golangci-lint run -v

.PHONY: format
format: ## Format the source
	@goimports -w $(GOFILES)

.PHONY: stop-venues-dev
stop-venues-dev:
	@docker compose -f docker-compose-venues.yml -p omiga-venues-dev --env-file .env.omiga.venues.dev down -v || true

.PHONY: restart-venues-dev
restart-venues-dev:
	@docker compose -f docker-compose-venues.yml -p omiga-venues-dev --env-file .env.omiga.venues.dev down -v || true
	@docker compose -f docker-compose-venues.yml -p omiga-venues-dev --env-file .env.omiga.venues.dev pull
	@docker compose -f docker-compose-venues.yml -p omiga-venues-dev --env-file .env.omiga.venues.dev up -d

.PHONY: stop-venues
stop-venues:
	@docker compose -f docker-compose-venues.yml -p omiga-venues --env-file .env.omiga.venues down -v || true

.PHONY: restart-venues
restart-venues:
	@docker compose -f docker-compose-venues.yml -p omiga-venues --env-file .env.omiga.venues down -v || true
	@docker compose -f docker-compose-venues.yml -p omiga-venues --env-file .env.omiga.venues pull
	@docker compose -f docker-compose-venues.yml -p omiga-venues --env-file .env.omiga.venues up -d

.PHONY: stop-apis-dev
stop-apis-dev:
	@docker compose -f docker-compose-apis.yml -p omiga-apis-dev --env-file .env.omiga.apis.dev down -v || true

.PHONY: restart-apis-dev
restart-apis-dev:
	@docker compose -f docker-compose-apis.yml -p omiga-apis-dev --env-file .env.omiga.apis.dev down -v || true
	@docker compose -f docker-compose-apis.yml -p omiga-apis-dev --env-file .env.omiga.apis.dev pull
	@docker compose -f docker-compose-apis.yml -p omiga-apis-dev --env-file .env.omiga.apis.dev up -d

.PHONY: stop-apis
stop-apis:
	@docker compose -f docker-compose-apis.yml -p omiga-apis --env-file .env.omiga.apis down -v || true

.PHONY: restart-apis
restart-apis:
	@docker compose -f docker-compose-apis.yml -p omiga-apis --env-file .env.omiga.apis down -v || true
	@docker compose -f docker-compose-apis.yml -p omiga-apis --env-file .env.omiga.apis pull
	@docker compose -f docker-compose-apis.yml -p omiga-apis --env-file .env.omiga.apis up -d

.PHONY: stop-web-dev
stop-web-dev:
	@docker compose -f docker-compose-web.yml -p omiga-web-dev --env-file .env.omiga.web.dev down -v || true

.PHONY: restart-web-dev
restart-web-dev:
	@docker compose -f docker-compose-web.yml -p omiga-web-dev --env-file .env.omiga.web.dev down -v || true
	@docker compose -f docker-compose-web.yml -p omiga-web-dev --env-file .env.omiga.web.dev pull
	@docker compose -f docker-compose-web.yml -p omiga-web-dev --env-file .env.omiga.web.dev up -d

.PHONY: stop-web
stop-web:
	@docker compose -f docker-compose-web.yml -p omiga-web --env-file .env.omiga.web down -v || true

.PHONY: restart-web
restart-web:
	@docker compose -f docker-compose-web.yml -p omiga-web --env-file .env.omiga.web down -v || true
	@docker compose -f docker-compose-web.yml -p omiga-web --env-file .env.omiga.web pull
	@docker compose -f docker-compose-web.yml -p omiga-web --env-file .env.omiga.web up -d

.PHONY: list
list: ## List all make targets
	@$(MAKE) -pRrn : -f $(MAKEFILE_LIST) 2>/dev/null | awk -v RS= -F: '/^# File/,/^# Finished Make data base/ {if ($$1 !~ "^[#.]") {print $$1}}' | egrep -v -e '^[^[:alnum:]]' -e '^$@$$' | sort

.PHONY: help
.DEFAULT_GOAL := help
help: ## Get help output
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

# Variable outputting/exporting rules
var-%: ; @echo $($*)
varexport-%: ; @echo $*=$($*)
