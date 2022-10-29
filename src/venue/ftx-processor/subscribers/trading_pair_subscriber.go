package subscribers

import (
	"context"

	"github.com/go-co-op/gocron"
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
	venueConfig           configuration.FtxConfig
	tradingPairRepository repositories.TradingPairRepository
}

func NewFtxTradingPairSubscriber(
	ctx context.Context,
	logger *zap.SugaredLogger,
	venueConfig configuration.FtxConfig,
	jobScheduler *gocron.Scheduler,
	tradingPairRepository repositories.TradingPairRepository) (FtxTradingPairSubscriber, error) {

	instance := &ftxTradingPairSubscriber{
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

func (ftps *ftxTradingPairSubscriber) Run() {
	ftps.logger.Infof("Start trading pairs sync for Venue: %s ...", ftps.venueConfig.Id)

	client, err := ftxv1.NewClientWithResponses(ftps.venueConfig.ApiUrl)
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

	if err = ftps.tradingPairRepository.CreateTradingPairs(
		ftps.ctx,
		ftps.venueConfig.Id,
		mappers.FtxMarketToTradingPairs(*response.JSON200.Result)); err != nil {
		ftps.logger.Errorf("Failed to create trading pairs. Error: %v", err)

		return
	}

	ftps.logger.Infof("Finished syncing trading pairs for Venue: %s", ftps.venueConfig.Id)
}
