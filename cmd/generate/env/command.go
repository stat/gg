package env

import (
	"fmt"

	"github.com/spf13/cobra"

	"gg/cmd/generate"
)

func init() {
	generate.Command.AddCommand(Command)
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
