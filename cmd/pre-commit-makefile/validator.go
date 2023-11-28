package main

import (
	"bufio"
	"fmt"
	"regexp"
	"strings"

	"github.com/spf13/afero"
)

func ValidateMakefile(fs afero.Fs, makefilePath string) error {
	file, err := fs.Open(makefilePath)
	if err != nil {
		return fmt.Errorf("error opening Makefile: %s", err)
	}
	defer file.Close()

	targetRegex := regexp.MustCompile(`^([a-zA-Z._-]+):`)
	phonyRegex := regexp.MustCompile(`^\.PHONY: (.*)$`)

	var targets []string
	var phonies []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		targetMatches := targetRegex.FindStringSubmatch(line)
		phonyMatches := phonyRegex.FindStringSubmatch(line)

		if len(targetMatches) == 2 && !strings.HasPrefix(line, ".PHONY:") {
			targets = append(targets, targetMatches[1])
		}

		if len(phonyMatches) == 2 {
			phonies = append(phonies, phonyMatches[1])
		}
	}

	fmt.Println(targets)
	fmt.Println(phonies)
	for _, target := range targets {
		found := false
		for _, phony := range phonies {
			if target == phony {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("target '%s' does not have a corresponding .PHONY definition", target)
		}
	}

	return nil
}
