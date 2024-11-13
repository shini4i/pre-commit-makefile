package main

import (
	"fmt"

	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

const version = "0.2.1"

func runCommand(readmePath *string, sectionName *string) *cobra.Command {
	return &cobra.Command{
		Use:   "run",
		Short: "Run the pre-commit-makefile",
		RunE: func(cmd *cobra.Command, args []string) error {
			app := &App{
				Fs: afero.NewOsFs(),
			}

			if err := app.Run(*readmePath, *sectionName); err != nil {
				return err
			}

			return nil
		},
	}
}

func validateCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "validate",
		Short: "Validate the pre-commit-makefile",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := ValidateMakefile(afero.NewOsFs(), "Makefile"); err != nil {
				return err
			}
			return nil
		},
		SilenceUsage:  true,
		SilenceErrors: true,
	}
}

func rootCommand(printVersion *bool) *cobra.Command {
	cmd := &cobra.Command{
		Use: "pre-commit-makefile",
		Run: func(cmd *cobra.Command, args []string) {
			if *printVersion {
				fmt.Println(version)
				return
			}

			// Else: Display default help
			if err := cmd.Help(); err != nil {
				fmt.Println(err)
			}
		},
	}

	cmd.PersistentFlags().BoolVarP(printVersion, "version", "v", false, "Print version and exit")
	return cmd
}

func cli() error {
	var readmePath string
	var sectionName string
	var printVersion bool

	cmdRun := runCommand(&readmePath, &sectionName)
	cmdRun.Flags().StringVarP(&readmePath, "readme-path", "r", "README.md", "Path to the readme file")
	cmdRun.Flags().StringVarP(&sectionName, "section-name", "s", "## Makefile targets", "Readme section name to put the description in")

	rootCmd := rootCommand(&printVersion)
	rootCmd.AddCommand(cmdRun)
	rootCmd.AddCommand(validateCommand())

	return rootCmd.Execute()
}
