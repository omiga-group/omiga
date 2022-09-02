package commands

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/omiga-group/omiga/src/exchange/binance-processor/appsetup"
	"github.com/omiga-group/omiga/src/exchange/binance-processor/configuration"
	orderbookv1 "github.com/omiga-group/omiga/src/shared/clients/events/omiga/order-book/v1"
	syntheticorderv1 "github.com/omiga-group/omiga/src/shared/clients/events/omiga/synthetic-order/v1"
	entconfiguration "github.com/omiga-group/omiga/src/shared/enterprise/configuration"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

func startCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start",
		Short: "Start binance-processor",
		Long:  "Start binance-processor",
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

			for _, symbolConfig := range config.Binance.OrderBook.Symbols {
				_, err = appsetup.NewBinanceOrderBookSubscriber(
					ctx,
					sugarLogger,
					config.App,
					config.Binance,
					symbolConfig,
					config.Pulsar,
					config.Postgres,
					orderbookv1.TopicName)
				if err != nil {
					sugarLogger.Fatal(err)
				}
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

	return cmd
}
