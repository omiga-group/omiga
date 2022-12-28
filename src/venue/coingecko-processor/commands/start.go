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
	"github.com/omiga-group/omiga/src/venue/coingecko-processor/appsetup"
	"github.com/omiga-group/omiga/src/venue/coingecko-processor/configuration"
	"github.com/spf13/cobra"
)

type startOptions struct {
	name string
}

func startCommand() *cobra.Command {
	opt := startOptions{}
	cmd := &cobra.Command{
		Use:   "start",
		Short: "Start coingecko",
		Long:  "Start coingecko",
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

			jobScheduler := gocron.NewScheduler(time.UTC)
			jobScheduler.StartAsync()
			defer jobScheduler.Stop()

			if _, err = appsetup.NewCoingeckoExchangeSubscriber(
				ctx,
				sugarLogger,
				jobScheduler,
				config.Coingecko,
				config.Exchanges,
				config.Postgres); err != nil {
				sugarLogger.Fatal(err)
			}

			if _, err = appsetup.NewCoingeckoCoinSubscriber(
				ctx,
				sugarLogger,
				jobScheduler,
				config.Coingecko,
				config.Exchanges,
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

	cmd.Flags().StringVar(&opt.name, "name", "", "The coingecko instance name")

	return cmd
}
