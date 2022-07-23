# Build variables
POC_API_BINARY_NAME = poc
BACKEND_DIR = backend/
BUILD_DIR ?= bin


.PHONY: all
all: dep generate build-poc #install ## Runs dep generate build-api build-simulator

.PHONY: clean
clean: ## Clean the working area and the project
	@rm -rf $(BUILD_DIR)/

.PHONY: dep
dep: ## Install dependencies
	@cd $(BACKEND_DIR) && \
	go install github.com/golang/mock/mockgen@v1.6.0 && \
	go get -u -d github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest && \
	go mod tidy && \

.PHONY: generate
generate: ## Generates (golang) codes from different sources
	@cd $(BACKEND_DIR) && \
	go generate ./...

.PHONY: build-poc
build-api: GOARGS += -o $(BUILD_DIR)/$(POC_API_BINARY_NAME) ## Build API
build-api:
	@go build -v $(GOARGS) ./backend/poc/main.go


.PHONY: test
test: ## Run unit tests
	@cd api && go test  -covermode=count ./...
	@cd simulator && go test  -covermode=count ./...
	@cd shared && go test  -covermode=count ./...


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
