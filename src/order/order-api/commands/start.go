package commands

import (
	"context"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/omiga-group/omiga/src/order/order-api/appsetup"
	"github.com/omiga-group/omiga/src/order/order-api/configuration"
	orderappsetup "github.com/omiga-group/omiga/src/order/shared/appsetup"
	enterpriseappsetup "github.com/omiga-group/omiga/src/shared/enterprise/appsetup"
	"github.com/omiga-group/omiga/src/shared/enterprise/logger"
	"github.com/spf13/cobra"
)

func startCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start",
		Short: "Start order-api",
		Long:  "Start order-api",
		Run: func(cmd *cobra.Command, args []string) {
			ctx := context.Background()
			sugarLogger := logger.CreateLogger()

			configurationHelper, err := enterpriseappsetup.NewConfigurationHelper(sugarLogger)
			if err != nil {
				sugarLogger.Fatal(err)
			}

			var config configuration.Config
			if err := configurationHelper.LoadYaml("config.yaml", &config); err != nil {
				sugarLogger.Fatal(err)
			}

			entgoClient, err := orderappsetup.NewEntgoClient(
				sugarLogger,
				config.Postgres)
			if err != nil {
				sugarLogger.Fatal(err)
			}

			jobScheduler := gocron.NewScheduler(time.UTC)
			jobScheduler.StartAsync()
			defer jobScheduler.Stop()

			pulsarClient, err := enterpriseappsetup.NewPulsarClient(
				sugarLogger,
				config.Pulsar)
			if err != nil {
				sugarLogger.Fatal(err)
			}

			defer pulsarClient.Close()

			outboxBackgroundService, err := orderappsetup.NewOutboxBackgroundService(
				ctx,
				sugarLogger,
				pulsarClient,
				config.Pulsar,
				config.Outbox,
				entgoClient,
				jobScheduler)
			if err != nil {
				sugarLogger.Fatal(err)
			}

			defer outboxBackgroundService.Close()

			httpServer, err := appsetup.NewHttpServer(
				sugarLogger,
				config.App,
				entgoClient,
				outboxBackgroundService)
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
