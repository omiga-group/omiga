package commands

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/adshao/go-binance/v2"
	"github.com/go-co-op/gocron"
	orderbookv1 "github.com/omiga-group/omiga/src/shared/clients/events/omiga/order-book/v1"
	enterpriseappsetup "github.com/omiga-group/omiga/src/shared/enterprise/appsetup"
	"github.com/omiga-group/omiga/src/shared/enterprise/logger"
	"github.com/omiga-group/omiga/src/venue/binance-processor/appsetup"
	"github.com/omiga-group/omiga/src/venue/binance-processor/configuration"
	"github.com/spf13/cobra"
)

func startCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start",
		Short: "Start binance-processor",
		Long:  "Start binance-processor",
		Run: func(cmd *cobra.Command, args []string) {
			sugarLogger := logger.CreateLogger()

			configurationHelper, err := enterpriseappsetup.NewConfigurationHelper(sugarLogger)
			if err != nil {
				sugarLogger.Fatal(err)
			}

			var config configuration.Config
			if err := configurationHelper.LoadYaml("config.yaml", &config); err != nil {
				sugarLogger.Fatal(err)
			}

			binance.UseTestnet = config.Binance.UseTestnet

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

			pulsarClient, err := enterpriseappsetup.NewPulsarClient(
				sugarLogger,
				config.Pulsar)
			if err != nil {
				sugarLogger.Fatal(err)
			}

			defer pulsarClient.Close()

			syntheticOrderConsumer, err := appsetup.NewSyntheticOrderConsumer(
				sugarLogger,
				pulsarClient,
				config.Pulsar)
			if err != nil {
				sugarLogger.Fatal(err)
			}

			defer syntheticOrderConsumer.Close()

			err = syntheticOrderConsumer.StartAsync(ctx)
			if err != nil {
				sugarLogger.Fatal(err)
			}

			for _, pairConfig := range config.Binance.OrderBook.Pairs {
				binanceOrderBookSubscriber, err := appsetup.NewBinanceOrderBookSubscriber(
					ctx,
					sugarLogger,
					pulsarClient,
					config.App,
					config.Binance,
					pairConfig,
					config.Pulsar,
					config.Postgres,
					orderbookv1.TopicName)
				if err != nil {
					sugarLogger.Fatal(err)
				}

				defer binanceOrderBookSubscriber.Close()
			}

			jobScheduler := gocron.NewScheduler(time.UTC)
			jobScheduler.StartAsync()
			defer jobScheduler.Stop()

			if _, err = appsetup.NewBinanceTradingPairSubscriber(
				ctx,
				sugarLogger,
				config.Binance,
				jobScheduler,
				config.Postgres); err != nil {
				sugarLogger.Fatal(err)
			}

			timeHelper, err := enterpriseappsetup.NewTimeHelper()
			if err != nil {
				sugarLogger.Fatal(err)
			}

			timeHelper.WaitUntilCancelled(ctx)
		},
	}

	return cmd
}
