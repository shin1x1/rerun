package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"shin1x1/rerun/handlers"
)

var (
	sleep int
)

var rootCmd = &cobra.Command{
	Use:   "rerun [command] [command args]",
	Short: "Rerun is a command re-runner",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		command := args[0]
		var commandArgs []string
		if len(args) > 1 {
			commandArgs = args[1:]
		}

		handlers.NewRerun(command, commandArgs, sleep).Run()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().IntVarP(&sleep, "sleep", "s", 10, "sleep time(sec) before re-run command")
}
