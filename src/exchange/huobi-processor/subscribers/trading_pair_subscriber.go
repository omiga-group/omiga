package subscribers

import (
	"context"

	"github.com/huobirdcenter/huobi_golang/pkg/client"
	"github.com/omiga-group/omiga/src/exchange/huobi-processor/configuration"
	"github.com/omiga-group/omiga/src/exchange/huobi-processor/mappers"
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
		htps.exchangeConfig.Id,
		mappers.HuobiSymbolsToTradingPairs(symbols)); err != nil {
		htps.logger.Errorf("Failed to create trading pairs. Error: %v", err)

		return
	}
}
