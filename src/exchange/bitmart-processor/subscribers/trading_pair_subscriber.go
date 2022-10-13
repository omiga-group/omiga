package subscribers

import (
	"context"

	"github.com/omiga-group/omiga/src/exchange/bitmart-processor/configuration"
	"github.com/omiga-group/omiga/src/exchange/bitmart-processor/mappers"
	exchangeConfiguration "github.com/omiga-group/omiga/src/exchange/shared/configuration"
	"github.com/omiga-group/omiga/src/exchange/shared/repositories"
	bitmartspotv1 "github.com/omiga-group/omiga/src/shared/clients/openapi/bitmart/spot/v1"
	"github.com/omiga-group/omiga/src/shared/enterprise/cron"
	"go.uber.org/zap"
)

type BitMartTradingPairSubscriber interface {
}

type bitMartTradingPairSubscriber struct {
	ctx                   context.Context
	logger                *zap.SugaredLogger
	bitMartConfig         configuration.BitMartConfig
	exchangeConfig        exchangeConfiguration.ExchangeConfig
	tradingPairRepository repositories.TradingPairRepository
}

func NewBitMartTradingPairSubscriber(
	ctx context.Context,
	logger *zap.SugaredLogger,
	bitMartConfig configuration.BitMartConfig,
	exchangeConfig exchangeConfiguration.ExchangeConfig,
	cronService cron.CronService,
	tradingPairRepository repositories.TradingPairRepository) (BitMartTradingPairSubscriber, error) {
	instance := &bitMartTradingPairSubscriber{
		ctx:                   ctx,
		logger:                logger,
		bitMartConfig:         bitMartConfig,
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

func (btps *bitMartTradingPairSubscriber) Run() {
	client, err := bitmartspotv1.NewClientWithResponses(btps.bitMartConfig.BaseUrl)
	if err != nil {
		btps.logger.Errorf("Failed to create client with response. Error: %v", err)

		return
	}

	response, err := client.GetAllSymbolDetailsWithResponse(btps.ctx)
	if err != nil {
		btps.logger.Errorf("Failed to call getSymbolDetails endpoint. Error: %v", err)

		return
	}

	if response.HTTPResponse.StatusCode != 200 {
		btps.logger.Errorf("Failed to call getSymbolDetails endpoint. Return status code is %d", response.HTTPResponse.StatusCode)

		return
	}

	if response.JSON200 == nil {
		btps.logger.Errorf("Returned JSON object is nil")

		return
	}

	if err = btps.tradingPairRepository.CreateTradingPairs(
		btps.ctx,
		btps.exchangeConfig.Id,
		mappers.BitMartSymbolsToTradingPairs(response.JSON200.Data.Symbols)); err != nil {
		btps.logger.Errorf("Failed to create trading pairs. Error: %v", err)

		return
	}
}
