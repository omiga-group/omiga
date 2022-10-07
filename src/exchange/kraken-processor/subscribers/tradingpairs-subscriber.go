package subscribers

import (
	"context"

	"github.com/aopoltorzhicky/go_kraken/rest"
	"github.com/omiga-group/omiga/src/exchange/kraken-processor/configuration"
	"github.com/omiga-group/omiga/src/exchange/kraken-processor/mappers"
	exchangeConfiguration "github.com/omiga-group/omiga/src/exchange/shared/configuration"
	"github.com/omiga-group/omiga/src/exchange/shared/repositories"
	"github.com/omiga-group/omiga/src/shared/enterprise/cron"
	"go.uber.org/zap"
)

type KrakenTradingPairsSubscriber interface {
}

type krakenTradingPairsSubscriber struct {
	ctx                   context.Context
	logger                *zap.SugaredLogger
	krakenConfig          configuration.KrakenConfig
	exchangeConfig        exchangeConfiguration.ExchangeConfig
	tradingPairRepository repositories.TradingPairRepository
}

func NewKrakenTradingPairsSubscriber(
	ctx context.Context,
	logger *zap.SugaredLogger,
	krakenConfig configuration.KrakenConfig,
	exchangeConfig exchangeConfiguration.ExchangeConfig,
	cronService cron.CronService,
	tradingPairRepository repositories.TradingPairRepository) (KrakenTradingPairsSubscriber, error) {

	instance := &krakenTradingPairsSubscriber{
		ctx:                   ctx,
		logger:                logger,
		krakenConfig:          krakenConfig,
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

func (ktps *krakenTradingPairsSubscriber) Run() {
	assetPairs, err := rest.
		New(ktps.krakenConfig.ApiKey, ktps.krakenConfig.SecretKey).
		AssetPairs()
	if err != nil {
		ktps.logger.Errorf("Failed to call assetPairs endpoint. Error: %v", err)

		return
	}

	if err = ktps.tradingPairRepository.CreateTradingPairs(
		ktps.ctx,
		ktps.exchangeConfig.Id,
		mappers.KrakenAssetPairsToTradingPairs(assetPairs)); err != nil {
		ktps.logger.Errorf("Failed to create trading pairs. Error: %v", err)

		return
	}
}
