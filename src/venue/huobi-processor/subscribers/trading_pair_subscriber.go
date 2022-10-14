package subscribers

import (
	"context"

	"github.com/huobirdcenter/huobi_golang/pkg/client"
	"github.com/omiga-group/omiga/src/shared/enterprise/cron"
	"github.com/omiga-group/omiga/src/venue/huobi-processor/configuration"
	"github.com/omiga-group/omiga/src/venue/huobi-processor/mappers"
	venueConfiguration "github.com/omiga-group/omiga/src/venue/shared/configuration"
	"github.com/omiga-group/omiga/src/venue/shared/repositories"
	"go.uber.org/zap"
)

type HuobiTradingPairSubscriber interface {
}

type huobiTradingPairSubscriber struct {
	ctx                   context.Context
	logger                *zap.SugaredLogger
	huobiConfig           configuration.HuobiConfig
	venueConfig           venueConfiguration.VenueConfig
	tradingPairRepository repositories.TradingPairRepository
}

func NewHuobiTradingPairSubscriber(
	ctx context.Context,
	logger *zap.SugaredLogger,
	huobiConfig configuration.HuobiConfig,
	venueConfig venueConfiguration.VenueConfig,
	cronService cron.CronService,
	tradingPairRepository repositories.TradingPairRepository) (HuobiTradingPairSubscriber, error) {

	instance := &huobiTradingPairSubscriber{
		ctx:                   ctx,
		logger:                logger,
		huobiConfig:           huobiConfig,
		venueConfig:           venueConfig,
		tradingPairRepository: tradingPairRepository,
	}

	// Run at every minute from 0 through 59.
	if _, err := cronService.GetCron().AddJob("* 0/1 * * * *", instance); err != nil {
		return nil, err
	}

	go instance.Run()

	return instance, nil
}

func (htps *huobiTradingPairSubscriber) Run() {
	client := new(client.CommonClient).Init(htps.huobiConfig.BaseUrl)

	symbols, err := client.GetSymbols()
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
}
