package cmd

import (	
	"os"
	"ilias-cli/cmd/exercises"
	"ilias-cli/cmd/grades"
	"ilias-cli/cmd/members"
	"ilias-cli/cmd/workspace"	
	"github.com/spf13/cobra"
)

var rootCommand = &cobra.Command{
	Use:   "ilias",
	Short: "A simple command line interface for managing ILIAS",
	SilenceErrors: true,

	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(1)
	},
}

func init() {
	rootCommand.AddCommand(exercises.RootCommand)
	rootCommand.AddCommand(members.RootCommand)
	rootCommand.AddCommand(workspace.RootCommand)
	rootCommand.AddCommand(grades.RootCommand)
	
}

func Execute() {
	if err := rootCommand.Execute(); err != nil {
		os.Exit(1)
	}
}
