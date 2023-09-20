package main

import (
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseMakefile(t *testing.T) {
	fs := afero.NewMemMapFs()
	err := afero.WriteFile(fs, "Makefile", []byte(`
	.PHONY: build
	build: ## Build the project
		@echo "Building..."

	.PHONY: clean
	clean: ## Clean build artifacts
		@rm -rf bin/

	some-rule:
		@echo "This rule has no description"
	`), 0644)

	assert.NoError(t, err)

	file, _ := fs.Open("Makefile")
	targets := ParseMakefile(file)

	expectedTargets := []Target{
		{
			Name:        "build",
			Description: "Build the project",
		},
		{
			Name:        "clean",
			Description: "Clean build artifacts",
		},
	}

	assert.Equal(t, expectedTargets, targets)
}

func TestUpdateReadme(t *testing.T) {
	fs := afero.NewMemMapFs()

	readmeContent := `
# Project Title

<!-- BEGINNING OF PRE-COMMIT-MAKEFILE HOOK -->

<!-- END OF PRE-COMMIT-MAKEFILE HOOK -->

## Another Section
`

	err := afero.WriteFile(fs, "README.md", []byte(readmeContent), 0644)
	assert.NoError(t, err)

	targets := []Target{
		{
			Name:        "build",
			Description: "Build the project",
		},
		{
			Name:        "clean",
			Description: "Clean build artifacts",
		},
	}

	// Assuming the function is adapted to use afero Fs
	err = UpdateReadme(fs, targets, "README.md")
	assert.NoError(t, err)

	updatedContent, _ := afero.ReadFile(fs, "README.md")
	expectedContent := `
# Project Title

<!-- BEGINNING OF PRE-COMMIT-MAKEFILE HOOK -->
## Makefile targets

▷ ` + "`build`" + `: Build the project

▷ ` + "`clean`" + `: Clean build artifacts

<!-- END OF PRE-COMMIT-MAKEFILE HOOK -->

## Another Section
`
	assert.Equal(t, expectedContent, string(updatedContent))
}
