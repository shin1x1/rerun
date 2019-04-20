package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"sync"
	"time"
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

		var wg sync.WaitGroup

		for {
			wg.Add(1)
			go func(command string, args []string) {
				defer wg.Done()
				fmt.Println(command, args)

				cmd := exec.Command(command, args...)
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr

				if err := cmd.Run(); err != nil {
					fmt.Println(err)
				}
				time.Sleep(time.Duration(sleep) * time.Second)

			}(command, commandArgs)

			wg.Wait()
		}
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
