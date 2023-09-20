package main

import (
	"bufio"
	"fmt"
	"github.com/spf13/afero"
	"os"
	"regexp"
	"strings"
)

type Target struct {
	Name        string
	Description string
}

func ParseMakefile(file afero.File) []Target {
	var targets []Target
	targetRegex := regexp.MustCompile(`^([a-zA-Z._-]+):.*?## (.*)$`)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		matches := targetRegex.FindStringSubmatch(line)
		if len(matches) == 3 {
			target := Target{
				Name:        matches[1],
				Description: matches[2],
			}
			targets = append(targets, target)
		}
	}
	return targets
}

func UpdateReadme(fs afero.Fs, targets []Target, readmePath string) error {
	currentContent, err := afero.ReadFile(fs, readmePath)
	if err != nil {
		return err
	}

	beginTag := "<!-- BEGINNING OF PRE-COMMIT-MAKEFILE HOOK -->"
	endTag := "<!-- END OF PRE-COMMIT-MAKEFILE HOOK -->"

	beginPos := strings.Index(string(currentContent), beginTag)
	endPos := strings.Index(string(currentContent), endTag)

	if beginPos == -1 || endPos == -1 {
		return fmt.Errorf("README.md does not contain required hook tags")
	}

	var newContent strings.Builder
	newContent.WriteString(string(currentContent[:beginPos+len(beginTag)]) + "\n")
	newContent.WriteString("## Makefile targets\n\n")

	for _, target := range targets {
		if target.Name != "help" {
			newContent.WriteString(fmt.Sprintf("▷ `%s`: %s\n\n", target.Name, target.Description))
		}
	}

	newContent.WriteString(string(currentContent[endPos:]))

	return afero.WriteFile(fs, readmePath, []byte(newContent.String()), 0644)
}

func main() {
	fs := afero.NewOsFs()

	file, err := fs.Open("Makefile")
	if err != nil {
		fmt.Printf("Error opening Makefile: %s", err)
		os.Exit(1)
	}

	targets := ParseMakefile(file)

	if err := UpdateReadme(fs, targets, "README.md"); err != nil {
		fmt.Printf("Error updating README.md: %s", err)
		os.Exit(1)
	}
}
