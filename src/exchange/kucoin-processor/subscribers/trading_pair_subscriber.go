package subscribers

import (
	"context"

	"github.com/omiga-group/omiga/src/exchange/kucoin-processor/configuration"
	exchangeConfiguration "github.com/omiga-group/omiga/src/exchange/shared/configuration"
	"github.com/omiga-group/omiga/src/exchange/shared/repositories"
	"github.com/omiga-group/omiga/src/shared/enterprise/cron"
	"go.uber.org/zap"
)

type KuCoinTradingPairSubscriber interface {
}

type kucoinTradingPairSubscriber struct {
	ctx                   context.Context
	logger                *zap.SugaredLogger
	kucoinConfig          configuration.KuCoinConfig
	exchangeConfig        exchangeConfiguration.ExchangeConfig
	tradingPairRepository repositories.TradingPairRepository
}

func NewKuCoinTradingPairSubscriber(
	ctx context.Context,
	logger *zap.SugaredLogger,
	kucoinConfig configuration.KuCoinConfig,
	exchangeConfig exchangeConfiguration.ExchangeConfig,
	cronService cron.CronService,
	tradingPairRepository repositories.TradingPairRepository) (KuCoinTradingPairSubscriber, error) {

	instance := &kucoinTradingPairSubscriber{
		ctx:                   ctx,
		logger:                logger,
		kucoinConfig:          kucoinConfig,
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

func (ctps *kucoinTradingPairSubscriber) Run() {
}
