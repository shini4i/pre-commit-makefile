package main

import (
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

const defaultReadmePath = "README.md"
const defaultReadmeSectionName = "## Makefile targets"

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

	err := afero.WriteFile(fs, defaultReadmePath, []byte(readmeContent), 0644)
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
	err = UpdateReadme(fs, targets, defaultReadmePath, defaultReadmeSectionName)
	assert.NoError(t, err)

	updatedContent, _ := afero.ReadFile(fs, "README.md")

	expectedContent := "\n# Project Title\n\n" +
		"<!-- BEGINNING OF PRE-COMMIT-MAKEFILE HOOK -->\n" +
		"## Makefile targets\n\n" +
		"To Build the project run:\n\n" +
		"```bash\n" +
		"make build\n" +
		"```\n\n" +
		"To Clean build artifacts run:\n\n" +
		"```bash\n" +
		"make clean\n" +
		"```\n\n" +
		"<!-- END OF PRE-COMMIT-MAKEFILE HOOK -->\n\n" +
		"## Another Section\n"

	assert.Equal(t, expectedContent, string(updatedContent))
}

func TestApp_Run(t *testing.T) {
	// Positive test
	t.Run("with valid Makefile and README.md", func(t *testing.T) {
		fs := afero.NewMemMapFs()

		// Setup Makefile
		err := afero.WriteFile(fs, "Makefile", []byte(`
			.PHONY: build
			build: ## Build the project
				@echo "Building..."`,
		), 0644)
		assert.NoError(t, err)

		// Setup README.md
		readmeContent := `
			# Project Title

			<!-- BEGINNING OF PRE-COMMIT-MAKEFILE HOOK -->

			<!-- END OF PRE-COMMIT-MAKEFILE HOOK -->

			## Another Section
		`
		err = afero.WriteFile(fs, defaultReadmePath, []byte(readmeContent), 0644)
		assert.NoError(t, err)

		app := &App{Fs: fs}
		assert.NoError(t, app.Run(defaultReadmePath, defaultReadmeSectionName))
	})

	// Negative test: Missing Makefile
	t.Run("with missing Makefile", func(t *testing.T) {
		fs := afero.NewMemMapFs()
		app := &App{Fs: fs}
		assert.Error(t, app.Run(defaultReadmePath, defaultReadmeSectionName))
	})

	// Negative test: Missing hook tags in README.md
	t.Run("with missing hook tags in README.md", func(t *testing.T) {
		fs := afero.NewMemMapFs()

		// Setup README.md without hook tags
		err := afero.WriteFile(fs, "README.md", []byte("# Project Title\n\n## Another Section"), 0644)
		assert.NoError(t, err)

		app := &App{Fs: fs}
		assert.Error(t, app.Run(defaultReadmePath, defaultReadmeSectionName))
	})
}
