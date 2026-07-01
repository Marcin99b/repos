package cmd

import (
	"fmt"
	"marcin99b/repos/internal"
	"os"

	"github.com/spf13/cobra"
)

var defdirCmd = &cobra.Command{
	Use:   "defdir",
	Short: "Get or set default repos directory",
	Long: `defdir without parameter will return default directory, 
	but defdir with argument will update default directory`,
	Run: runDefdir,
}

func init() {
	rootCmd.AddCommand(defdirCmd)
}

func runDefdir(cmd *cobra.Command, args []string) {
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
	return c.Defdir, err
}

func update(value string) error {
	defdir, err := os.Stat(value)
	if err != nil || !defdir.IsDir() {
		return fmt.Errorf("%s is not a dictionary", value)
	}

	c, err := internal.ReadConfig()
	if err != nil {
		return err
	}

	c.Defdir = value
	return internal.UpdateConfig(c)
}
