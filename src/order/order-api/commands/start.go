package commands

import (
	"context"
	"log"

	"github.com/omiga-group/omiga/src/order/order-api/appsetup"
	"github.com/omiga-group/omiga/src/order/order-api/configuration"
	orderv1 "github.com/omiga-group/omiga/src/shared/clients/events/omiga/order/v1"
	entconfiguration "github.com/omiga-group/omiga/src/shared/enterprise/configuration"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

func startCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start",
		Short: "Start order-api",
		Long:  "Start order-api",
		Run: func(cmd *cobra.Command, args []string) {
			ctx := context.Background()

			logger, err := zap.NewDevelopment()
			if err != nil {
				log.Fatal(err)
			}

			sugarLogger := logger.Sugar()

			var config configuration.Config
			if err := entconfiguration.LoadConfig("config.yaml", &config); err != nil {
				sugarLogger.Fatal(err)
			}

			entgoClient, err := appsetup.NewEntgoClient(
				sugarLogger,
				config.Postgres)
			if err != nil {
				sugarLogger.Fatal(err)
			}

			cronService, err := appsetup.NewCronService(
				sugarLogger)
			if err != nil {
				sugarLogger.Fatal(err)
			}

			defer cronService.Close()

			orderOutboxBackgroundService, err := appsetup.NewOrderOutboxBackgroundService(
				ctx,
				sugarLogger,
				config.Pulsar,
				config.Outbox,
				orderv1.TopicName,
				entgoClient,
				cronService)
			if err != nil {
				sugarLogger.Fatal(err)
			}

			httpServer, err := appsetup.NewHttpServer(
				sugarLogger,
				config.App,
				entgoClient,
				orderOutboxBackgroundService)
			if err != nil {
				sugarLogger.Fatal(err)
			}

			err = httpServer.ListenAndServe()
			if err != nil {
				sugarLogger.Fatal(err)
			}
		},
	}

	return cmd
}
