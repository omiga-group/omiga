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
	bybitprocessorappsetup "github.com/omiga-group/omiga/src/venue/bybit-processor/appsetup"
	bybitprocessorconfiguration "github.com/omiga-group/omiga/src/venue/bybit-processor/configuration"
	coinbaseprocessorappsetup "github.com/omiga-group/omiga/src/venue/coinbase-processor/appsetup"
	coinbaseprocessorconfiguration "github.com/omiga-group/omiga/src/venue/coinbase-processor/configuration"
	cryptoprocessorappsetup "github.com/omiga-group/omiga/src/venue/crypto-processor/appsetup"
	cryptoprocessorconfiguration "github.com/omiga-group/omiga/src/venue/crypto-processor/configuration"
	dextradeprocessorappsetup "github.com/omiga-group/omiga/src/venue/dextrade-processor/appsetup"
	dextradeprocessorconfiguration "github.com/omiga-group/omiga/src/venue/dextrade-processor/configuration"
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
	ftxprocessorappsetup "github.com/omiga-group/omiga/src/venue/ftx-processor/appsetup"
	ftxprocessorconfiguration "github.com/omiga-group/omiga/src/venue/ftx-processor/configuration"
	"github.com/omiga-group/omiga/src/venue/venues-all-in-one/appsetup"
	xtprocessorappsetup "github.com/omiga-group/omiga/src/venue/xt-processor/appsetup"
	xtprocessorconfiguration "github.com/omiga-group/omiga/src/venue/xt-processor/configuration"
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

			var bybitProcessorConfig bybitprocessorconfiguration.Config
			if err := entconfiguration.LoadConfig("bybit-processor-config.yaml", &bybitProcessorConfig); err != nil {
				sugarLogger.Fatal(err)
			}

			var coinbaseProcessorConfig coinbaseprocessorconfiguration.Config
			if err := entconfiguration.LoadConfig("coinbase-processor-config.yaml", &coinbaseProcessorConfig); err != nil {
				sugarLogger.Fatal(err)
			}

			var cryptoProcessorConfig cryptoprocessorconfiguration.Config
			if err := entconfiguration.LoadConfig("crypto-processor-config.yaml", &cryptoProcessorConfig); err != nil {
				sugarLogger.Fatal(err)
			}

			var dextradeProcessorConfig dextradeprocessorconfiguration.Config
			if err := entconfiguration.LoadConfig("dextrade-processor-config.yaml", &dextradeProcessorConfig); err != nil {
				sugarLogger.Fatal(err)
			}

			var geminiProcessorConfig geminiprocessorconfiguration.Config
			if err := entconfiguration.LoadConfig("gemini-processor-config.yaml", &geminiProcessorConfig); err != nil {
				sugarLogger.Fatal(err)
			}

			var huobiProcessorConfig huobiprocessorconfiguration.Config
			if err := entconfiguration.LoadConfig("huobi-processor-config.yaml", &huobiProcessorConfig); err != nil {
				sugarLogger.Fatal(err)
			}

			var krakenProcessorConfig krakenprocessorconfiguration.Config
			if err := entconfiguration.LoadConfig("kraken-processor-config.yaml", &krakenProcessorConfig); err != nil {
				sugarLogger.Fatal(err)
			}

			var kucoinProcessorConfig kucoinprocessorconfiguration.Config
			if err := entconfiguration.LoadConfig("kucoin-processor-config.yaml", &kucoinProcessorConfig); err != nil {
				sugarLogger.Fatal(err)
			}

			var mexcProcessorConfig mexcprocessorconfiguration.Config
			if err := entconfiguration.LoadConfig("mexc-processor-config.yaml", &mexcProcessorConfig); err != nil {
				sugarLogger.Fatal(err)
			}

			var xtProcessorConfig xtprocessorconfiguration.Config
			if err := entconfiguration.LoadConfig("xt-processor-config.yaml", &xtProcessorConfig); err != nil {
				sugarLogger.Fatal(err)
			}

			var ftxProcessorConfig ftxprocessorconfiguration.Config
			if err := entconfiguration.LoadConfig("ftx-processor-config.yaml", &ftxProcessorConfig); err != nil {
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

			if _, err = bybitprocessorappsetup.NewBybitTradingPairSubscriber(
				ctx,
				sugarLogger,
				bybitProcessorConfig.Bybit,
				cronService,
				bybitProcessorConfig.Postgres); err != nil {
				sugarLogger.Fatal(err)
			}

			if _, err = coinbaseprocessorappsetup.NewCoinbaseTradingPairSubscriber(
				ctx,
				sugarLogger,
				coinbaseProcessorConfig.Coinbase,
				cronService,
				coinbaseProcessorConfig.Postgres); err != nil {
				sugarLogger.Fatal(err)
			}

			if _, err = cryptoprocessorappsetup.NewCryptoTradingPairSubscriber(
				ctx,
				sugarLogger,
				cryptoProcessorConfig.Crypto,
				cronService,
				cryptoProcessorConfig.Postgres); err != nil {
				sugarLogger.Fatal(err)
			}

			if _, err = dextradeprocessorappsetup.NewDexTradeTradingPairSubscriber(
				ctx,
				sugarLogger,
				dextradeProcessorConfig.DexTrade,
				cronService,
				dextradeProcessorConfig.Postgres); err != nil {
				sugarLogger.Fatal(err)
			}

			if _, err = geminiprocessorappsetup.NewGeminiTradingPairSubscriber(
				ctx,
				sugarLogger,
				geminiProcessorConfig.Gemini,
				cronService,
				geminiProcessorConfig.Postgres); err != nil {
				sugarLogger.Fatal(err)
			}

			if _, err = huobiprocessorappsetup.NewHuobiTradingPairSubscriber(
				ctx,
				sugarLogger,
				huobiProcessorConfig.Huobi,
				cronService,
				huobiProcessorConfig.Postgres); err != nil {
				sugarLogger.Fatal(err)
			}

			if _, err = krakenprocessorappsetup.NewKrakenTradingPairSubscriber(
				ctx,
				sugarLogger,
				krakenProcessorConfig.Kraken,
				cronService,
				krakenProcessorConfig.Postgres); err != nil {
				sugarLogger.Fatal(err)
			}

			if _, err = kucoinprocessorappsetup.NewKuCoinTradingPairSubscriber(
				ctx,
				sugarLogger,
				kucoinProcessorConfig.KuCoin,
				cronService,
				kucoinProcessorConfig.Postgres); err != nil {
				sugarLogger.Fatal(err)
			}

			if _, err = mexcprocessorappsetup.NewMexcTradingPairSubscriber(
				ctx,
				sugarLogger,
				mexcProcessorConfig.Mexc,
				cronService,
				mexcProcessorConfig.Postgres); err != nil {
				sugarLogger.Fatal(err)
			}

			if _, err = xtprocessorappsetup.NewXtTradingPairSubscriber(
				ctx,
				sugarLogger,
				xtProcessorConfig.Xt,
				cronService,
				xtProcessorConfig.Postgres); err != nil {
				sugarLogger.Fatal(err)
			}

			if _, err = ftxprocessorappsetup.NewFtxTradingPairSubscriber(
				ctx,
				sugarLogger,
				ftxProcessorConfig.Ftx,
				cronService,
				ftxProcessorConfig.Postgres); err != nil {
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
