package subscribers

import (
	"context"

	"github.com/omiga-group/omiga/src/shared/enterprise/cron"
	"github.com/omiga-group/omiga/src/venue/bittrex-processor/configuration"
	"github.com/omiga-group/omiga/src/venue/bittrex-processor/mappers"
	"github.com/omiga-group/omiga/src/venue/shared/repositories"
	"github.com/toorop/go-bittrex"
	"go.uber.org/zap"
)

type BittrexTradingPairSubscriber interface {
}

type bittrexTradingPairSubscriber struct {
	ctx                   context.Context
	logger                *zap.SugaredLogger
	venueConfig           configuration.BittrexConfig
	tradingPairRepository repositories.TradingPairRepository
}

func NewBittrexTradingPairSubscriber(
	ctx context.Context,
	logger *zap.SugaredLogger,
	venueConfig configuration.BittrexConfig,
	cronService cron.CronService,
	tradingPairRepository repositories.TradingPairRepository) (BittrexTradingPairSubscriber, error) {

	instance := &bittrexTradingPairSubscriber{
		ctx:                   ctx,
		logger:                logger,
		venueConfig:           venueConfig,
		tradingPairRepository: tradingPairRepository,
	}

	// Run at every 5th minute from 0 through 59..
	if _, err := cronService.GetCron().AddJob("* 0/5 * * * *", instance); err != nil {
		return nil, err
	}

	return instance, nil
}

func (btps *bittrexTradingPairSubscriber) Run() {
	btps.logger.Errorf("Start trading pairs sync for Venue: %s ...", btps.venueConfig.Id)

	markets, err := bittrex.
		New(btps.venueConfig.ApiKey, btps.venueConfig.SecretKey).
		GetMarkets()
	if err != nil {
		btps.logger.Errorf("Failed to call markets endpoint. Error: %v", err)

		return
	}

	if err = btps.tradingPairRepository.CreateTradingPairs(
		btps.ctx,
		btps.venueConfig.Id,
		mappers.BittrexMarketsToTradingPairs(markets)); err != nil {
		btps.logger.Errorf("Failed to create trading pairs. Error: %v", err)

		return
	}

	btps.logger.Errorf("Finished syncing trading pairs for Venue: %s", btps.venueConfig.Id)
}
