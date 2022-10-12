package subscribers

import (
	"context"
	"fmt"

	"go.uber.org/zap"

	"github.com/life4/genesis/maps"
	"github.com/life4/genesis/slices"
	"github.com/omiga-group/omiga/src/exchange/gemini-processor/configuration"
	"github.com/omiga-group/omiga/src/exchange/gemini-processor/mappers"
	exchangeConfiguration "github.com/omiga-group/omiga/src/exchange/shared/configuration"
	"github.com/omiga-group/omiga/src/exchange/shared/repositories"
	geminiClient "github.com/omiga-group/omiga/src/shared/clients/openapi/gemini/v1"
	"github.com/omiga-group/omiga/src/shared/enterprise/cron"
)

type GeminiTradingPairSubscriber interface {
}

type tradingPairSubscriber struct {
	ctx                   context.Context
	logger                *zap.SugaredLogger
	geminiConfig          configuration.GeminiConfig
	exchangeConfig        exchangeConfiguration.ExchangeConfig
	tradingPairRepository repositories.TradingPairRepository
}

func NewGeminiTradingPairSubscriber(
	ctx context.Context,
	logger *zap.SugaredLogger,
	geminiConfig configuration.GeminiConfig,
	exchangeConfig exchangeConfiguration.ExchangeConfig,
	cronService cron.CronService,
	tradingPairRepository repositories.TradingPairRepository) (GeminiTradingPairSubscriber, error) {

	instance := &tradingPairSubscriber{
		ctx:                   ctx,
		logger:                logger,
		geminiConfig:          geminiConfig,
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

func (tps *tradingPairSubscriber) Run() {
	client, err := geminiClient.NewClientWithResponses(tps.geminiConfig.ApiUrl)
	if err != nil {
		tps.logger.Errorf("Failed to call assetPairs endpoint. Error: %v", err)
		return
	}

	priceFeedRes, err := client.GetSymbolsWithResponse(tps.ctx)
	if err != nil {
		tps.logger.Errorf("Failed to call assetPairs endpoint. Error: %v", err)
		return
	}

	if priceFeedRes.HTTPResponse.StatusCode != 200 {
		tps.logger.Errorf("Failed to get exchanges list. Return status code is %d", priceFeedRes.HTTPResponse.StatusCode)
		return
	}

	if priceFeedRes.JSON200 == nil {
		return
	}

	all, err := slices.ReduceWhile(*priceFeedRes.JSON200,
		make(map[string]geminiClient.TradingPair),
		func(el string, acc map[string]geminiClient.TradingPair) (map[string]geminiClient.TradingPair, error) {
			details, err := client.GetSymbolsDetailsSymbolWithResponse(tps.ctx, el)
			if err != nil {
				return acc, err
			}

			if details.HTTPResponse.StatusCode != 200 {
				return acc, fmt.Errorf("failed to get exchanges list. Return status code is %d", priceFeedRes.HTTPResponse.StatusCode)
			}

			acc[el] = *details.JSON200
			return acc, err
		})
	if err != nil {
		tps.logger.Errorf("Failed to get exchanges list. Return status code is %d", priceFeedRes.HTTPResponse.StatusCode)
		return
	}

	if err = tps.tradingPairRepository.CreateTradingPairs(
		tps.ctx,
		tps.exchangeConfig.Id,
		mappers.GeminiTradingPairsToTradingPairs(maps.Values(all))); err != nil {
		tps.logger.Errorf("Failed to create trading pairs. Error: %v", err)

		return
	}
}
