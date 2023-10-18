.DEFAULT_GOAL := help

.PHONY: help
help: ## print this help
	@echo "Usage: make [target]"
	@grep -E '^[a-z.A-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: install-deps
install-deps: ## install dependencies
	@echo "===> Installing dependencies"
	@go get ./...

.PHONY: ensure-dir
ensure-dir:
	@mkdir -p bin

.PHONY: build
build: ensure-dir ## build project binary
	@echo "Building project binary..."
	@@CGO_ENABLED=0 go build -ldflags="-s -w" -o bin/pre-commit-makefile ./cmd/pre-commit-makefile

.PHONY: test
test: ## run tests
	@go test -v ./... -count=1

.PHONY: test-coverage
test-coverage: ## run tests with coverage
	@go test -v -coverprofile=coverage.out ./... -count=1 2>&1

.PHONY: clean
clean: ## remove build artifacts
	@rm -rf bin

.PHONY: bump-patch
bump-patch:
	@bump2version patch --allow-dirty

.PHONY: bump-minor
bump-minor:
	@bump2version minor --allow-dirty
