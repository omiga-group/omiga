package subscribers

import (
	"context"

	"github.com/omiga-group/omiga/src/exchange/coinbase-processor/configuration"
	exchangeConfiguration "github.com/omiga-group/omiga/src/exchange/shared/configuration"
	"github.com/omiga-group/omiga/src/exchange/shared/repositories"
	"github.com/omiga-group/omiga/src/shared/enterprise/cron"
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

func NewCoinbaseTradingPairsSubscriber(
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
}
