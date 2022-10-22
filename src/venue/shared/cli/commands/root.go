package commands

import (
	"github.com/micro-business/go-core/pkg/util"
	"github.com/spf13/cobra"
)

// RootCommand returns root CLI application command interface
func Root() *cobra.Command {
	var cmd = &cobra.Command{
		Use: "venue-database-migration",
		PreRun: func(cmd *cobra.Command, args []string) {
			printHeader()
		},
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	// Register all commands
	cmd.AddCommand(
		versionCommand(),
		migrateCommand(),
	)

	return cmd
}

func printHeader() {
	util.PrintInfo("Exchange Database Migration Tool")
}
