package main

import (
	"fmt"

	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

const version = "0.1.5"

func cli() error {
	var readmePath string
	var sectionName string
	var printVersion bool

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

	var rootCmd = &cobra.Command{
		Use: "pre-commit-makefile",
		Run: func(cmd *cobra.Command, args []string) {
			if printVersion {
				fmt.Println(version)
				return
			}

			// Else: Display default help
			if err := cmd.Help(); err != nil {
				return
			}
		},
	}

	rootCmd.PersistentFlags().BoolVarP(&printVersion, "version", "v", false, "Print version and exit")
	rootCmd.AddCommand(cmdRun)

	if err := rootCmd.Execute(); err != nil {
		return err
	}
	return nil
}
