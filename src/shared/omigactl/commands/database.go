package commands

import (
	"github.com/omiga-group/omiga/src/shared/omigactl/commands/database"
	"github.com/spf13/cobra"
)

type databaseOptions struct {
	connectionString string
}

func databaseCommand() *cobra.Command {
	options := databaseOptions{}

	cmd := &cobra.Command{
		Use: "database",
	}

	cmd.PersistentFlags().StringVar(&options.connectionString, "connectionString", "", "Specify the database connection string")

	cmd.AddCommand(
		database.ProvisionCommand(&options.connectionString),
		database.MigrateCommand(&options.connectionString),
	)

	return cmd
}
