package migrate

import (
	"context"

	"github.com/omiga-group/omiga/src/shared/enterprise/appsetup"
	"github.com/omiga-group/omiga/src/shared/enterprise/logger"
	"github.com/omiga-group/omiga/src/shared/omigactl/configuration"
	"github.com/spf13/cobra"
)

type provisionDatabaseOptions struct {
	name string
}

func AddCommand(connectionString *string, path *string) *cobra.Command {
	options := provisionDatabaseOptions{}
	sugarLogger := logger.CreateLogger()

	cmd := &cobra.Command{
		Use:   "add",
		Short: "Add new migration",
		Long:  "Add new migration",
		Run: func(cmd *cobra.Command, args []string) {
			_ = context.Background()

			configurationHelper, err := appsetup.NewConfigurationHelper(sugarLogger)
			if err != nil {
				sugarLogger.Fatal(err)
			}

			var config configuration.Config
			if err := configurationHelper.LoadYaml("config.yaml", &config); err != nil {
				sugarLogger.Fatal(err)
			}

			if len(*connectionString) != 0 {
				config.Postgres.ConnectionString = *connectionString
			}
		},
	}

	cmd.Flags().StringVar(&options.name, "name", "", "Specify the database name to provision")

	if err := cmd.MarkFlagRequired("name"); err != nil {
		sugarLogger.Fatal(err)
	}

	return cmd
}
