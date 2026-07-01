package cmd

import (
	"fmt"
	"marcin99b/repos/internal"
	"os"

	"github.com/go-git/go-git/v6"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Clone git repository",
	Long:  `Clone git repository to directory {defdir}`,
	Run:   runGet,
}

func init() {
	rootCmd.AddCommand(getCmd)
}

func runGet(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		fmt.Println("Arguments not provided")
		return
	}

	c, err := internal.ReadConfig()
	if err != nil || len(c.Defdir) == 0 {
		fmt.Println("defdir is empty")
		return
	}

	//todo create folder for repo
	_, err = git.PlainClone(c.Defdir, &git.CloneOptions{
		URL:      args[0],
		Progress: os.Stdout,
	})
}
