package commands

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/adshao/go-binance/v2"
	orderbookv1 "github.com/omiga-group/omiga/src/shared/clients/events/omiga/order-book/v1"
	entconfiguration "github.com/omiga-group/omiga/src/shared/enterprise/configuration"
	"github.com/omiga-group/omiga/src/venue/binance-processor/appsetup"
	"github.com/omiga-group/omiga/src/venue/binance-processor/configuration"
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

			syntheticOrderConsumer, err := appsetup.NewSyntheticOrderConsumer(
				sugarLogger,
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

			cronService, err := appsetup.NewCronService(sugarLogger)
			if err != nil {
				sugarLogger.Fatal(err)
			}

			defer cronService.Close()

			if _, err = appsetup.NewBinanceTradingPairSubscriber(
				ctx,
				sugarLogger,
				config.Binance,
				config.Exchange,
				cronService,
				config.Postgres); err != nil {
				sugarLogger.Fatal(err)
			}

			timeHelper, err := appsetup.NewTimeHelper()
			if err != nil {
				sugarLogger.Fatal(err)
			}

			timeHelper.WaitUntilCancelled(ctx)
		},
	}

	return cmd
}
