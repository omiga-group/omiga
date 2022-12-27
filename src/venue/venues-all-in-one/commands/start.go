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
	binanceprocessorappsetup "github.com/omiga-group/omiga/src/venue/binance-processor/appsetup"
	binanceprocessorconfiguration "github.com/omiga-group/omiga/src/venue/binance-processor/configuration"
	bitmartprocessorappsetup "github.com/omiga-group/omiga/src/venue/bitmart-processor/appsetup"
	bitmartprocessorconfiguration "github.com/omiga-group/omiga/src/venue/bitmart-processor/configuration"
	bittrexprocessorappsetup "github.com/omiga-group/omiga/src/venue/bittrex-processor/appsetup"
	bittrexprocessorconfiguration "github.com/omiga-group/omiga/src/venue/bittrex-processor/configuration"
	bybitprocessorappsetup "github.com/omiga-group/omiga/src/venue/bybit-processor/appsetup"
	bybitprocessorconfiguration "github.com/omiga-group/omiga/src/venue/bybit-processor/configuration"
	coinbaseprocessorappsetup "github.com/omiga-group/omiga/src/venue/coinbase-processor/appsetup"
	coinbaseprocessorconfiguration "github.com/omiga-group/omiga/src/venue/coinbase-processor/configuration"
	cryptoprocessorappsetup "github.com/omiga-group/omiga/src/venue/crypto-processor/appsetup"
	cryptoprocessorconfiguration "github.com/omiga-group/omiga/src/venue/crypto-processor/configuration"
	dextradeprocessorappsetup "github.com/omiga-group/omiga/src/venue/dextrade-processor/appsetup"
	dextradeprocessorconfiguration "github.com/omiga-group/omiga/src/venue/dextrade-processor/configuration"
	ftxprocessorappsetup "github.com/omiga-group/omiga/src/venue/ftx-processor/appsetup"
	ftxprocessorconfiguration "github.com/omiga-group/omiga/src/venue/ftx-processor/configuration"
	geminiprocessorappsetup "github.com/omiga-group/omiga/src/venue/gemini-processor/appsetup"
	geminiprocessorconfiguration "github.com/omiga-group/omiga/src/venue/gemini-processor/configuration"
	huobiprocessorappsetup "github.com/omiga-group/omiga/src/venue/huobi-processor/appsetup"
	huobiprocessorconfiguration "github.com/omiga-group/omiga/src/venue/huobi-processor/configuration"
	krakenprocessorappsetup "github.com/omiga-group/omiga/src/venue/kraken-processor/appsetup"
	krakenprocessorconfiguration "github.com/omiga-group/omiga/src/venue/kraken-processor/configuration"
	kucoinprocessorappsetup "github.com/omiga-group/omiga/src/venue/kucoin-processor/appsetup"
	kucoinprocessorconfiguration "github.com/omiga-group/omiga/src/venue/kucoin-processor/configuration"
	mexcprocessorappsetup "github.com/omiga-group/omiga/src/venue/mexc-processor/appsetup"
	mexcprocessorconfiguration "github.com/omiga-group/omiga/src/venue/mexc-processor/configuration"
	rainprocessorappsetup "github.com/omiga-group/omiga/src/venue/rain-processor/appsetup"
	rainprocessorconfiguration "github.com/omiga-group/omiga/src/venue/rain-processor/configuration"
	xtprocessorappsetup "github.com/omiga-group/omiga/src/venue/xt-processor/appsetup"
	xtprocessorconfiguration "github.com/omiga-group/omiga/src/venue/xt-processor/configuration"
	"github.com/spf13/cobra"
)

func startCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start",
		Short: "Start venues-all-in-one",
		Long:  "Start venues-all-in-one",
		Run: func(cmd *cobra.Command, args []string) {
			sugarLogger := logger.CreateLogger()

			configurationHelper, err := enterpriseappsetup.NewConfigurationHelper(sugarLogger)
			if err != nil {
				sugarLogger.Fatal(err)
			}

			var binanceProcessorConfig binanceprocessorconfiguration.Config
			if err := configurationHelper.LoadYaml("binance-processor-config.yaml", &binanceProcessorConfig); err != nil {
				sugarLogger.Fatal(err)
			}

			var bitmartProcessorConfig bitmartprocessorconfiguration.Config
			if err := configurationHelper.LoadYaml("bitmart-processor-config.yaml", &bitmartProcessorConfig); err != nil {
				sugarLogger.Fatal(err)
			}

			var bittrexProcessorConfig bittrexprocessorconfiguration.Config
			if err := configurationHelper.LoadYaml("bittrex-processor-config.yaml", &bittrexProcessorConfig); err != nil {
				sugarLogger.Fatal(err)
			}

			var bybitProcessorConfig bybitprocessorconfiguration.Config
			if err := configurationHelper.LoadYaml("bybit-processor-config.yaml", &bybitProcessorConfig); err != nil {
				sugarLogger.Fatal(err)
			}

			var coinbaseProcessorConfig coinbaseprocessorconfiguration.Config
			if err := configurationHelper.LoadYaml("coinbase-processor-config.yaml", &coinbaseProcessorConfig); err != nil {
				sugarLogger.Fatal(err)
			}

			var cryptoProcessorConfig cryptoprocessorconfiguration.Config
			if err := configurationHelper.LoadYaml("crypto-processor-config.yaml", &cryptoProcessorConfig); err != nil {
				sugarLogger.Fatal(err)
			}

			var dextradeProcessorConfig dextradeprocessorconfiguration.Config
			if err := configurationHelper.LoadYaml("dextrade-processor-config.yaml", &dextradeProcessorConfig); err != nil {
				sugarLogger.Fatal(err)
			}

			var geminiProcessorConfig geminiprocessorconfiguration.Config
			if err := configurationHelper.LoadYaml("gemini-processor-config.yaml", &geminiProcessorConfig); err != nil {
				sugarLogger.Fatal(err)
			}

			var huobiProcessorConfig huobiprocessorconfiguration.Config
			if err := configurationHelper.LoadYaml("huobi-processor-config.yaml", &huobiProcessorConfig); err != nil {
				sugarLogger.Fatal(err)
			}

			var krakenProcessorConfig krakenprocessorconfiguration.Config
			if err := configurationHelper.LoadYaml("kraken-processor-config.yaml", &krakenProcessorConfig); err != nil {
				sugarLogger.Fatal(err)
			}

			var kucoinProcessorConfig kucoinprocessorconfiguration.Config
			if err := configurationHelper.LoadYaml("kucoin-processor-config.yaml", &kucoinProcessorConfig); err != nil {
				sugarLogger.Fatal(err)
			}

			var mexcProcessorConfig mexcprocessorconfiguration.Config
			if err := configurationHelper.LoadYaml("mexc-processor-config.yaml", &mexcProcessorConfig); err != nil {
				sugarLogger.Fatal(err)
			}

			var rainProcessorConfig rainprocessorconfiguration.Config
			if err := configurationHelper.LoadYaml("rain-processor-config.yaml", &rainProcessorConfig); err != nil {
				sugarLogger.Fatal(err)
			}

			var xtProcessorConfig xtprocessorconfiguration.Config
			if err := configurationHelper.LoadYaml("xt-processor-config.yaml", &xtProcessorConfig); err != nil {
				sugarLogger.Fatal(err)
			}

			var ftxProcessorConfig ftxprocessorconfiguration.Config
			if err := configurationHelper.LoadYaml("ftx-processor-config.yaml", &ftxProcessorConfig); err != nil {
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

			if _, err = binanceprocessorappsetup.NewBinanceTradingPairSubscriber(
				ctx,
				sugarLogger,
				binanceProcessorConfig.Binance,
				jobScheduler,
				binanceProcessorConfig.Postgres); err != nil {
				sugarLogger.Fatal(err)
			}

			if _, err = bitmartprocessorappsetup.NewBitmartTradingPairSubscriber(
				ctx,
				sugarLogger,
				bitmartProcessorConfig.Bitmart,
				jobScheduler,
				bitmartProcessorConfig.Postgres); err != nil {
				sugarLogger.Fatal(err)
			}

			if _, err = bittrexprocessorappsetup.NewBittrexTradingPairSubscriber(
				ctx,
				sugarLogger,
				bittrexProcessorConfig.Bittrex,
				jobScheduler,
				bittrexProcessorConfig.Postgres); err != nil {
				sugarLogger.Fatal(err)
			}

			if _, err = bybitprocessorappsetup.NewBybitTradingPairSubscriber(
				ctx,
				sugarLogger,
				bybitProcessorConfig.Bybit,
				jobScheduler,
				bybitProcessorConfig.Postgres); err != nil {
				sugarLogger.Fatal(err)
			}

			if _, err = coinbaseprocessorappsetup.NewCoinbaseTradingPairSubscriber(
				ctx,
				sugarLogger,
				coinbaseProcessorConfig.Coinbase,
				jobScheduler,
				coinbaseProcessorConfig.Postgres); err != nil {
				sugarLogger.Fatal(err)
			}

			if _, err = cryptoprocessorappsetup.NewCryptoTradingPairSubscriber(
				ctx,
				sugarLogger,
				cryptoProcessorConfig.Crypto,
				jobScheduler,
				cryptoProcessorConfig.Postgres); err != nil {
				sugarLogger.Fatal(err)
			}

			if _, err = dextradeprocessorappsetup.NewDextradeTradingPairSubscriber(
				ctx,
				sugarLogger,
				dextradeProcessorConfig.Dextrade,
				jobScheduler,
				dextradeProcessorConfig.Postgres); err != nil {
				sugarLogger.Fatal(err)
			}

			if _, err = geminiprocessorappsetup.NewGeminiTradingPairSubscriber(
				ctx,
				sugarLogger,
				geminiProcessorConfig.Gemini,
				jobScheduler,
				geminiProcessorConfig.Postgres); err != nil {
				sugarLogger.Fatal(err)
			}

			if _, err = huobiprocessorappsetup.NewHuobiTradingPairSubscriber(
				ctx,
				sugarLogger,
				huobiProcessorConfig.Huobi,
				jobScheduler,
				huobiProcessorConfig.Postgres); err != nil {
				sugarLogger.Fatal(err)
			}

			if _, err = krakenprocessorappsetup.NewKrakenTradingPairSubscriber(
				ctx,
				sugarLogger,
				krakenProcessorConfig.Kraken,
				jobScheduler,
				krakenProcessorConfig.Postgres); err != nil {
				sugarLogger.Fatal(err)
			}

			if _, err = kucoinprocessorappsetup.NewKucoinTradingPairSubscriber(
				ctx,
				sugarLogger,
				kucoinProcessorConfig.Kucoin,
				jobScheduler,
				kucoinProcessorConfig.Postgres); err != nil {
				sugarLogger.Fatal(err)
			}

			if _, err = mexcprocessorappsetup.NewMexcTradingPairSubscriber(
				ctx,
				sugarLogger,
				mexcProcessorConfig.Mexc,
				jobScheduler,
				mexcProcessorConfig.Postgres); err != nil {
				sugarLogger.Fatal(err)
			}

			if _, err = rainprocessorappsetup.NewRainTradingPairSubscriber(
				ctx,
				sugarLogger,
				rainProcessorConfig.Rain,
				jobScheduler,
				rainProcessorConfig.Postgres); err != nil {
				sugarLogger.Fatal(err)
			}

			if _, err = xtprocessorappsetup.NewXtTradingPairSubscriber(
				ctx,
				sugarLogger,
				xtProcessorConfig.Xt,
				jobScheduler,
				xtProcessorConfig.Postgres); err != nil {
				sugarLogger.Fatal(err)
			}

			if _, err = ftxprocessorappsetup.NewFtxTradingPairSubscriber(
				ctx,
				sugarLogger,
				ftxProcessorConfig.Ftx,
				jobScheduler,
				ftxProcessorConfig.Postgres); err != nil {
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
