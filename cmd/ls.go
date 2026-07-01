package cmd

import (
	"fmt"
	"marcin99b/repos/internal"
	"os"

	"github.com/spf13/cobra"
)

var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List repositories in working directory",
	Long:  `List repositories in working directory`,
	Run:   runLs,
}

func init() {
	rootCmd.AddCommand(lsCmd)
}

func runLs(cmd *cobra.Command, args []string) {
	c, err := internal.ReadConfig()
	if err != nil || len(c.Wd) == 0 {
		fmt.Println("wd is empty")
		return
	}

	entries, err := os.ReadDir(c.Wd)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, e := range entries {
		if !e.Type().IsDir() {
			continue
		}

		fmt.Println(e.Name())
	}
}
