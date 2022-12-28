package commands

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-co-op/gocron"
	enterpriseappsetup "github.com/omiga-group/omiga/src/shared/enterprise/appsetup"
	"github.com/omiga-group/omiga/src/shared/enterprise/logger"
	"github.com/omiga-group/omiga/src/venue/kucoin-processor/appsetup"
	"github.com/omiga-group/omiga/src/venue/kucoin-processor/configuration"
	"github.com/spf13/cobra"
)

func startCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start",
		Short: "Start kucoin-processor",
		Long:  "Start kucoin-processor",
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

			jobScheduler := gocron.NewScheduler(time.UTC)
			jobScheduler.StartAsync()
			defer jobScheduler.Stop()

			if _, err = appsetup.NewKucoinTradingPairSubscriber(
				ctx,
				sugarLogger,
				config.Kucoin,
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
