package commands

import (
	"context"
	"log"
	"os"

	"github.com/omiga-group/omiga/src/order/shared/repositories/cli/configuration"
	"github.com/omiga-group/omiga/src/order/shared/repositories/migrate"
	entconfiguration "github.com/omiga-group/omiga/src/shared/enterprise/configuration"
	"github.com/spf13/cobra"
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

			var config configuration.Config
			if err := entconfiguration.LoadConfig("config.yaml", &config); err != nil {
				sugarLogger.Fatal(err)
			}

			entgoClient, err := NewEntgoClient(
				sugarLogger,
				config.Postgres)
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
