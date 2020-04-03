package commands

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ServerCommand = &cobra.Command{
	Use:                "hello",
	Aliases:            []string{"say"},
	Short:              "say hello.",
	DisableFlagParsing: true,
	RunE: func(c *cobra.Command, args []string) error {
		fmt.Println("test server")
		return nil
	},
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}
