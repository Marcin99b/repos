package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "repos",
	Short: "CLI for easier management of git repositories on local computer",
	Long: `repos is a CLI for easier management of git repositories on local computer.

It lets you configure a single working directory as your repositories store,
clone repositories into it, and list the repositories already there.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}
