package cmd

import (
	"fmt"
	"marcin99b/repos/internal"
	"os"

	"github.com/spf13/cobra"
)

var wdCmd = &cobra.Command{
	Use:   "wd",
	Short: "Get or set working directory (repositories dictionary)",
	Long:  `wd without parameter will return working directory, but wd with argument will update working directory`,
	Run:   runWd,
}

func init() {
	rootCmd.AddCommand(wdCmd)
}

func runWd(cmd *cobra.Command, args []string) {
	if len(args) > 1 {
		fmt.Println("Too many arguments...")
		return
	}

	if len(args) == 0 {
		value, err := get()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println(value)
	} else {
		err := update(args[0])
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}
}

func get() (string, error) {
	c, err := internal.ReadConfig()
	return c.Wd, err
}

func update(value string) error {
	wd, err := os.Stat(value)
	if err != nil || !wd.IsDir() {
		return fmt.Errorf("%s is not a dictionary", value)
	}

	c, err := internal.ReadConfig()
	if err != nil {
		return err
	}

	c.Wd = value
	return internal.UpdateConfig(c)
}
