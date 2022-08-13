package commands

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/mitchellh/mapstructure"
	orderv1 "github.com/omiga-group/omiga/src/shared/clients/events/omiga/order/v1"
	"github.com/omiga-group/omiga/src/shared/enterprise/configuration"
	"github.com/omiga-group/omiga/src/shared/enterprise/database/postgres"
	"github.com/omiga-group/omiga/src/shared/enterprise/messaging/pulsar"
	"github.com/omiga-group/omiga/src/shared/enterprise/outbox"
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

			viper, err := configuration.SetupConfigReader(".")
			if err != nil {
				sugarLogger.Fatal(err)
			}

			var appSettings configuration.AppSettings
			if err := mapstructure.Decode(viper.Get(configuration.AppSettingsConfigKey), &appSettings); err != nil {
				sugarLogger.Fatal(err)
			}

			var postgresSettings postgres.PostgresSettings
			if err := mapstructure.Decode(viper.Get(postgres.ConfigKey), &postgresSettings); err != nil {
				sugarLogger.Fatal(err)
			}

			var pulsarSettings pulsar.PulsarSettings
			if err := mapstructure.Decode(viper.Get(pulsar.ConfigKey), &pulsarSettings); err != nil {
				sugarLogger.Fatal(err)
			}

			pulsarSettings.ProducerName = pulsarSettings.ProducerName + uuid.NewString()

			var outboxSettings outbox.OutboxSettings
			if err := mapstructure.Decode(viper.Get(outbox.ConfigKey), &outboxSettings); err != nil {
				sugarLogger.Fatal(err)
			}

			entgoClient, err := NewEntgoClient(
				sugarLogger,
				postgresSettings)
			if err != nil {
				sugarLogger.Fatal(err)
			}

			cronService, err := NewCronService(
				sugarLogger)
			if err != nil {
				sugarLogger.Fatal(err)
			}

			defer cronService.Close()

			orderOutboxBackgroundService, err := NewOrderOutboxBackgroundService(
				ctx,
				sugarLogger,
				pulsarSettings,
				outboxSettings,
				orderv1.TopicName,
				entgoClient,
				cronService)
			if err != nil {
				sugarLogger.Fatal(err)
			}

			httpServer, err := NewHttpServer(
				sugarLogger,
				appSettings,
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
