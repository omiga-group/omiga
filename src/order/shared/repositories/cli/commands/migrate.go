package commands

import (
	"context"
	"log"
	"os"

	"github.com/mitchellh/mapstructure"
	"github.com/omiga-group/omiga/src/order/shared/repositories/migrate"
	"github.com/omiga-group/omiga/src/shared/enterprise/database/postgres"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func migrateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "migrate",
		Short: "Migrate database to the latest version",
		Long:  "Migrate database to the latest version",
		Run: func(cmd *cobra.Command, args []string) {
			logger, err := zap.NewDevelopment()
			if err != nil {
				log.Fatal(err)
			}

			sugarLogger := logger.Sugar()

			var postgresSettings postgres.PostgresSettings
			if err := mapstructure.Decode(viper.Get(postgres.ConfigKey), &postgresSettings); err != nil {
				sugarLogger.Fatal(err)
			}

			entgoClient, err := NewEntgoClient(sugarLogger, postgresSettings)
			if err != nil {
				sugarLogger.Fatal(err)
			}

			defer entgoClient.Close()

			client := entgoClient.GetClient()

			ctx := context.Background()

			if err = client.Schema.WriteTo(
				ctx,
				os.Stdout,
				migrate.WithGlobalUniqueID(true)); err != nil {
				sugarLogger.Fatal(err)
			}

			if err = client.Schema.Create(
				ctx,
				migrate.WithGlobalUniqueID(true)); err != nil {
				sugarLogger.Fatal(err)
			}

			sugarLogger.Info("Successfully migrated database to the latest schema")
		},
	}

	return cmd
}
