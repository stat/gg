package root

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use:   "gg",
	Short: "gg",
	Long:  `gg`,
	Run:   Run,
}

func Run(cmd *cobra.Command, args []string) {
	fmt.Println("Root Command")
}
