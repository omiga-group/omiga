package subscribers

import (
	"context"

	"github.com/omiga-group/omiga/src/shared/enterprise/cron"
	"github.com/omiga-group/omiga/src/venue/dextrade-processor/configuration"
	dextradev1 "github.com/omiga-group/omiga/src/venue/dextrade-processor/dextradeclient/v1"
	"github.com/omiga-group/omiga/src/venue/dextrade-processor/mappers"
	"github.com/omiga-group/omiga/src/venue/shared/repositories"
	"go.uber.org/zap"
)

type DextradeTradingPairSubscriber interface {
}

type dexTradeTradingPairSubscriber struct {
	ctx                   context.Context
	logger                *zap.SugaredLogger
	venueConfig           configuration.DextradeConfig
	tradingPairRepository repositories.TradingPairRepository
}

func NewDextradeTradingPairSubscriber(
	ctx context.Context,
	logger *zap.SugaredLogger,
	venueConfig configuration.DextradeConfig,
	cronService cron.CronService,
	tradingPairRepository repositories.TradingPairRepository) (DextradeTradingPairSubscriber, error) {

	instance := &dexTradeTradingPairSubscriber{
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

func (dtps *dexTradeTradingPairSubscriber) Run() {
	dtps.logger.Errorf("Start trading pairs sync for Venue: %s ...", dtps.venueConfig.Id)

	client, err := dextradev1.NewClientWithResponses(dtps.venueConfig.BaseUrl)
	if err != nil {
		dtps.logger.Errorf("Failed to create client with response. Error: %v", err)

		return
	}

	response, err := client.GetAllSymbolsWithResponse(dtps.ctx)
	if err != nil {
		dtps.logger.Errorf("Failed to call getAllSymbols endpoint. Error: %v", err)

		return
	}

	if response.HTTPResponse.StatusCode != 200 {
		dtps.logger.Errorf("Failed to call getAllSymbols endpoint. Return status code is %d", response.HTTPResponse.StatusCode)

		return
	}

	if response.JSON200 == nil {
		dtps.logger.Errorf("Returned JSON object is nil")

		return
	}

	if err = dtps.tradingPairRepository.CreateTradingPairs(
		dtps.ctx,
		dtps.venueConfig.Id,
		mappers.DextradeSymbolsToTradingPairs(response.JSON200.Data)); err != nil {
		dtps.logger.Errorf("Failed to create trading pairs. Error: %v", err)

		return
	}

	dtps.logger.Errorf("Finished syncing trading pairs for Venue: %s", dtps.venueConfig.Id)
}
