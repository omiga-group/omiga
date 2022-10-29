package subscribers

import (
	"context"

	"github.com/go-co-op/gocron"
	"github.com/huobirdcenter/huobi_golang/pkg/client"
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
	jobScheduler *gocron.Scheduler,
	tradingPairRepository repositories.TradingPairRepository) (HuobiTradingPairSubscriber, error) {

	instance := &huobiTradingPairSubscriber{
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

func (htps *huobiTradingPairSubscriber) Run() {
	htps.logger.Infof("Start trading pairs sync for Venue: %s ...", htps.venueConfig.Id)

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

	htps.logger.Infof("Finished syncing trading pairs for Venue: %s", htps.venueConfig.Id)
}
