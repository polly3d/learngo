package cobraExample

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number.",
	Long:  "Let me show you version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("version 0.1")
	},
}
