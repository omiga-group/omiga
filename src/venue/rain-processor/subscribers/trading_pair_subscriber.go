package subscribers

import (
	"context"

	"github.com/omiga-group/omiga/src/shared/enterprise/cron"
	"github.com/omiga-group/omiga/src/venue/rain-processor/configuration"
	"github.com/omiga-group/omiga/src/venue/shared/repositories"
	"go.uber.org/zap"
)

type RainTradingPairSubscriber interface {
}

type rainTradingPairSubscriber struct {
	ctx                   context.Context
	logger                *zap.SugaredLogger
	rainConfig            configuration.RainConfig
	tradingPairRepository repositories.TradingPairRepository
}

func NewRainTradingPairSubscriber(
	ctx context.Context,
	logger *zap.SugaredLogger,
	rainConfig configuration.RainConfig,
	cronService cron.CronService,
	tradingPairRepository repositories.TradingPairRepository) (RainTradingPairSubscriber, error) {

	instance := &rainTradingPairSubscriber{
		ctx:                   ctx,
		logger:                logger,
		rainConfig:            rainConfig,
		tradingPairRepository: tradingPairRepository,
	}

	// Run at every minute from 0 through 59.
	if _, err := cronService.GetCron().AddJob("* 0/1 * * * *", instance); err != nil {
		return nil, err
	}

	go instance.Run()

	return instance, nil
}

func (mtps *rainTradingPairSubscriber) Run() {

}
