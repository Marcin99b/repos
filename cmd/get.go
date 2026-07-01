package cmd

import (
	"fmt"
	"marcin99b/repos/internal"
	"net/url"
	"os"
	"path"
	"strings"

	"github.com/go-git/go-git/v6"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get <url>",
	Short: "Clone git repository",
	Long:  `Clone git repository to directory {wd}`,
	Args:  cobra.ExactArgs(1),
	Run:   runGet,
}

func init() {
	rootCmd.AddCommand(getCmd)
}

func runGet(cmd *cobra.Command, args []string) {
	c, err := internal.ReadConfig()
	if err != nil || len(c.Wd) == 0 {
		fmt.Println("wd is empty")
		return
	}

	uri, err := url.Parse(args[0])
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	splitted := strings.Split(uri.Path, "/")
	if len(uri.Path) == 0 || len(splitted) == 0 {
		fmt.Println("Url has no project name")
		return
	}

	projectName := splitted[len(splitted)-1]
	projectName = strings.Replace(projectName, ".git", "", 1)
	path := path.Join(c.Wd, projectName)

	_, err = git.PlainClone(path, &git.CloneOptions{
		URL:      args[0],
		Progress: os.Stdout,
	})
}
