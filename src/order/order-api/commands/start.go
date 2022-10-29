package commands

import (
	"context"
	"log"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/omiga-group/omiga/src/order/order-api/appsetup"
	"github.com/omiga-group/omiga/src/order/order-api/configuration"
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

			jobScheduler := gocron.NewScheduler(time.UTC)
			jobScheduler.StartAsync()
			defer jobScheduler.Stop()

			outboxBackgroundService, err := appsetup.NewOutboxBackgroundService(
				ctx,
				sugarLogger,
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
