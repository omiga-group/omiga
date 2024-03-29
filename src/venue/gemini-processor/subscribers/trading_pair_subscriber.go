package subscribers

import (
	"context"
	"fmt"

	"go.uber.org/zap"

	"github.com/go-co-op/gocron"
	"github.com/life4/genesis/maps"
	"github.com/life4/genesis/slices"
	"github.com/omiga-group/omiga/src/venue/gemini-processor/configuration"
	geminiv1 "github.com/omiga-group/omiga/src/venue/gemini-processor/geminiclient/v1"
	"github.com/omiga-group/omiga/src/venue/gemini-processor/mappers"
	"github.com/omiga-group/omiga/src/venue/shared/repositories"
)

type GeminiTradingPairSubscriber interface {
}

type geminiTradingPairSubscriber struct {
	ctx                   context.Context
	logger                *zap.SugaredLogger
	venueConfig           configuration.GeminiConfig
	tradingPairRepository repositories.TradingPairRepository
}

func NewGeminiTradingPairSubscriber(
	ctx context.Context,
	logger *zap.SugaredLogger,
	venueConfig configuration.GeminiConfig,
	jobScheduler *gocron.Scheduler,
	tradingPairRepository repositories.TradingPairRepository) (GeminiTradingPairSubscriber, error) {

	instance := &geminiTradingPairSubscriber{
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

func (gtps *geminiTradingPairSubscriber) Run() {
	gtps.logger.Infof("Start trading pairs sync for Venue: %s ...", gtps.venueConfig.Id)

	client, err := geminiv1.NewClientWithResponses(gtps.venueConfig.ApiUrl)
	if err != nil {
		gtps.logger.Errorf("Failed to create client with response. Error: %v", err)

		return
	}

	symbolsReponse, err := client.GetAllSymbolsWithResponse(gtps.ctx)
	if err != nil {
		gtps.logger.Errorf("Failed to call getAllSymbols endpoint. Error: %v", err)

		return
	}

	if symbolsReponse.HTTPResponse.StatusCode != 200 {
		gtps.logger.Errorf("Failed to call getAllSymbols endpoint. Return status code is %d", symbolsReponse.HTTPResponse.StatusCode)

		return
	}

	if symbolsReponse.JSON200 == nil {
		gtps.logger.Errorf("Returned JSON object is nil")

		return
	}

	all, err := slices.ReduceWhile(*symbolsReponse.JSON200,
		make(map[string]geminiv1.TradingPair),
		func(el string, acc map[string]geminiv1.TradingPair) (map[string]geminiv1.TradingPair, error) {
			symbolDetailsResponse, err := client.GetSymbolDetailsWithResponse(gtps.ctx, el)
			if err != nil {
				return acc, err
			}

			if symbolDetailsResponse.HTTPResponse.StatusCode != 200 {
				return acc, fmt.Errorf("failed to get symbol details. Return status code is %d", symbolsReponse.HTTPResponse.StatusCode)
			}

			if symbolDetailsResponse.JSON200 == nil {
				return acc, fmt.Errorf("returned JSON object is nil")
			}

			acc[el] = *symbolDetailsResponse.JSON200

			return acc, nil
		})

	if err != nil {
		gtps.logger.Errorf("Failed to get symbol details. Return status code is %d", symbolsReponse.HTTPResponse.StatusCode)

		return
	}

	if err = gtps.tradingPairRepository.CreateTradingPairs(
		gtps.ctx,
		gtps.venueConfig.Id,
		mappers.GeminiTradingPairsToTradingPairs(maps.Values(all))); err != nil {
		gtps.logger.Errorf("Failed to create trading pairs. Error: %v", err)

		return
	}

	gtps.logger.Infof("Finished syncing trading pairs for Venue: %s", gtps.venueConfig.Id)
}
