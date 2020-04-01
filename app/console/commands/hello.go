package commands

import (
	"fmt"
	"github.com/spf13/cobra"
)

var HelloCommand = &cobra.Command{
	Use:                "hello",
	Aliases:            []string{"say"},
	Short:              "say hello.",
	DisableFlagParsing: true,
	RunE: func(c *cobra.Command, args []string) error {
		//redis.Init()
		//spew.Dump(productRepo.GetProductsFromCacheOrRpcThenSetCache(Ctx, []int{2455}))
		for _, msg := range args {
			fmt.Println(msg)
		}
		return nil
	},
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}
