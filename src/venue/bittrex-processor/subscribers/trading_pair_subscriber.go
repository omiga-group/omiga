package subscribers

import (
	"context"

	"github.com/go-co-op/gocron"
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
	jobScheduler *gocron.Scheduler,
	tradingPairRepository repositories.TradingPairRepository) (BittrexTradingPairSubscriber, error) {

	instance := &bittrexTradingPairSubscriber{
		ctx:                   ctx,
		logger:                logger,
		venueConfig:           venueConfig,
		tradingPairRepository: tradingPairRepository,
	}

	if _, err := jobScheduler.Every(5).Minutes().Do(func() {
		instance.Run()
	}); err != nil {
		return nil, err
	}

	return instance, nil
}

func (btps *bittrexTradingPairSubscriber) Run() {
	btps.logger.Infof("Start trading pairs sync for Venue: %s ...", btps.venueConfig.Id)

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

	btps.logger.Infof("Finished syncing trading pairs for Venue: %s", btps.venueConfig.Id)
}
