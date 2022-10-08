package commands

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/omiga-group/omiga/src/exchange/coingecko-processor/appsetup"
	"github.com/omiga-group/omiga/src/exchange/coingecko-processor/configuration"
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
		Short: "Start coingecko",
		Long:  "Start coingecko",
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

			cronService, err := appsetup.NewCronService(sugarLogger)
			if err != nil {
				sugarLogger.Fatal(err)
			}

			defer cronService.Close()

			if _, err = appsetup.NewCoingeckoExchangeSubscriber(
				ctx,
				sugarLogger,
				cronService,
				config.Coingecko,
				config.Exchanges,
				config.Postgres); err != nil {
				sugarLogger.Fatal(err)
			}

			if _, err = appsetup.NewCoingeckoCoinSubscriber(
				ctx,
				sugarLogger,
				cronService,
				config.Coingecko,
				config.Exchanges,
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

	cmd.Flags().StringVar(&opt.name, "name", "", "The coingecko instance name")

	return cmd
}
