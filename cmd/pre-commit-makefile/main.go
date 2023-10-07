package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/spf13/afero"
)

type Target struct {
	Name        string
	Description string
}

type App struct {
	Fs afero.Fs
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

func UpdateReadme(fs afero.Fs, targets []Target, readmePath, readmeSectionName string) error {
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
	newContent.WriteString(fmt.Sprintf("%s\n\n", readmeSectionName))

	for _, target := range targets {
		if target.Name != "help" {
			newContent.WriteString(fmt.Sprintf("To %s run:\n\n", target.Description))
			newContent.WriteString(fmt.Sprintf("```bash\nmake %s\n```\n\n", target.Name))
		}
	}

	newContent.WriteString(string(currentContent[endPos:]))

	return afero.WriteFile(fs, readmePath, []byte(newContent.String()), 0644)
}

func (app *App) Run(readmePath, readmeSectionName string) error {
	file, err := app.Fs.Open("Makefile")
	if err != nil {
		return fmt.Errorf("error opening Makefile: %s", err)
	}

	targets := ParseMakefile(file)

	if err := UpdateReadme(app.Fs, targets, readmePath, readmeSectionName); err != nil {
		return fmt.Errorf("error updating README.md: %s", err)
	}

	return nil
}

func main() {
	if err := cli(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
