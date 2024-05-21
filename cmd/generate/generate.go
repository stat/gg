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
		Run:   Execute,
	}
)

func Execute(cmd *cobra.Command, args []string) {
	fmt.Println("I'M HERE!")
}
