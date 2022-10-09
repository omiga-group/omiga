package subscribers

import (
	"context"
	"encoding/json"

	"go.uber.org/zap"

	"github.com/omiga-group/omiga/src/exchange/gemini-processor/configuration"
	exchangeConfiguration "github.com/omiga-group/omiga/src/exchange/shared/configuration"
	"github.com/omiga-group/omiga/src/exchange/shared/repositories"
	geminiClient "github.com/omiga-group/omiga/src/shared/clients/openapi/gemini/v1"
	"github.com/omiga-group/omiga/src/shared/enterprise/cron"
)

type TradingPairsSubscriber interface {
}

type tradingPairsSubscriber struct {
	ctx                   context.Context
	logger                *zap.SugaredLogger
	geminiConfig          configuration.GeminiConfig
	exchangeConfig        exchangeConfiguration.ExchangeConfig
	tradingPairRepository repositories.TradingPairRepository
}

func NewGeminiTradingPairsSubscriber(
	ctx context.Context,
	logger *zap.SugaredLogger,
	geminiConfig configuration.GeminiConfig,
	exchangeConfig exchangeConfiguration.ExchangeConfig,
	cronService cron.CronService,
	tradingPairRepository repositories.TradingPairRepository) (TradingPairsSubscriber, error) {

	instance := &tradingPairsSubscriber{
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

func (tps *tradingPairsSubscriber) Run() {
	client, err := geminiClient.NewClientWithResponses(tps.geminiConfig.ApiUrl)
	if err != nil {
		tps.logger.Errorf("Failed to call assetPairs endpoint. Error: %v", err)
		return
	}

	priceFeedRes, err := client.GetPricefeedWithResponse(tps.ctx)
	if err != nil {
		tps.logger.Errorf("Failed to call assetPairs endpoint. Error: %v", err)
		return
	}

	if priceFeedRes.HTTPResponse.StatusCode != 200 {
		tps.logger.Errorf("Failed to get exchanges list. Return status code is %d", priceFeedRes.HTTPResponse.StatusCode)
		return
	}

	if priceFeedRes.JSON200 == nil || len(*priceFeedRes.JSON200) == 0 || priceFeedRes.Body == nil {
		return
	}

	priceFeeds := make([]geminiClient.PriceFeed, 0)
	err = json.Unmarshal(priceFeedRes.Body, &priceFeeds)
	if err != nil {
		tps.logger.Errorf("Failed to call assetPairs endpoint. Error: %v", err)
		return
	}
}
