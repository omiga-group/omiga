package subscribers

import (
	"context"

	"github.com/omiga-group/omiga/src/shared/enterprise/cron"
	"github.com/omiga-group/omiga/src/venue/mexc-processor/configuration"
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

func (btps *mexcTradingPairSubscriber) Run() {

}
