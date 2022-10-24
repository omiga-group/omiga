package commands

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	orderbookv1 "github.com/omiga-group/omiga/src/shared/clients/events/omiga/order-book/v1"
	entconfiguration "github.com/omiga-group/omiga/src/shared/enterprise/configuration"
	"github.com/omiga-group/omiga/src/venue/$VENUE@LOW$-processor/appsetup"
	"github.com/omiga-group/omiga/src/venue/$VENUE@LOW$-processor/configuration"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

func startCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start",
		Short: "Start $VENUE@LOW$-processor",
		Long:  "Start $VENUE@LOW$-processor",
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

			$VENUE@LOW$.UseTestnet = config.$VENUE@PAS$.UseTestnet

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

			for _, pairConfig := range config.$VENUE@PAS$.OrderBook.Pairs {
				$VENUE@LOW$OrderBookSubscriber, err := appsetup.New$VENUE@PAS$OrderBookSubscriber(
					ctx,
					sugarLogger,
					config.App,
					config.$VENUE@PAS$,
					pairConfig,
					config.Pulsar,
					config.Postgres,
					orderbookv1.TopicName)
				if err != nil {
					sugarLogger.Fatal(err)
				}

				defer $VENUE@LOW$OrderBookSubscriber.Close()
			}

			cronService, err := appsetup.NewCronService(sugarLogger)
			if err != nil {
				sugarLogger.Fatal(err)
			}

			defer cronService.Close()

			if _, err = appsetup.New$VENUE@PAS$TradingPairSubscriber(
				ctx,
				sugarLogger,
				config.$VENUE@PAS$,
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
