package subscribers

import (
	"context"

	"github.com/omiga-group/omiga/src/shared/enterprise/cron"
	"github.com/omiga-group/omiga/src/venue/dextrade-processor/configuration"
	"github.com/omiga-group/omiga/src/venue/shared/repositories"
	"go.uber.org/zap"
)

type DexTradeTradingPairSubscriber interface {
}

type dexTradeTradingPairSubscriber struct {
	ctx                   context.Context
	logger                *zap.SugaredLogger
	dexTradeConfig        configuration.DexTradeConfig
	tradingPairRepository repositories.TradingPairRepository
}

func NewDexTradeTradingPairSubscriber(
	ctx context.Context,
	logger *zap.SugaredLogger,
	dexTradeConfig configuration.DexTradeConfig,
	cronService cron.CronService,
	tradingPairRepository repositories.TradingPairRepository) (DexTradeTradingPairSubscriber, error) {

	instance := &dexTradeTradingPairSubscriber{
		ctx:                   ctx,
		logger:                logger,
		dexTradeConfig:        dexTradeConfig,
		tradingPairRepository: tradingPairRepository,
	}

	// Run at every minute from 0 through 59.
	if _, err := cronService.GetCron().AddJob("* 0/1 * * * *", instance); err != nil {
		return nil, err
	}

	go instance.Run()

	return instance, nil
}

func (dtps *dexTradeTradingPairSubscriber) Run() {

}
