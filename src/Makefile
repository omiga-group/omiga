SHELL = /bin/bash
CURRENT_DIRECTORY = $(shell pwd)

# Go variables
GOFILES = $(shell find . -type f -name '*.go' -not -path "*/mock/*.go" -not -path "*.pb.go")

.PHONY: all
all: dep generate ## Runs dep generate

.PHONY: dep
dep: ## Install dependencies
	@go get -d golang.org/x/tools/cmd/cover
	@go get -d github.com/mattn/goveralls
	@go mod tidy

.PHONY: generate
generate: ## Generate models from GraphQL Schema
	@../api-definitions/asyncapi/scripts/generate-all.sh
	@go install github.com/google/wire/cmd/wire@latest
	@go install entgo.io/ent/cmd/ent@latest
	@GOFLAGS=-mod=mod go generate ./...

.PHONY: lint
lint: ## run golanci-lint locally
	@docker run --rm -v $(CURRENT_DIRECTORY):/app -w /app golangci/golangci-lint:latest golangci-lint run -v

.PHONY: format
format: ## Format the source
	@goimports -w $(GOFILES)

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
