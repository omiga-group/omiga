package commands

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/omiga-group/omiga/src/order/order-processor/appsetup"
	"github.com/omiga-group/omiga/src/order/order-processor/configuration"
	orderbookv1 "github.com/omiga-group/omiga/src/shared/clients/events/omiga/order-book/v1"
	orderv1 "github.com/omiga-group/omiga/src/shared/clients/events/omiga/order/v1"
	entconfiguration "github.com/omiga-group/omiga/src/shared/enterprise/configuration"
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

			entgoClient, err := appsetup.NewEntgoClient(
				sugarLogger,
				config.Postgres)
			if err != nil {
				sugarLogger.Fatal(err)
			}

			orderMessageConsumer, err := appsetup.NewMessageConsumer(
				sugarLogger,
				config.Pulsar,
				orderv1.TopicName)
			if err != nil {
				sugarLogger.Fatal(err)
			}

			defer orderMessageConsumer.Close(ctx)

			orderConsumer, err := appsetup.NewOrderConsumer(
				sugarLogger,
				orderMessageConsumer)
			if err != nil {
				sugarLogger.Fatal(err)
			}

			err = orderConsumer.StartAsync(ctx)
			if err != nil {
				sugarLogger.Fatal(err)
			}

			orderBookMessageConsumer, err := appsetup.NewMessageConsumer(
				sugarLogger,
				config.Pulsar,
				orderbookv1.TopicName)
			if err != nil {
				sugarLogger.Fatal(err)
			}

			defer orderMessageConsumer.Close(ctx)

			orderBookConsumer, err := appsetup.NewOrderBookConsumer(
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
