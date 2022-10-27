package subscribers

import (
	"context"

	"github.com/omiga-group/omiga/src/shared/enterprise/cron"
	bybitpotv3 "github.com/omiga-group/omiga/src/venue/bybit-processor/bybitclient/spot/v3"
	"github.com/omiga-group/omiga/src/venue/bybit-processor/configuration"
	"github.com/omiga-group/omiga/src/venue/bybit-processor/mappers"
	"github.com/omiga-group/omiga/src/venue/shared/repositories"
	"go.uber.org/zap"
)

type BybitTradingPairSubscriber interface {
}

type bybitTradingPairSubscriber struct {
	ctx                   context.Context
	logger                *zap.SugaredLogger
	bybitConfig           configuration.BybitConfig
	tradingPairRepository repositories.TradingPairRepository
}

func NewBybitTradingPairSubscriber(
	ctx context.Context,
	logger *zap.SugaredLogger,
	bybitConfig configuration.BybitConfig,
	cronService cron.CronService,
	tradingPairRepository repositories.TradingPairRepository) (BybitTradingPairSubscriber, error) {

	instance := &bybitTradingPairSubscriber{
		ctx:                   ctx,
		logger:                logger,
		bybitConfig:           bybitConfig,
		tradingPairRepository: tradingPairRepository,
	}

	// Run at every minute from 0 through 59.
	if _, err := cronService.GetCron().AddJob("* 0/1 * * * *", instance); err != nil {
		return nil, err
	}

	return instance, nil
}

func (btps *bybitTradingPairSubscriber) Run() {
	client, err := bybitpotv3.NewClientWithResponses(btps.bybitConfig.BaseUrl)
	if err != nil {
		btps.logger.Errorf("Failed to create client with response. Error: %v", err)

		return
	}

	response, err := client.GetAllSymbolsWithResponse(btps.ctx)
	if err != nil {
		btps.logger.Errorf("Failed to call getAllSymbols endpoint. Error: %v", err)

		return
	}

	if response.HTTPResponse.StatusCode != 200 {
		btps.logger.Errorf("Failed to call getAllSymbols endpoint. Return status code is %d", response.HTTPResponse.StatusCode)

		return
	}

	if response.JSON200 == nil {
		btps.logger.Errorf("Returned JSON object is nil")

		return
	}

	if err = btps.tradingPairRepository.CreateTradingPairs(
		btps.ctx,
		btps.bybitConfig.Id,
		mappers.BybitSymbolToTradingPairs(response.JSON200.Result.List)); err != nil {
		btps.logger.Errorf("Failed to create trading pairs. Error: %v", err)

		return
	}
}
