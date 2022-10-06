package subscribers

import (
	"context"

	"github.com/adshao/go-binance/v2"
	"github.com/omiga-group/omiga/src/exchange/binance-processor/configuration"
	"github.com/omiga-group/omiga/src/exchange/binance-processor/mappers"
	exchangeConfiguration "github.com/omiga-group/omiga/src/exchange/shared/configuration"
	"github.com/omiga-group/omiga/src/exchange/shared/repositories"
	"github.com/omiga-group/omiga/src/shared/enterprise/cron"
	"go.uber.org/zap"
)

type BinanceTradingPairsSubscriber interface {
}

type binanceTradingPairsSubscriber struct {
	ctx                    context.Context
	logger                 *zap.SugaredLogger
	binanceConfig          configuration.BinanceConfig
	exchangeConfig         exchangeConfiguration.ExchangeConfig
	tradingPairsRepository repositories.TradingPairsRepository
}

func NewBinanceTradingPairsSubscriber(
	ctx context.Context,
	logger *zap.SugaredLogger,
	binanceConfig configuration.BinanceConfig,
	exchangeConfig exchangeConfiguration.ExchangeConfig,
	cronService cron.CronService,
	tradingPairsRepository repositories.TradingPairsRepository) (BinanceTradingPairsSubscriber, error) {

	instance := &binanceTradingPairsSubscriber{
		ctx:                    ctx,
		logger:                 logger,
		binanceConfig:          binanceConfig,
		exchangeConfig:         exchangeConfig,
		tradingPairsRepository: tradingPairsRepository,
	}

	// Run at minute 0
	if _, err := cronService.GetCron().AddJob("* 0 * * * *", instance); err != nil {
		return nil, err
	}

	go instance.Run()

	return instance, nil
}

func (btps *binanceTradingPairsSubscriber) Run() {
	exchangeInfo, err := binance.
		NewClient(btps.binanceConfig.ApiKey, btps.binanceConfig.SecretKey).
		NewExchangeInfoService().
		Do(btps.ctx)
	if err != nil {
		btps.logger.Errorf("Failed to call exchangeInfo endpoint. Error: %v", err)

		return
	}

	if err = btps.tradingPairsRepository.CreateTradingPairs(
		btps.ctx,
		btps.exchangeConfig.Id,
		mappers.FromBinanceSymbolsToTradingPairs(exchangeInfo.Symbols)); err != nil {
		btps.logger.Errorf("Failed to create trading pairs. Error: %v", err)

		return
	}
}
