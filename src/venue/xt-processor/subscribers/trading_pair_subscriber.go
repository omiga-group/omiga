package subscribers

import (
	"context"

	"github.com/omiga-group/omiga/src/shared/enterprise/cron"
	"github.com/omiga-group/omiga/src/venue/shared/repositories"
	"github.com/omiga-group/omiga/src/venue/xt-processor/configuration"
	"github.com/omiga-group/omiga/src/venue/xt-processor/mappers"
	xtv1 "github.com/omiga-group/omiga/src/venue/xt-processor/xtclient/v1"
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
	cronService cron.CronService,
	tradingPairRepository repositories.TradingPairRepository) (XtTradingPairSubscriber, error) {

	instance := &xtTradingPairSubscriber{
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

func (mtps *xtTradingPairSubscriber) Run() {
	client, err := xtv1.NewClientWithResponses(mtps.venueConfig.BaseUrl)
	if err != nil {
		mtps.logger.Errorf("Failed to create client with response. Error: %v", err)

		return
	}

	response, err := client.GetAllMarketConfigWithResponse(mtps.ctx)
	if err != nil {
		mtps.logger.Errorf("Failed to call getAllMarketConfig endpoint. Error: %v", err)

		return
	}

	if response.HTTPResponse.StatusCode != 200 {
		mtps.logger.Errorf("Failed to call getAllMarketConfig endpoint. Return status code is %d", response.HTTPResponse.StatusCode)

		return
	}

	if response.JSON200 == nil {
		mtps.logger.Errorf("Returned JSON object is nil")

		return
	}

	if err = mtps.tradingPairRepository.CreateTradingPairs(
		mtps.ctx,
		mtps.venueConfig.Id,
		mappers.XtMarketConfigsToTradingPairs(*response.JSON200)); err != nil {
		mtps.logger.Errorf("Failed to create trading pairs. Error: %v", err)

		return
	}
}
