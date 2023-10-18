<div align="center">

# 🛠 pre-commit-makefile 🛠

This project allows users to automatically update their README.md with descriptions of Makefile targets.

![GitHub Workflow Status](https://img.shields.io/github/actions/workflow/status/shini4i/pre-commit-makefile/run-tests.yml?branch=main)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/shini4i/pre-commit-makefile)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/shini4i/pre-commit-makefile)
[![codecov](https://codecov.io/gh/shini4i/pre-commit-makefile/graph/badge.svg?token=JXN63AUFFW)](https://codecov.io/gh/shini4i/pre-commit-makefile)
[![Go Report Card](https://goreportcard.com/badge/github.com/shini4i/pre-commit-makefile)](https://goreportcard.com/report/github.com/shini4i/pre-commit-makefile)
![GitHub](https://img.shields.io/github/license/shini4i/pre-commit-makefile)

</div>

## Prerequisites

To use this project, you need to install `pre-commit-makefile` binary. You can do this by running:

```bash
brew install shini4i/tap/pre-commit-makefile
````

Or by downloading the desired version from [releases](https://github.com/shini4i/pre-commit-makefile/releases) page.

The expected `Makefile` format is the following:

```makefile
.PHONY: help
help: ## print this help
	@echo "Usage: make [target]"
	@grep -E '^[a-z.A-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: test
test: ## run tests
	@go test -v ./... -count=1
```

> **Note**: The content after  `##` will be used as a target description

## Configuration

To start using this project, add the following to your `.pre-commit-config.yaml`:

```yaml
repos:
  - repo: https://github.com/shini4i/pre-commit-makefile
    rev: v0.1.4 # Replace with the latest release version
    hooks:
      - id: makefile-readme-updater
```

The following comment markers should be added to your `README.md`:

```markdown
<!-- BEGINNING OF PRE-COMMIT-MAKEFILE HOOK -->
<!-- END OF PRE-COMMIT-MAKEFILE HOOK -->
```

The dynamically generated content will be placed between the markers.

Additionally, it is possible to override a few settings by adding the following arguments to your `.pre-commit-config.yaml`:

```yaml
args:
  - --readme-path=docs/README.md
  - --section-name=## Usage
```

## Example
The `Makefile` in this repository will produce the following output:
<!-- BEGINNING OF PRE-COMMIT-MAKEFILE HOOK -->
To install dependencies run:

```bash
make install-deps
```

To build project binary run:

```bash
make build
```

To run tests run:

```bash
make test
```

To run tests with coverage run:

```bash
make test-coverage
```

To remove build artifacts run:

```bash
make clean
```

> Note: The generated content will be placed under `## Makefile targets` section.

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.
