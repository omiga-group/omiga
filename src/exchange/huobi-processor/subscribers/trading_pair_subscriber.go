package subscribers

import (
	"context"

	"github.com/omiga-group/omiga/src/exchange/huobi-processor/configuration"
	exchangeConfiguration "github.com/omiga-group/omiga/src/exchange/shared/configuration"
	"github.com/omiga-group/omiga/src/exchange/shared/repositories"
	"github.com/omiga-group/omiga/src/shared/enterprise/cron"
	"go.uber.org/zap"
)

type HuobiTradingPairSubscriber interface {
}

type huobiTradingPairSubscriber struct {
	ctx                   context.Context
	logger                *zap.SugaredLogger
	huobiConfig           configuration.HuobiConfig
	exchangeConfig        exchangeConfiguration.ExchangeConfig
	tradingPairRepository repositories.TradingPairRepository
}

func NewHuobiTradingPairSubscriber(
	ctx context.Context,
	logger *zap.SugaredLogger,
	huobiConfig configuration.HuobiConfig,
	exchangeConfig exchangeConfiguration.ExchangeConfig,
	cronService cron.CronService,
	tradingPairRepository repositories.TradingPairRepository) (HuobiTradingPairSubscriber, error) {

	instance := &huobiTradingPairSubscriber{
		ctx:                   ctx,
		logger:                logger,
		huobiConfig:           huobiConfig,
		exchangeConfig:        exchangeConfig,
		tradingPairRepository: tradingPairRepository,
	}

	// Run at minute 0
	if _, err := cronService.GetCron().AddJob("* 0 * * * *", instance); err != nil {
		return nil, err
	}

	go instance.Run()

	return instance, nil
}

func (ctps *huobiTradingPairSubscriber) Run() {
}
