package commands

import (
	"github.com/omiga-group/omiga/src/shared/omigactl/commands/database"
	"github.com/spf13/cobra"
)

type databaseOptions struct {
	connectionString string
}

func databaseCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "database",
		PreRun: func(cmd *cobra.Command, args []string) {
			printHeader()
		},
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	options := databaseOptions{}

	cmd.PersistentFlags().StringVar(&options.connectionString, "connectionString", "", "Specify the database connection string")

	// Register all commands
	cmd.AddCommand(
		database.CreateCommand(cmd),
	)

	return cmd
}
