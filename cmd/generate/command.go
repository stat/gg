package generate

import (
	"fmt"

	"github.com/spf13/cobra"

	"gg/cmd/generate/env"
)

func init() {
	// sub commands

	Command.AddCommand(env.Command)
}

var (
	Command = &cobra.Command{
		Use:   "generate",
		Short: "Runs all generators",
		Long:  `Runs all generators`,
		Run:   Run,
	}
)

func Run(cmd *cobra.Command, args []string) {
	fmt.Println("Generate Command")
}
