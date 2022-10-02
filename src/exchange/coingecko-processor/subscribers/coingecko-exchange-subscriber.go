package subscribers

import (
	"context"
	"time"

	"github.com/life4/genesis/maps"
	"github.com/life4/genesis/slices"
	"github.com/omiga-group/omiga/src/exchange/coingecko-processor/configuration"
	"github.com/omiga-group/omiga/src/exchange/coingecko-processor/mappers"
	coingeckorepositories "github.com/omiga-group/omiga/src/exchange/coingecko-processor/repositories"
	"github.com/omiga-group/omiga/src/exchange/shared/models"
	"github.com/omiga-group/omiga/src/exchange/shared/repositories"
	coingeckov3 "github.com/omiga-group/omiga/src/shared/clients/openapi/coingecko/v3"
	"github.com/omiga-group/omiga/src/shared/enterprise/cron"
	timeex "github.com/omiga-group/omiga/src/shared/enterprise/time"
	"go.uber.org/zap"
)

type CoingeckoExchangeSubscriber interface {
}

type coingeckoExchangeSubscriber struct {
	ctx                context.Context
	logger             *zap.SugaredLogger
	coingeckoConfig    configuration.CoingeckoConfig
	exchanges          map[string]configuration.Exchange
	entgoClient        repositories.EntgoClient
	timeHelper         timeex.TimeHelper
	exchangeRepository coingeckorepositories.ExchangeRepository
}

func NewCoingeckoExchangeSubscriber(
	ctx context.Context,
	logger *zap.SugaredLogger,
	cronService cron.CronService,
	coingeckoConfig configuration.CoingeckoConfig,
	exchanges map[string]configuration.Exchange,
	entgoClient repositories.EntgoClient,
	timeHelper timeex.TimeHelper,
	exchangeRepository coingeckorepositories.ExchangeRepository) (CoingeckoExchangeSubscriber, error) {
	instance := &coingeckoExchangeSubscriber{
		ctx:                ctx,
		logger:             logger,
		coingeckoConfig:    coingeckoConfig,
		exchanges:          exchanges,
		entgoClient:        entgoClient,
		timeHelper:         timeHelper,
		exchangeRepository: exchangeRepository,
	}

	// Run at every second from 0 through 59.
	if _, err := cronService.GetCron().AddJob("0/1 * * * * *", instance); err != nil {
		return nil, err
	}

	return instance, nil
}

func (ces *coingeckoExchangeSubscriber) Run() {
	exchangesWithManualFeesOnlyMap := maps.Map(ces.exchanges, func(id string, exchange configuration.Exchange) (string, models.Exchange) {
		return id, mappers.FromConfigurationExchangeToExchange(exchange)
	})
	exchangesWithManualFeesOnly := maps.Values(exchangesWithManualFeesOnlyMap)

	if err := ces.exchangeRepository.CreateExchanges(ces.ctx, exchangesWithManualFeesOnly); err != nil {
		ces.logger.Errorf("Failed to create exchanges. Error: %v", err)

		return
	}

	coingeckoClient, err := coingeckov3.NewClientWithResponses(ces.coingeckoConfig.BaseUrl)
	if err != nil {
		ces.logger.Errorf("Failed to create coingecko client. Error: %v", err)
		return
	}

	perPage := 250
	exchanges := make([]coingeckov3.Exchange, 0)

	for page := 1; ; page++ {
		exchangesWithResponse, err := coingeckoClient.GetExchangesWithResponse(ces.ctx, &coingeckov3.GetExchangesParams{
			PerPage: &perPage,
			Page:    &page,
		})
		if err != nil {
			ces.logger.Errorf("Failed to get exchanges list. Error: %v", err)

			return
		}

		if exchangesWithResponse.HTTPResponse.StatusCode != 200 {
			ces.logger.Errorf(
				"Failed to get exchanges list. Return status code is %d",
				exchangesWithResponse.HTTPResponse.StatusCode)

			return
		}

		if exchangesWithResponse.JSON200 == nil || len(*exchangesWithResponse.JSON200) == 0 {
			break
		}

		exchanges = append(exchanges, *exchangesWithResponse.JSON200...)
	}

	if err := ces.exchangeRepository.CreateExchanges(
		ces.ctx,
		slices.Map(exchanges, func(exchange coingeckov3.Exchange) models.Exchange {
			if extraDetails, ok := ces.exchanges[exchange.Id]; ok {
				return mappers.FromCoingeckoExchangeToExchange(exchange, &extraDetails)
			}

			return mappers.FromCoingeckoExchangeToExchange(exchange, nil)
		})); err != nil {
		ces.logger.Errorf("Failed to create exchanges. Error: %v", err)

		return
	}

	for _, exchange := range exchanges {
		exchangeId := exchange.Id

		// This is to avoid coingecko rate limiter blocking us from querying exchanges details
		ces.timeHelper.SleepOrWaitForContextGetCancelled(ces.ctx, 2*time.Second)

		if ces.ctx.Err() == context.Canceled {
			break
		}

		exchangeIdResponse, err := coingeckoClient.GetExchangeWithResponse(
			ces.ctx,
			exchangeId)
		if err != nil {
			ces.logger.Errorf("Failed to get exchange details. Error: %v", err)

			continue
		}

		if exchangeIdResponse.HTTPResponse.StatusCode != 200 {
			ces.logger.Errorf(
				"Failed to get exchange details. Return status code is %d",
				exchangeIdResponse.HTTPResponse.StatusCode)

			continue
		}

		var mappedExchange models.Exchange

		if extraDetails, ok := ces.exchanges[mappedExchange.ExchangeId]; ok {
			mappedExchange = mappers.FromCoingeckoExchangeDetailsToExchange(
				exchangeId,
				*exchangeIdResponse.JSON200,
				&extraDetails)
		} else {
			mappedExchange = mappers.FromCoingeckoExchangeDetailsToExchange(
				exchangeId,
				*exchangeIdResponse.JSON200,
				nil)
		}

		if err := ces.exchangeRepository.CreateExchange(
			ces.ctx,
			mappedExchange); err != nil {

			ces.logger.Errorf("Failed to create exchange. Error: %v", err)

			return
		}
	}
}
