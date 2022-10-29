package subscribers

import (
	"context"

	"github.com/huobirdcenter/huobi_golang/pkg/client"
	"github.com/omiga-group/omiga/src/shared/enterprise/cron"
	"github.com/omiga-group/omiga/src/venue/huobi-processor/configuration"
	"github.com/omiga-group/omiga/src/venue/huobi-processor/mappers"
	"github.com/omiga-group/omiga/src/venue/shared/repositories"
	"go.uber.org/zap"
)

type HuobiTradingPairSubscriber interface {
}

type huobiTradingPairSubscriber struct {
	ctx                   context.Context
	logger                *zap.SugaredLogger
	venueConfig           configuration.HuobiConfig
	tradingPairRepository repositories.TradingPairRepository
}

func NewHuobiTradingPairSubscriber(
	ctx context.Context,
	logger *zap.SugaredLogger,
	venueConfig configuration.HuobiConfig,
	cronService cron.CronService,
	tradingPairRepository repositories.TradingPairRepository) (HuobiTradingPairSubscriber, error) {

	instance := &huobiTradingPairSubscriber{
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

func (htps *huobiTradingPairSubscriber) Run() {
	htps.logger.Errorf("Start trading pairs sync for Venue: %s ...", htps.venueConfig.Id)

	client := &client.CommonClient{}

	symbols, err := client.
		Init(htps.venueConfig.BaseUrl).
		GetSymbols()
	if err != nil {
		htps.logger.Errorf("Failed to call common/symbols endpoint. Error: %v", err)

		return
	}

	if err = htps.tradingPairRepository.CreateTradingPairs(
		htps.ctx,
		htps.venueConfig.Id,
		mappers.HuobiSymbolsToTradingPairs(symbols)); err != nil {
		htps.logger.Errorf("Failed to create trading pairs. Error: %v", err)

		return
	}

	htps.logger.Errorf("Finished syncing trading pairs for Venue: %s", htps.venueConfig.Id)
}
