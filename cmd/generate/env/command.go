package env

import (
	"fmt"

	"github.com/spf13/cobra"

	"gg/cmd/generate/env/tests"
)

func init() {
	// sub commands

	Command.AddCommand(tests.Command)
}

var (
	Command = &cobra.Command{
		Use:   "env",
		Short: "Runs all env generators",
		Long:  `Runs all env generators`,
		Run:   Run,
	}
)

func Run(cmd *cobra.Command, args []string) {
	fmt.Println("Generate ENV Command")
}
