package subscribers

import (
	"context"

	"github.com/omiga-group/omiga/src/shared/enterprise/cron"
	"github.com/omiga-group/omiga/src/venue/ftx-processor/configuration"
	ftxv1 "github.com/omiga-group/omiga/src/venue/ftx-processor/ftxclient/v1"
	"github.com/omiga-group/omiga/src/venue/ftx-processor/mappers"
	"github.com/omiga-group/omiga/src/venue/shared/repositories"
	"go.uber.org/zap"
)

type FtxTradingPairSubscriber interface {
}

type ftxTradingPairSubscriber struct {
	ctx                   context.Context
	logger                *zap.SugaredLogger
	ftxConfig             configuration.FtxConfig
	tradingPairRepository repositories.TradingPairRepository
}

func NewFtxTradingPairSubscriber(
	ctx context.Context,
	logger *zap.SugaredLogger,
	ftxConfig configuration.FtxConfig,
	cronService cron.CronService,
	tradingPairRepository repositories.TradingPairRepository) (FtxTradingPairSubscriber, error) {

	instance := &ftxTradingPairSubscriber{
		ctx:                   ctx,
		logger:                logger,
		ftxConfig:             ftxConfig,
		tradingPairRepository: tradingPairRepository,
	}

	// Run at every minute from 0 through 59.
	if _, err := cronService.GetCron().AddJob("* 0/1 * * * *", instance); err != nil {
		return nil, err
	}

	return instance, nil
}

func (ftps *ftxTradingPairSubscriber) Run() {
	client, err := ftxv1.NewClientWithResponses(ftps.ftxConfig.ApiUrl)
	if err != nil {
		ftps.logger.Errorf("Failed to create client with response. Error: %v", err)

		return
	}

	response, err := client.GetMarketsWithResponse(ftps.ctx)
	if err != nil {
		ftps.logger.Errorf("Failed to call getMarkets endpoint. Error: %v", err)

		return
	}

	if response.HTTPResponse.StatusCode != 200 {
		ftps.logger.Errorf("Failed to call getMarkets endpoint. Return status code is %d", response.HTTPResponse.StatusCode)

		return
	}

	if response.JSON200 == nil {
		ftps.logger.Errorf("Returned JSON object is nil")

		return
	}

	m := *response.JSON200.Result
	if err = ftps.tradingPairRepository.CreateTradingPairs(
		ftps.ctx,
		ftps.ftxConfig.Id,
		mappers.FtxMarketToTradingPairs(m)); err != nil {
		ftps.logger.Errorf("Failed to create trading pairs. Error: %v", err)

		return
	}
}
