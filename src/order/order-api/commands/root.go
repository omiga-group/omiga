package commands

import (
	"github.com/micro-business/go-core/pkg/util"
	"github.com/spf13/cobra"
)

// RootCommand returns root CLI application command interface
func Root() *cobra.Command {
	var cmd = &cobra.Command{
		Use: "order-api",
		PreRun: func(cmd *cobra.Command, args []string) {
			printHeader()
		},
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	cmd.AddCommand(
		versionCommand(),
		installDependeciesCommand(),
		startCommand(),
	)

	return cmd
}

func printHeader() {
	util.PrintInfo("Order API")
}
