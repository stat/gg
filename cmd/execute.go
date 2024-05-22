package cmd

import (
	"os"

	"gg/cmd/root"
)

func init() {}

func Execute() {
	if err := root.Command.Execute(); err != nil {
		os.Exit(1)
	}
}
