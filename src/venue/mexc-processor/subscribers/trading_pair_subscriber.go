package subscribers

import (
	"context"

	"github.com/omiga-group/omiga/src/shared/enterprise/cron"
	"github.com/omiga-group/omiga/src/venue/mexc-processor/configuration"
	"github.com/omiga-group/omiga/src/venue/mexc-processor/mappers"
	mexcpotv2 "github.com/omiga-group/omiga/src/venue/mexc-processor/mexcclient/spot/v2"
	"github.com/omiga-group/omiga/src/venue/shared/repositories"
	"go.uber.org/zap"
)

type MexcTradingPairSubscriber interface {
}

type mexcTradingPairSubscriber struct {
	ctx                   context.Context
	logger                *zap.SugaredLogger
	mexcConfig            configuration.MexcConfig
	tradingPairRepository repositories.TradingPairRepository
}

func NewMexcTradingPairSubscriber(
	ctx context.Context,
	logger *zap.SugaredLogger,
	mexcConfig configuration.MexcConfig,
	cronService cron.CronService,
	tradingPairRepository repositories.TradingPairRepository) (MexcTradingPairSubscriber, error) {

	instance := &mexcTradingPairSubscriber{
		ctx:                   ctx,
		logger:                logger,
		mexcConfig:            mexcConfig,
		tradingPairRepository: tradingPairRepository,
	}

	// Run at every minute from 0 through 59.
	if _, err := cronService.GetCron().AddJob("* 0/1 * * * *", instance); err != nil {
		return nil, err
	}

	go instance.Run()

	return instance, nil
}

func (mtps *mexcTradingPairSubscriber) Run() {
	client, err := mexcpotv2.NewClientWithResponses(mtps.mexcConfig.BaseUrl)
	if err != nil {
		mtps.logger.Errorf("Failed to create client with response. Error: %v", err)

		return
	}

	response, err := client.GetAllSymbolsWithResponse(mtps.ctx)
	if err != nil {
		mtps.logger.Errorf("Failed to call getAllSymbols endpoint. Error: %v", err)

		return
	}

	if response.HTTPResponse.StatusCode != 200 {
		mtps.logger.Errorf("Failed to call getAllSymbols endpoint. Return status code is %d", response.HTTPResponse.StatusCode)

		return
	}

	if response.JSON200 == nil {
		mtps.logger.Errorf("Returned JSON object is nil")

		return
	}

	if err = mtps.tradingPairRepository.CreateTradingPairs(
		mtps.ctx,
		mtps.mexcConfig.Id,
		mappers.MexcSymbolsToTradingPairs(response.JSON200.Data)); err != nil {
		mtps.logger.Errorf("Failed to create trading pairs. Error: %v", err)

		return
	}
}
