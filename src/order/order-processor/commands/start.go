package commands

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/mitchellh/mapstructure"
	orderbookv1 "github.com/omiga-group/omiga/src/shared/clients/events/omiga/order-book/v1"
	orderv1 "github.com/omiga-group/omiga/src/shared/clients/events/omiga/order/v1"
	"github.com/omiga-group/omiga/src/shared/enterprise/configuration"
	"github.com/omiga-group/omiga/src/shared/enterprise/database/postgres"
	"github.com/omiga-group/omiga/src/shared/enterprise/messaging/pulsar"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

func startCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start",
		Short: "Start order-processor",
		Long:  "Start order-processor",
		Run: func(cmd *cobra.Command, args []string) {
			logger, err := zap.NewDevelopment()
			if err != nil {
				log.Fatal(err)
			}

			sugarLogger := logger.Sugar()

			viper, err := configuration.SetupConfigReader(".")
			if err != nil {
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

			ctx, cancelFunc := context.WithCancel(context.Background())

			sigc := make(chan os.Signal, 1)
			signal.Notify(sigc,
				syscall.SIGHUP,
				syscall.SIGINT,
				syscall.SIGTERM,
				syscall.SIGQUIT)
			go func() {
				<-sigc
				cancelFunc()
			}()

			entgoClient, err := NewEntgoClient(
				sugarLogger,
				postgresSettings)
			if err != nil {
				sugarLogger.Fatal(err)
			}

			orderMessageConsumer, err := NewMessageConsumer(
				sugarLogger,
				pulsarSettings,
				orderv1.TopicName)
			if err != nil {
				sugarLogger.Fatal(err)
			}

			defer orderMessageConsumer.Close(ctx)

			orderConsumer, err := NewOrderConsumer(
				sugarLogger,
				orderMessageConsumer)
			if err != nil {
				sugarLogger.Fatal(err)
			}

			err = orderConsumer.StartAsync(ctx)
			if err != nil {
				sugarLogger.Fatal(err)
			}

			orderBookMessageConsumer, err := NewMessageConsumer(
				sugarLogger,
				pulsarSettings,
				orderbookv1.TopicName)
			if err != nil {
				sugarLogger.Fatal(err)
			}

			defer orderMessageConsumer.Close(ctx)

			orderBookConsumer, err := NewOrderBookConsumer(
				sugarLogger,
				orderBookMessageConsumer,
				entgoClient)
			if err != nil {
				sugarLogger.Fatal(err)
			}

			err = orderBookConsumer.StartAsync(ctx)
			if err != nil {
				sugarLogger.Fatal(err)
			}

			for {
				if ctx.Err() == context.Canceled {
					break
				}

				select {
				case <-ctx.Done():
				case <-time.After(time.Second):
				}
			}
		},
	}

	return cmd
}
