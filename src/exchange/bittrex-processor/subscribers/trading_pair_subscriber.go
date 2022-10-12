package subscribers

import (
	"context"

	"github.com/omiga-group/omiga/src/exchange/bittrex-processor/configuration"
	"github.com/omiga-group/omiga/src/exchange/bittrex-processor/mappers"
	exchangeConfiguration "github.com/omiga-group/omiga/src/exchange/shared/configuration"
	"github.com/omiga-group/omiga/src/exchange/shared/repositories"
	"github.com/omiga-group/omiga/src/shared/enterprise/cron"
	"github.com/toorop/go-bittrex"
	"go.uber.org/zap"
)

type BittrexTradingPairSubscriber interface {
}

type bittrexTradingPairSubscriber struct {
	ctx                   context.Context
	logger                *zap.SugaredLogger
	bittrexConfig         configuration.BittrexConfig
	exchangeConfig        exchangeConfiguration.ExchangeConfig
	tradingPairRepository repositories.TradingPairRepository
}

func NewBittrexTradingPairSubscriber(
	ctx context.Context,
	logger *zap.SugaredLogger,
	bittrexConfig configuration.BittrexConfig,
	exchangeConfig exchangeConfiguration.ExchangeConfig,
	cronService cron.CronService,
	tradingPairRepository repositories.TradingPairRepository) (BittrexTradingPairSubscriber, error) {

	instance := &bittrexTradingPairSubscriber{
		ctx:                   ctx,
		logger:                logger,
		bittrexConfig:         bittrexConfig,
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

func (btps *bittrexTradingPairSubscriber) Run() {
	client := bittrex.New(btps.bittrexConfig.ApiKey, btps.bittrexConfig.SecretKey)

	markets, err := client.GetMarkets()
	if err != nil {
		btps.logger.Errorf("Failed to call markets endpoint. Error: %v", err)

		return
	}

	if err = btps.tradingPairRepository.CreateTradingPairs(
		btps.ctx,
		btps.exchangeConfig.Id,
		mappers.BittrexMarketsToTradingPairs(markets)); err != nil {
		btps.logger.Errorf("Failed to create trading pairs. Error: %v", err)

		return
	}
}
