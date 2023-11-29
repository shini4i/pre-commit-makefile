package main

import (
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func TestValidateMakefile(t *testing.T) {
	fs := afero.NewMemMapFs()

	testMakefilePath := "/test/Makefile"

	// Creating a fake Makefile
	makefileContent := `
.PHONY: start-server
start-server:

.PHONY: test-build
test-build:
`

	if err := afero.WriteFile(fs, testMakefilePath, []byte(makefileContent), 0644); err != nil {
		t.Fatalf("error writing Makefile: %v", err)
	}

	assert.NoError(t, ValidateMakefile(fs, testMakefilePath))

	// Creating a fake Makefile with a target without a .PHONY definition
	makefileContent = `
.PHONY: start-server
start-server:


test-build:`

	if err := afero.WriteFile(fs, testMakefilePath, []byte(makefileContent), 0644); err != nil {
		t.Fatalf("error writing Makefile: %v", err)
	}

	assert.Error(t, ValidateMakefile(fs, testMakefilePath))
}
