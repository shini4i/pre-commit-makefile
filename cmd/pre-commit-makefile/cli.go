package main

import (
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

func cli() error {
	var readmePath string
	var sectionName string

	var cmdRun = &cobra.Command{
		Use:   "run",
		Short: "Run the pre-commit-makefile",
		RunE: func(cmd *cobra.Command, args []string) error {
			app := &App{
				Fs: afero.NewOsFs(),
			}

			if err := app.Run(readmePath, sectionName); err != nil {
				return err
			}

			return nil
		},
	}

	cmdRun.Flags().StringVarP(&readmePath, "readme-path", "r", "README.md", "Path to the readme file")
	cmdRun.Flags().StringVarP(&sectionName, "section-name", "s", "## Makefile targets", "Readme section name to put the description in")

	var rootCmd = &cobra.Command{Use: "pre-commit-makefile"}
	rootCmd.AddCommand(cmdRun)

	if err := rootCmd.Execute(); err != nil {
		return err
	}
	return nil
}
