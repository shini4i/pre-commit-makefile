package main

import (
	"bufio"
	"fmt"
	"regexp"
	"strings"

	"github.com/spf13/afero"
)

var (
	targetRegex = regexp.MustCompile(`^([a-zA-Z._-]+):`)
	phonyRegex  = regexp.MustCompile(`^\.PHONY: (.*)$`)
)

func ValidateMakefile(fs afero.Fs, makefilePath string) error {
	file, err := fs.Open(makefilePath)
	if err != nil {
		return fmt.Errorf("error opening Makefile: %w", err)
	}

	defer func(file afero.File) {
		if err := file.Close(); err != nil {
			fmt.Printf("error closing Makefile: %v", err)
		}
	}(file)

	var targets []string
	phonies := make(map[string]bool)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		targetMatches := targetRegex.FindStringSubmatch(line)
		phonyMatches := phonyRegex.FindStringSubmatch(line)

		if len(targetMatches) == 2 && !strings.HasPrefix(line, ".PHONY:") {
			targets = append(targets, targetMatches[1])
		}

		if len(phonyMatches) == 2 {
			phonies[phonyMatches[1]] = true
		}
	}

	for _, target := range targets {
		if !phonies[target] {
			return fmt.Errorf("target '%s' does not have a corresponding .PHONY definition", target)
		}
	}

	return nil
}
