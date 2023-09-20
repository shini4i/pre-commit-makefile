# pre-commit-makefile

This project addresses a specific and straightforward use case. It enables users to dynamically incorporate descriptions of Makefile targets into their README.md.

## Prerequisites

The binary can be installed using homebrew:

```bash
brew install shini4i/tap/pre-commit-makefile
````

## Usage

To use this project, add the following to your `.pre-commit-config.yaml`:

```yaml
- repo: https://github.com/shini4i/pre-commit-makefile
  rev: v0.1.0
  hooks:
    - id: makefile-readme-updater
```

To use this project, add the following to your README.md:

```markdown
<!-- BEGINNING OF PRE-COMMIT-MAKEFILE HOOK -->
<!-- END OF PRE-COMMIT-MAKEFILE HOOK -->
```

The following content will be dynamically generated and inserted between the above markers:
```
## Makefile targets

▷ `install-deps`: Install dependencies

▷ `build`: Build project binary

▷ `test`: Run tests

▷ `test-coverage`: Run tests with coverage

▷ `clean`: Remove build artifacts
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.