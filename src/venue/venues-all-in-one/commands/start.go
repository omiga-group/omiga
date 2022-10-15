package commands

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	entconfiguration "github.com/omiga-group/omiga/src/shared/enterprise/configuration"
	binanceprocessorappsetup "github.com/omiga-group/omiga/src/venue/binance-processor/appsetup"
	binanceprocessorconfiguration "github.com/omiga-group/omiga/src/venue/binance-processor/configuration"
	bitmartprocessorappsetup "github.com/omiga-group/omiga/src/venue/bitmart-processor/appsetup"
	bitmartprocessorconfiguration "github.com/omiga-group/omiga/src/venue/bitmart-processor/configuration"
	bittrexprocessorappsetup "github.com/omiga-group/omiga/src/venue/bittrex-processor/appsetup"
	bittrexprocessorconfiguration "github.com/omiga-group/omiga/src/venue/bittrex-processor/configuration"
	"github.com/omiga-group/omiga/src/venue/venues-all-in-one/appsetup"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

func startCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start",
		Short: "Start venues-all-in-one",
		Long:  "Start venues-all-in-one",
		Run: func(cmd *cobra.Command, args []string) {
			logger, err := zap.NewDevelopment()
			if err != nil {
				log.Fatal(err)
			}

			sugarLogger := logger.Sugar()

			var binanceProcessorConfig binanceprocessorconfiguration.Config
			if err := entconfiguration.LoadConfig("binance-processor-config.yaml", &binanceProcessorConfig); err != nil {
				sugarLogger.Fatal(err)
			}

			var bitmartProcessorConfig bitmartprocessorconfiguration.Config
			if err := entconfiguration.LoadConfig("bitmart-processor-config.yaml", &bitmartProcessorConfig); err != nil {
				sugarLogger.Fatal(err)
			}

			var bittrexProcessorConfig bittrexprocessorconfiguration.Config
			if err := entconfiguration.LoadConfig("bittrex-processor-config.yaml", &bittrexProcessorConfig); err != nil {
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

			if _, err = binanceprocessorappsetup.NewBinanceTradingPairSubscriber(
				ctx,
				sugarLogger,
				binanceProcessorConfig.Binance,
				cronService,
				binanceProcessorConfig.Postgres); err != nil {
				sugarLogger.Fatal(err)
			}

			if _, err = bitmartprocessorappsetup.NewBitMartTradingPairSubscriber(
				ctx,
				sugarLogger,
				bitmartProcessorConfig.BitMart,
				cronService,
				bitmartProcessorConfig.Postgres); err != nil {
				sugarLogger.Fatal(err)
			}

			if _, err = bittrexprocessorappsetup.NewBittrexTradingPairSubscriber(
				ctx,
				sugarLogger,
				bittrexProcessorConfig.Bittrex,
				cronService,
				bittrexProcessorConfig.Postgres); err != nil {
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
