package generate

import (
	"fmt"

	"github.com/spf13/cobra"

	"gg/cmd/root"
)

func init() {
	root.Command.AddCommand(Command)
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
