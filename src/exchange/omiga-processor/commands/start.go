package commands

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/omiga-group/omiga/src/exchange/omiga-processor/appsetup"
	"github.com/omiga-group/omiga/src/exchange/omiga-processor/configuration"
	"github.com/omiga-group/omiga/src/exchange/omiga-processor/simulators"
	orderbookv1 "github.com/omiga-group/omiga/src/shared/clients/events/omiga/order-book/v1"
	syntheticorderv1 "github.com/omiga-group/omiga/src/shared/clients/events/omiga/synthetic-order/v1"
	entconfiguration "github.com/omiga-group/omiga/src/shared/enterprise/configuration"
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

			var config configuration.Config
			if err := entconfiguration.LoadConfig("config.yaml", &config); err != nil {
				sugarLogger.Fatal(err)
			}

			config.App.Source = config.App.Source + "::" + opt.name

			config.Pulsar.SubscriptionName = config.Pulsar.SubscriptionName + "-" + opt.name
			config.Pulsar.ProducerName = config.Pulsar.ProducerName + opt.name

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

			_, err = appsetup.NewOrderBookSimulator(
				ctx,
				sugarLogger,
				config.App,
				config.Pulsar,
				orderbookv1.TopicName,
				simulators.OrderBookSimulatorConfig{
					ExchangeName: opt.name,
				})
			if err != nil {
				sugarLogger.Fatal(err)
			}

			syntheticMessageConsumer, err := appsetup.NewMessageConsumer(
				sugarLogger,
				config.Pulsar,
				syntheticorderv1.TopicName)
			if err != nil {
				sugarLogger.Fatal(err)
			}

			defer syntheticMessageConsumer.Close(ctx)

			syntheticOrderConsumer, err := appsetup.NewSyntheticOrderConsumer(
				sugarLogger,
				syntheticMessageConsumer)
			if err != nil {
				sugarLogger.Fatal(err)
			}

			err = syntheticOrderConsumer.StartAsync(ctx)
			if err != nil {
				sugarLogger.Fatal(err)
			}

			timeHelper, err := appsetup.NewTimeHelper()
			if err != nil {
				sugarLogger.Fatal(err)
			}

			for {
				if ctx.Err() == context.Canceled {
					break
				}

				timeHelper.SleepOrWaitForContextGetCancelled(
					ctx,
					time.Second)
			}
		},
	}

	cmd.Flags().StringVar(&opt.name, "name", "", "The omiga-processor instance name")

	return cmd
}
