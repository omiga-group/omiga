package subscribers

import (
	"context"

	"github.com/adshao/go-binance/v2"
	"github.com/omiga-group/omiga/src/shared/enterprise/cron"
	"github.com/omiga-group/omiga/src/venue/binance-processor/configuration"
	"github.com/omiga-group/omiga/src/venue/binance-processor/mappers"
	exchangeConfiguration "github.com/omiga-group/omiga/src/venue/shared/configuration"
	"github.com/omiga-group/omiga/src/venue/shared/repositories"
	"go.uber.org/zap"
)

type BinanceTradingPairSubscriber interface {
}

type binanceTradingPairSubscriber struct {
	ctx                   context.Context
	logger                *zap.SugaredLogger
	binanceConfig         configuration.BinanceConfig
	exchangeConfig        exchangeConfiguration.ExchangeConfig
	tradingPairRepository repositories.TradingPairRepository
}

func NewBinanceTradingPairSubscriber(
	ctx context.Context,
	logger *zap.SugaredLogger,
	binanceConfig configuration.BinanceConfig,
	exchangeConfig exchangeConfiguration.ExchangeConfig,
	cronService cron.CronService,
	tradingPairRepository repositories.TradingPairRepository) (BinanceTradingPairSubscriber, error) {

	instance := &binanceTradingPairSubscriber{
		ctx:                   ctx,
		logger:                logger,
		binanceConfig:         binanceConfig,
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

func (btps *binanceTradingPairSubscriber) Run() {
	exchangeInfo, err := binance.
		NewClient(btps.binanceConfig.ApiKey, btps.binanceConfig.SecretKey).
		NewExchangeInfoService().
		Do(btps.ctx)
	if err != nil {
		btps.logger.Errorf("Failed to call exchangeInfo endpoint. Error: %v", err)

		return
	}

	if err = btps.tradingPairRepository.CreateTradingPairs(
		btps.ctx,
		btps.exchangeConfig.Id,
		mappers.BinanceSymbolsToTradingPairs(exchangeInfo.Symbols)); err != nil {
		btps.logger.Errorf("Failed to create trading pairs. Error: %v", err)

		return
	}
}
