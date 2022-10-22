package subscribers

import (
	"context"

	"github.com/omiga-group/omiga/src/shared/enterprise/cron"
	"github.com/omiga-group/omiga/src/venue/ftx-processor/configuration"
	"github.com/omiga-group/omiga/src/venue/shared/repositories"
	"go.uber.org/zap"
)

type FTXTradingPairSubscriber interface {
}

type ftxTradingPairSubscriber struct {
	ctx                   context.Context
	logger                *zap.SugaredLogger
	ftxConfig        	  configuration.FtxConfig
	tradingPairRepository repositories.TradingPairRepository
}

func NewFTXTradingPairSubscriber(
	ctx context.Context,
	logger *zap.SugaredLogger,
	ftxConfig configuration.FtxConfig,
	cronService cron.CronService,
	tradingPairRepository repositories.TradingPairRepository) (FTXTradingPairSubscriber, error) {

	instance := &ftxTradingPairSubscriber{
		ctx:                   ctx,
		logger:                logger,
		ftxConfig:        	   ftxConfig,
		tradingPairRepository: tradingPairRepository,
	}

	// Run at every minute from 0 through 59.
	if _, err := cronService.GetCron().AddJob("* 0/1 * * * *", instance); err != nil {
		return nil, err
	}

	go instance.Run()

	return instance, nil
}

func (ftps *ftxTradingPairSubscriber) Run() {
	exchangeInfo, err := ftx.
		NewClient(btps.binanceConfig.ApiKey, btps.binanceConfig.SecretKey).
		NewExchangeInfoService().
		Do(btps.ctx)
	if err != nil {
		btps.logger.Errorf("Failed to call exchangeInfo endpoint. Error: %v", err)

		return
	}

	if err = btps.tradingPairRepository.CreateTradingPairs(
		btps.ctx,
		btps.binanceConfig.Id,
		mappers.BinanceSymbolsToTradingPairs(exchangeInfo.Symbols)); err != nil {
		btps.logger.Errorf("Failed to create trading pairs. Error: %v", err)

		return
	}
}