SHELL = /bin/bash
CURRENT_DIRECTORY = $(shell pwd)

# Go variables
GOFILES = $(shell find . -type f -name '*.go' -not -path "*/mock/*.go" -not -path "*.pb.go" -not -path "*_eventgen.go" -not -path "*_gen.go")

.PHONY: all
all: dep generate ## Runs dep generate

.PHONY: dep
dep: ## Install dependencies
	@go mod tidy

.PHONY: generate
generate: ## Generate code
	@../api-definitions/asyncapi/events/scripts/generate-all.sh
	@go install github.com/google/wire/cmd/wire@latest
	@go install entgo.io/ent/cmd/ent@latest
	@go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest
	@GOFLAGS=-mod=mod go generate ./...

generate-via-docker: ## Generate code
	@../api-definitions/asyncapi/events/scripts/generate-all.sh
	@./shared/scripts/generate.sh
	@./exchange/shared/scripts/generate.sh
	@./exchange/exchange-api/scripts/generate.sh
	@./exchange/binance-processor/scripts/generate.sh
	@./exchange/coingeko-processor/scripts/generate.sh
	@./exchange/gemini-processor/scripts/generate.sh
	@./exchange/kraken-processor/scripts/generate.sh
	@./exchange/ftx-processor/scripts/generate.sh
	@./order/shared/scripts/generate.sh
	@./order/order-api/scripts/generate.sh
	@./order/order-processor/scripts/generate.sh

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
