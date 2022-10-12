package subscribers

import (
	"context"

	"github.com/omiga-group/omiga/src/exchange/coinbase-processor/configuration"
	"github.com/omiga-group/omiga/src/exchange/coinbase-processor/mappers"
	exchangeConfiguration "github.com/omiga-group/omiga/src/exchange/shared/configuration"
	"github.com/omiga-group/omiga/src/exchange/shared/repositories"
	"github.com/omiga-group/omiga/src/shared/enterprise/cron"
	"github.com/preichenberger/go-coinbasepro/v2"
	"go.uber.org/zap"
)

type CoinbaseTradingPairSubscriber interface {
}

type coinbaseTradingPairSubscriber struct {
	ctx                   context.Context
	logger                *zap.SugaredLogger
	coinbaseConfig        configuration.CoinbaseConfig
	exchangeConfig        exchangeConfiguration.ExchangeConfig
	tradingPairRepository repositories.TradingPairRepository
}

func NewCoinbaseTradingPairSubscriber(
	ctx context.Context,
	logger *zap.SugaredLogger,
	coinbaseConfig configuration.CoinbaseConfig,
	exchangeConfig exchangeConfiguration.ExchangeConfig,
	cronService cron.CronService,
	tradingPairRepository repositories.TradingPairRepository) (CoinbaseTradingPairSubscriber, error) {

	instance := &coinbaseTradingPairSubscriber{
		ctx:                   ctx,
		logger:                logger,
		coinbaseConfig:        coinbaseConfig,
		exchangeConfig:        exchangeConfig,
		tradingPairRepository: tradingPairRepository,
	}

	// Run at minute 0
	if _, err := cronService.GetCron().AddJob("* 0 * * * *", instance); err != nil {
		return nil, err
	}

	go instance.Run()

	return instance, nil
}

func (ctps *coinbaseTradingPairSubscriber) Run() {
	client := coinbasepro.NewClient()

	client.UpdateConfig(&coinbasepro.ClientConfig{
		BaseURL:    ctps.coinbaseConfig.BaseUrl,
		Key:        ctps.coinbaseConfig.ApiKey,
		Passphrase: ctps.coinbaseConfig.Passphrase,
		Secret:     ctps.coinbaseConfig.SecretKey,
	})

	products, err := client.GetProducts()
	if err != nil {
		ctps.logger.Errorf("Failed to call getProducts endpoint. Error: %v", err)

		return
	}

	if err = ctps.tradingPairRepository.CreateTradingPairs(
		ctps.ctx,
		ctps.exchangeConfig.Id,
		mappers.BinanceSymbolsToTradingPairs(products)); err != nil {
		ctps.logger.Errorf("Failed to create trading pairs. Error: %v", err)

		return
	}
}
