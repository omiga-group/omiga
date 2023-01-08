package database

import (
	"github.com/omiga-group/omiga/src/shared/enterprise/logger"
	"github.com/omiga-group/omiga/src/shared/omigactl/commands/database/migrate"
	"github.com/spf13/cobra"
)

type migrateOptions struct {
	path string
}

func MigrateCommand(connectionString *string) *cobra.Command {
	options := migrateOptions{}
	sugarLogger := logger.CreateLogger()

	cmd := &cobra.Command{
		Use: "migrate",
	}

	cmd.Flags().StringVar(&options.path, "path", "", "Specify the path to the migration scripts")

	if err := cmd.MarkFlagRequired("path"); err != nil {
		sugarLogger.Fatal(err)
	}

	cmd.AddCommand(
		migrate.AddCommand(connectionString, &options.path),
	)

	return cmd
}
