package commands

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/mitchellh/mapstructure"
	"github.com/omiga-group/omiga/src/exchange/ftx-processor/configurations"
	syntheticorderv1 "github.com/omiga-group/omiga/src/shared/clients/events/omiga/synthetic-order/v1"
	"github.com/omiga-group/omiga/src/shared/enterprise/configuration"
	"github.com/omiga-group/omiga/src/shared/enterprise/messaging/pulsar"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

func startCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start",
		Short: "Start ftx-processor",
		Long:  "Start ftx-processor",
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

			var ftxSettings configurations.FtxSettings
			if err := mapstructure.Decode(viper.Get(configurations.ConfigKey), &ftxSettings); err != nil {
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

			syntheticMessageConsumer, err := NewMessageConsumer(
				sugarLogger,
				pulsarSettings,
				syntheticorderv1.TopicName)
			if err != nil {
				sugarLogger.Fatal(err)
			}

			defer syntheticMessageConsumer.Close(ctx)

			syntheticOrderConsumer, err := NewSyntheticOrderConsumer(
				sugarLogger,
				syntheticMessageConsumer)
			if err != nil {
				sugarLogger.Fatal(err)
			}

			err = syntheticOrderConsumer.StartAsync(ctx)
			if err != nil {
				sugarLogger.Fatal(err)
			}

			_, err = NewFtxOrderBookSubscriber(
				ctx,
				sugarLogger,
				ftxSettings,
				"ETH-PERP")
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
