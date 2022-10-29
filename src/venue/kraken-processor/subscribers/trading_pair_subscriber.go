package subscribers

import (
	"context"

	"github.com/aopoltorzhicky/go_kraken/rest"
	"github.com/go-co-op/gocron"
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
	venueConfig           configuration.KrakenConfig
	tradingPairRepository repositories.TradingPairRepository
}

func NewKrakenTradingPairSubscriber(
	ctx context.Context,
	logger *zap.SugaredLogger,
	venueConfig configuration.KrakenConfig,
	jobScheduler *gocron.Scheduler,
	tradingPairRepository repositories.TradingPairRepository) (KrakenTradingPairSubscriber, error) {

	instance := &krakenTradingPairSubscriber{
		ctx:                   ctx,
		logger:                logger,
		venueConfig:           venueConfig,
		tradingPairRepository: tradingPairRepository,
	}

	jobScheduler.Every(5).Minutes().Do(func() {
		instance.Run()
	})

	return instance, nil
}

func (ktps *krakenTradingPairSubscriber) Run() {
	ktps.logger.Infof("Start trading pairs sync for Venue: %s ...", ktps.venueConfig.Id)

	assetPairs, err := rest.
		New(ktps.venueConfig.ApiKey, ktps.venueConfig.SecretKey).
		AssetPairs()
	if err != nil {
		ktps.logger.Errorf("Failed to call assetPairs endpoint. Error: %v", err)

		return
	}

	if err = ktps.tradingPairRepository.CreateTradingPairs(
		ktps.ctx,
		ktps.venueConfig.Id,
		mappers.KrakenAssetPairsToTradingPairs(assetPairs)); err != nil {
		ktps.logger.Errorf("Failed to create trading pairs. Error: %v", err)

		return
	}

	ktps.logger.Infof("Finished syncing trading pairs for Venue: %s", ktps.venueConfig.Id)
}
