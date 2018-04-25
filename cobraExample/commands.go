package cobraExample

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "learn",
	Short: "learn from cobra, a command-line tool.",
	Long:  `Help to learn from cobra to generate a command-line tool for this project`,
}

var cmdEcho = &cobra.Command{
	Use:   "echo [string to echo]",
	Short: "Echo every anyting to scree",
	Long:  "Echo Echo Echo",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Print: " + strings.Join(args, " "))
	},
}

var echoTimes int
var cmdTimes = &cobra.Command{
	Use:   "times [# times] [string to echo]",
	Short: "echo more times",
	Long:  "echo echo",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for i := 0; i < echoTimes; i++ {
			fmt.Println("Echo: " + strings.Join(args, " "))
		}
	},
}

//Execute is a cobra func
func Execute() {
	cmdTimes.Flags().IntVarP(&echoTimes, "times", "t", 1, "times to echo the input")
	cmdEcho.AddCommand(cmdTimes)
	rootCmd.AddCommand(cmdEcho)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
