<div align="center">

# pre-commit-makefile

This project allows users to automatically update their README.md with descriptions of Makefile targets.

![GitHub Workflow Status](https://img.shields.io/github/actions/workflow/status/shini4i/pre-commit-makefile/run-tests.yml?branch=main)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/shini4i/pre-commit-makefile)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/shini4i/pre-commit-makefile)
[![codecov](https://codecov.io/gh/shini4i/pre-commit-makefile/branch/main/graph/badge.svg?token=48E1OZHLPY)](https://codecov.io/gh/shini4i/pre-commit-makefile)
[![Go Report Card](https://goreportcard.com/badge/github.com/shini4i/pre-commit-makefile)](https://goreportcard.com/report/github.com/shini4i/pre-commit-makefile)
![GitHub](https://img.shields.io/github/license/shini4i/pre-commit-makefile)

</div>

## Prerequisites

The binary can be installed using homebrew:

```bash
brew install shini4i/tap/pre-commit-makefile
````

## Usage

To use this project, add the following to your `.pre-commit-config.yaml`:

```yaml
- repo: https://github.com/shini4i/pre-commit-makefile
  rev: v0.1.1 # Replace with the latest release version
  hooks:
    - id: makefile-readme-updater
```

The following comments should be added to your `README.md`:

```markdown
<!-- BEGINNING OF PRE-COMMIT-MAKEFILE HOOK -->
<!-- END OF PRE-COMMIT-MAKEFILE HOOK -->
```

The content between these markers will be dynamically generated and will look something like this:

```
## Makefile targets

▷ `install-deps`: Install dependencies

▷ `build`: Build project binary

▷ `test`: Run tests

▷ `test-coverage`: Run tests with coverage

▷ `clean`: Remove build artifacts
```

This project expects the following `Makefile` format, the content after  `##` will be used as a target description:

```makefile
.PHONY: help
help: ## Print this help
	@echo "Usage: make [target]"
	@grep -E '^[a-z.A-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: test
test: ## Run tests
	@go test -v ./... -count=1
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.
