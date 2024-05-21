package root

import (
	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use:   "help",
	Short: "gg",
	Long:  `gg`,
	Run:   func(cmd *cobra.Command, args []string) {},
}
