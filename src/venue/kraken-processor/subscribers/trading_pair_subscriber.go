package subscribers

import (
	"context"

	"github.com/aopoltorzhicky/go_kraken/rest"
	"github.com/omiga-group/omiga/src/shared/enterprise/cron"
	"github.com/omiga-group/omiga/src/venue/kraken-processor/configuration"
	"github.com/omiga-group/omiga/src/venue/kraken-processor/mappers"
	"github.com/omiga-group/omiga/src/venue/shared/repositories"
	"go.uber.org/zap"
)

type KrakenTradingPairSubscriber interface {
}

type krakenTradingPairSubscriber struct {
	ctx                   context.Context
	logger                *zap.SugaredLogger
	krakenConfig          configuration.KrakenConfig
	tradingPairRepository repositories.TradingPairRepository
}

func NewKrakenTradingPairSubscriber(
	ctx context.Context,
	logger *zap.SugaredLogger,
	krakenConfig configuration.KrakenConfig,
	cronService cron.CronService,
	tradingPairRepository repositories.TradingPairRepository) (KrakenTradingPairSubscriber, error) {

	instance := &krakenTradingPairSubscriber{
		ctx:                   ctx,
		logger:                logger,
		krakenConfig:          krakenConfig,
		tradingPairRepository: tradingPairRepository,
	}

	// Run at every 5th minute from 0 through 59..
	if _, err := cronService.GetCron().AddJob("* 0/5 * * * *", instance); err != nil {
		return nil, err
	}

	return instance, nil
}

func (ktps *krakenTradingPairSubscriber) Run() {
	assetPairs, err := rest.
		New(ktps.krakenConfig.ApiKey, ktps.krakenConfig.SecretKey).
		AssetPairs()
	if err != nil {
		ktps.logger.Errorf("Failed to call assetPairs endpoint. Error: %v", err)

		return
	}

	if err = ktps.tradingPairRepository.CreateTradingPairs(
		ktps.ctx,
		ktps.krakenConfig.Id,
		mappers.KrakenAssetPairsToTradingPairs(assetPairs)); err != nil {
		ktps.logger.Errorf("Failed to create trading pairs. Error: %v", err)

		return
	}
}
