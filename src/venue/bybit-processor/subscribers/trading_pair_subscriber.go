package subscribers

import (
	"context"

	"github.com/omiga-group/omiga/src/shared/enterprise/cron"
	"github.com/omiga-group/omiga/src/venue/bybit-processor/configuration"
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

	go instance.Run()

	return instance, nil
}

func (htps *bybitTradingPairSubscriber) Run() {
}
