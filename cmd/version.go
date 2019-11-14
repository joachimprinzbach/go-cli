package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of go-cli",
	Long:  `All software has versions. This is go-cli's'`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("go-cli 0.0.1-buggy-shizzl -- HEAD")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
