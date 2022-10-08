package commands

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/omiga-group/omiga/src/order/order-processor/appsetup"
	"github.com/omiga-group/omiga/src/order/order-processor/configuration"
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

			orderConsumer, err := appsetup.NewOrderConsumer(
				sugarLogger,
				config.Pulsar)
			if err != nil {
				sugarLogger.Fatal(err)
			}

			defer orderConsumer.Close()

			err = orderConsumer.StartAsync(ctx)
			if err != nil {
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
