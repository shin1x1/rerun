package cmd

import (
	"fmt"
	"github.com/shin1x1/rerun/version"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version",
	Long:  "Show version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version.Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
