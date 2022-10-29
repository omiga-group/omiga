package subscribers

import (
	"context"

	"github.com/adshao/go-binance/v2"
	"github.com/go-co-op/gocron"
	"github.com/omiga-group/omiga/src/venue/binance-processor/configuration"
	"github.com/omiga-group/omiga/src/venue/binance-processor/mappers"
	"github.com/omiga-group/omiga/src/venue/shared/repositories"
	"go.uber.org/zap"
)

type BinanceTradingPairSubscriber interface {
}

type binanceTradingPairSubscriber struct {
	ctx                   context.Context
	logger                *zap.SugaredLogger
	venueConfig           configuration.BinanceConfig
	tradingPairRepository repositories.TradingPairRepository
}

func NewBinanceTradingPairSubscriber(
	ctx context.Context,
	logger *zap.SugaredLogger,
	venueConfig configuration.BinanceConfig,
	jobScheduler *gocron.Scheduler,
	tradingPairRepository repositories.TradingPairRepository) (BinanceTradingPairSubscriber, error) {

	instance := &binanceTradingPairSubscriber{
		ctx:                   ctx,
		logger:                logger,
		venueConfig:           venueConfig,
		tradingPairRepository: tradingPairRepository,
	}

	if _, err := jobScheduler.Every(5).Minutes().Do(func() {
		instance.Run()
	}); err != nil {
		return nil, err
	}

	return instance, nil
}

func (btps *binanceTradingPairSubscriber) Run() {
	btps.logger.Infof("Start trading pairs sync for Venue: %s ...", btps.venueConfig.Id)

	exchangeInfo, err := binance.
		NewClient(btps.venueConfig.ApiKey, btps.venueConfig.SecretKey).
		NewExchangeInfoService().
		Do(btps.ctx)
	if err != nil {
		btps.logger.Errorf("Failed to call exchangeInfo endpoint. Error: %v", err)

		return
	}

	if err = btps.tradingPairRepository.CreateTradingPairs(
		btps.ctx,
		btps.venueConfig.Id,
		mappers.BinanceSymbolsToTradingPairs(exchangeInfo.Symbols)); err != nil {
		btps.logger.Errorf("Failed to create trading pairs. Error: %v", err)

		return
	}

	btps.logger.Infof("Finished syncing trading pairs for Venue: %s", btps.venueConfig.Id)
}
