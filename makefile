PROJECT_NAME := "algostructures"
PKG := "github.com/martinomburajr/$(PROJECT_NAME)"
PKG_LIST := $(shell go list ${PKG}/... | grep -v /vendor/)
GO_FILES := $(shell find . -name '*.go' | grep -v /vendor/ | grep -v _test.go)

.PHONY: all dep build clean test coverage coveragehtml lint

all: build

lint: dep
	@golint -set_exit_status ${PKG_LIST}

test: ## Run unit tests
	@go test -short ${PKG_LIST}

race: dep ## Run data race detector
	@go test -race -short ${PKG_LIST}

dep: ## Get dependencies
	@go get -v -d ./...

clean: ## Remove previous builds
	@rm -f ${PROJECT_NAME}

cover: ## reports coverage for tests
	./coverage.sh;

coverhtml: ## Generage global code coverage report in HTML
	./coverage.sh html;

help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'