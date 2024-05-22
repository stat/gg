package aquire_types

import (
	"fmt"
	"math/rand"

	"github.com/dave/jennifer/jen"
	"github.com/spf13/cobra"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	// . "github.com/dave/jennifer/jen"
)

/*
	//go:generate go run ../main.go generate env tests aquire-types $GOFILE
*/

var (
	Command = &cobra.Command{
		Use:   "aquire-types [target]",
		Short: "Runs all env tests generators",
		Long:  `Runs all env tests generators`,
		Args:  cobra.MatchAll(cobra.ExactArgs(2), cobra.OnlyValidArgs),
		Run:   Run,
	}
)

func Run(cmd *cobra.Command, args []string) {
	fmt.Println("Generate ENV Aquire Types Tests Command")

	pkg := args[0]
	src := args[1]

	tests := []struct {
		Type  string
		Value any
	}{
		{Type: "int", Value: rand.Int()},
	}

	f := jen.NewFile(pkg)
	caser := cases.Title(language.English)

	for _, test := range tests {
		fId := fmt.Sprintf("TestAquire%s", caser.String(test.Type))

		f.Func().Id(fId).Params().Block(
			jen.Qual("fmt", "Println").Call(jen.Lit("Hello, world")),
		)
	}

	fmt.Println("writing:", src)
	fmt.Printf("%#v", f)
	// f.Save(target)
}
