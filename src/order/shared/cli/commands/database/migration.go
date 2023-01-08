package database

import (
	"github.com/omiga-group/omiga/src/order/shared/cli/commands/database/migration"
	"github.com/spf13/cobra"
)

func MigrateCommand(connectionString *string) *cobra.Command {
	cmd := &cobra.Command{
		Use: "migration",
	}

	cmd.AddCommand(
		migration.AddCommand(connectionString),
	)

	return cmd
}
