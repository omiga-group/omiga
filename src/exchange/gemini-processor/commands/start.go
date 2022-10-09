package commands

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/omiga-group/omiga/src/exchange/gemini-processor/appsetup"
	"github.com/omiga-group/omiga/src/exchange/gemini-processor/configuration"
	orderbookv1 "github.com/omiga-group/omiga/src/shared/clients/events/omiga/order-book/v1"
	entconfiguration "github.com/omiga-group/omiga/src/shared/enterprise/configuration"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

func startCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start",
		Short: "Start gemini-processor",
		Long:  "Start gemini-processor",
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

			geminiOrderBookSubscriber, err := appsetup.NewGeminiOrderBookSubscriber(
				ctx,
				sugarLogger,
				config.App,
				config.Gemini,
				config.Pulsar,
				orderbookv1.TopicName,
			)
			if err != nil {
				sugarLogger.Fatal(err)
			}
			defer geminiOrderBookSubscriber.Close()

			cronService, err := appsetup.NewCronService(sugarLogger)
			if err != nil {
				sugarLogger.Fatal(err)
			}
			defer cronService.Close()

			if _, err = appsetup.NewGeminiTradingPairsSubscriber(
				ctx,
				sugarLogger,
				config.Gemini,
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
