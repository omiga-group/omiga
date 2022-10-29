package subscribers

import (
	"context"

	"github.com/go-co-op/gocron"
	"github.com/omiga-group/omiga/src/venue/$VENUE@LOW$-processor/configuration"
	"github.com/omiga-group/omiga/src/venue/$VENUE@LOW$-processor/mappers"
	"github.com/omiga-group/omiga/src/venue/shared/repositories"
	"go.uber.org/zap"
)

type $VENUE@PAS$TradingPairSubscriber interface {
}

type $VENUE@LOW$TradingPairSubscriber struct {
	ctx                   context.Context
	logger                *zap.SugaredLogger
	venueConfig         configuration.$VENUE@PAS$Config
	tradingPairRepository repositories.TradingPairRepository
}

func New$VENUE@PAS$TradingPairSubscriber(
	ctx context.Context,
	logger *zap.SugaredLogger,
	venueConfig configuration.$VENUE@PAS$Config,
	jobScheduler *gocron.Scheduler,
	tradingPairRepository repositories.TradingPairRepository) ($VENUE@PAS$TradingPairSubscriber, error) {

	instance := &$VENUE@LOW$TradingPairSubscriber{
		ctx:                   ctx,
		logger:                logger,
		venueConfig:         venueConfig,
		tradingPairRepository: tradingPairRepository,
	}

	if _, err := jobScheduler.Every(5).Minutes().Do(func() {
		instance.Run()
	}); err != nil {
		return nil, err
	}

	return instance, nil
}

func (btps *$VENUE@LOW$TradingPairSubscriber) Run() {
	btps.logger.Infof("Start trading pairs sync for Venue: %s ...", btps.venueConfig.Id)

	exchangeInfo, err := $VENUE@LOW$.
		NewClient(btps.venueConfig.ApiKey, btps.venueConfig.SecretKey).
		NewExchangeInfoService().
		Do(btps.ctx)
	if err != nil {
		btps.logger.Errorf("Failed to call exchangeInfo endpoint. Error: %v", err)

		return
	}

	if err = btps.tradingPairRepository.CreateTradingPairs(
		btps.ctx,
		btps.venueConfig.Id,
		mappers.$VENUE@PAS$SymbolsToTradingPairs(exchangeInfo.Symbols)); err != nil {
		btps.logger.Errorf("Failed to create trading pairs. Error: %v", err)

		return
	}
}
