package tests

import (
	"fmt"

	"github.com/spf13/cobra"

	"gg/cmd/generate/env"
)

func init() {
	env.Command.AddCommand(Command)
}

var (
	Command = &cobra.Command{
		Use:   "tests",
		Short: "Runs all env tests generators",
		Long:  `Runs all env tests generators`,
		Run:   Run,
	}
)

func Run(cmd *cobra.Command, args []string) {
	fmt.Println("Generate ENV Tests Command")
}
