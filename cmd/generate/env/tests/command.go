package tests

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"gg/cmd/generate/env/tests/aquire_types"
)

func init() {
	Command.AddCommand(aquire_types.Command)
}

var (
	Command = &cobra.Command{
		Use:   "tests [target directory]",
		Short: "Runs all env tests generators",
		Long:  `Runs all env tests generators`,
		// Args:  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
		Run: Run,
	}
)

func Run(cmd *cobra.Command, args []string) {
	fmt.Println("Generate ENV Tests Command")

	// pwd, err := os.Getwd()

	// if err != nil {
	//   panic(err)
	// }

	target := args[0]

	log.Printf("Writing ENV Tests to: %s", target)

	// f := NewFile(fmt.Sprintf("%s/aquire_types_test.go", target))
	// f.Func().Id("main").Params().Block(
	//   Qual("fmt", "Println").Call(Lit("Hello, world")),
	// )
	// fmt.Printf("%#v", f)
}
