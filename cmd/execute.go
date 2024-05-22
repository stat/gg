package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"gg/cmd/generate"
)

func init() {
	// sub commands

	Command.AddCommand(generate.Command)
}

var Command = &cobra.Command{
	Use:   "gg",
	Short: "gg",
	Long:  `gg`,
	Run:   Run,
}

func Run(cmd *cobra.Command, args []string) {
	fmt.Println("Root Command")
}

func Execute() {
	if err := Command.Execute(); err != nil {
		os.Exit(1)
	}
}
