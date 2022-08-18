package commands

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/omiga-group/omiga/src/exchange/omiga-processor/simulators"
	orderbookv1 "github.com/omiga-group/omiga/src/shared/clients/events/omiga/order-book/v1"
	syntheticorderv1 "github.com/omiga-group/omiga/src/shared/clients/events/omiga/synthetic-order/v1"
	"github.com/omiga-group/omiga/src/shared/enterprise/configuration"
	"github.com/omiga-group/omiga/src/shared/enterprise/messaging/pulsar"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

type startOptions struct {
	name string
}

func startCommand() *cobra.Command {
	opt := startOptions{}
	cmd := &cobra.Command{
		Use:   "start",
		Short: "Start omiga-processor",
		Long:  "Start omiga-processor",
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

			appSettings := configuration.GetAppSettings(viper)
			appSettings.Source = appSettings.Source + "::" + opt.name

			pulsarSettings := pulsar.GetPulsarSettings(viper)
			pulsarSettings.SubscriptionName = pulsarSettings.SubscriptionName + "-" + opt.name
			pulsarSettings.ProducerName = pulsarSettings.ProducerName + opt.name

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

			_, err = NewOrderBookSimulator(
				ctx,
				sugarLogger,
				appSettings,
				pulsarSettings,
				orderbookv1.TopicName,
				simulators.OrderBookSimulatorSettings{
					ExchangeName: opt.name,
				})
			if err != nil {
				sugarLogger.Fatal(err)
			}

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

	cmd.Flags().StringVar(&opt.name, "name", "", "The omiga-processor instance name")

	return cmd
}
