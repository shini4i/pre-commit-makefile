package main

import (
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func TestValidateMakefile(t *testing.T) {

	// Creating a new in-memory file system
	fs := afero.NewMemMapFs()
	makefilePath := "Makefile"

	t.Run("Correct Makefile", func(t *testing.T) {
		makefileContent := `
.PHONY: start-server
start-server:

.PHONY: test-build
test-build:
`
		if err := afero.WriteFile(fs, makefilePath, []byte(makefileContent), 0644); err != nil {
			t.Fatal(err)
		}

		err := ValidateMakefile(fs, makefilePath)
		assert.NoError(t, err, "ValidateMakefile should not return an error for a correct Makefile")
	})

	t.Run("Malformed Makefile", func(t *testing.T) {
		makefileContent := `
.PHONY: start-server
start-server:

test-build:`

		if err := afero.WriteFile(fs, makefilePath, []byte(makefileContent), 0644); err != nil {
			t.Fatal(err)
		}

		err := ValidateMakefile(fs, makefilePath)
		assert.Error(t, err, "ValidateMakefile should return an error for a malformed Makefile")
	})

	t.Run("Non-existent file", func(t *testing.T) {
		err := ValidateMakefile(fs, "/nonexistent/path")
		assert.Error(t, err, "ValidateMakefile should return an error for a non-existent Makefile")
	})
}
