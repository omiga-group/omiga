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
	dexTradeConfig        configuration.DextradeConfig
	tradingPairRepository repositories.TradingPairRepository
}

func NewDextradeTradingPairSubscriber(
	ctx context.Context,
	logger *zap.SugaredLogger,
	dexTradeConfig configuration.DextradeConfig,
	cronService cron.CronService,
	tradingPairRepository repositories.TradingPairRepository) (DextradeTradingPairSubscriber, error) {

	instance := &dexTradeTradingPairSubscriber{
		ctx:                   ctx,
		logger:                logger,
		dexTradeConfig:        dexTradeConfig,
		tradingPairRepository: tradingPairRepository,
	}

	// Run at every 5th minute from 0 through 59..
	if _, err := cronService.GetCron().AddJob("* 0/5 * * * *", instance); err != nil {
		return nil, err
	}

	return instance, nil
}

func (dtps *dexTradeTradingPairSubscriber) Run() {
	client, err := dextradev1.NewClientWithResponses(dtps.dexTradeConfig.BaseUrl)
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
		dtps.dexTradeConfig.Id,
		mappers.DextradeSymbolsToTradingPairs(response.JSON200.Data)); err != nil {
		dtps.logger.Errorf("Failed to create trading pairs. Error: %v", err)

		return
	}
}
