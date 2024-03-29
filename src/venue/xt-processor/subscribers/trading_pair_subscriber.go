package subscribers

import (
	"context"

	"github.com/go-co-op/gocron"
	"github.com/omiga-group/omiga/src/venue/shared/repositories"
	"github.com/omiga-group/omiga/src/venue/xt-processor/configuration"
	"github.com/omiga-group/omiga/src/venue/xt-processor/mappers"
	xtv4 "github.com/omiga-group/omiga/src/venue/xt-processor/xtclient/v4"
	"go.uber.org/zap"
)

type XtTradingPairSubscriber interface {
}

type xtTradingPairSubscriber struct {
	ctx                   context.Context
	logger                *zap.SugaredLogger
	venueConfig           configuration.XtConfig
	tradingPairRepository repositories.TradingPairRepository
}

func NewXtTradingPairSubscriber(
	ctx context.Context,
	logger *zap.SugaredLogger,
	venueConfig configuration.XtConfig,
	jobScheduler *gocron.Scheduler,
	tradingPairRepository repositories.TradingPairRepository) (XtTradingPairSubscriber, error) {

	instance := &xtTradingPairSubscriber{
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

func (xtps *xtTradingPairSubscriber) Run() {
	xtps.logger.Infof("Start trading pairs sync for Venue: %s ...", xtps.venueConfig.Id)

	client, err := xtv4.NewClientWithResponses(xtps.venueConfig.BaseUrl)
	if err != nil {
		xtps.logger.Errorf("Failed to create client with response. Error: %v", err)

		return
	}

	response, err := client.GetAllSymbolsWithResponse(xtps.ctx)
	if err != nil {
		xtps.logger.Errorf("Failed to call getAllMarketConfig endpoint. Error: %v", err)

		return
	}

	if response.HTTPResponse.StatusCode != 200 {
		xtps.logger.Errorf("Failed to call getAllMarketConfig endpoint. Return status code is %d", response.HTTPResponse.StatusCode)

		return
	}

	if response.JSON200 == nil {
		xtps.logger.Errorf("Returned JSON object is nil")

		return
	}

	if err = xtps.tradingPairRepository.CreateTradingPairs(
		xtps.ctx,
		xtps.venueConfig.Id,
		mappers.XtSymbolsToTradingPairs(response.JSON200.Result.Symbols)); err != nil {
		xtps.logger.Errorf("Failed to create trading pairs. Error: %v", err)

		return
	}

	xtps.logger.Infof("Finished syncing trading pairs for Venue: %s", xtps.venueConfig.Id)
}
