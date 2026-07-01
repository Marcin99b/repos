package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var defdirCmd = &cobra.Command{
	Use:   "defdir",
	Short: "Get or set default repos directory",
	Long: `defdir without parameter will return default directory, 
	but defdir with argument will update default directory`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("defdir called")
	},
}

func init() {
	rootCmd.AddCommand(defdirCmd)
}
