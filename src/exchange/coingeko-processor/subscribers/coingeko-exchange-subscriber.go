package subscribers

import (
	"context"
	"time"

	"github.com/life4/genesis/maps"
	"github.com/life4/genesis/slices"
	"github.com/omiga-group/omiga/src/exchange/coingeko-processor/configuration"
	"github.com/omiga-group/omiga/src/exchange/coingeko-processor/mappers"
	coingekorepositories "github.com/omiga-group/omiga/src/exchange/coingeko-processor/repositories"
	"github.com/omiga-group/omiga/src/exchange/shared/models"
	"github.com/omiga-group/omiga/src/exchange/shared/repositories"
	coingekov3 "github.com/omiga-group/omiga/src/shared/clients/openapi/coingeko/v3"
	"github.com/omiga-group/omiga/src/shared/enterprise/cron"
	timeex "github.com/omiga-group/omiga/src/shared/enterprise/time"
	"go.uber.org/zap"
)

type CoingekoExchangeSubscriber interface {
}

type coingekoExchangeSubscriber struct {
	ctx                context.Context
	logger             *zap.SugaredLogger
	coingekoConfig     configuration.CoingekoConfig
	exchanges          map[string]configuration.Exchange
	entgoClient        repositories.EntgoClient
	timeHelper         timeex.TimeHelper
	exchangeRepository coingekorepositories.ExchangeRepository
}

func NewCoingekoExchangeSubscriber(
	ctx context.Context,
	logger *zap.SugaredLogger,
	cronService cron.CronService,
	coingekoConfig configuration.CoingekoConfig,
	exchanges map[string]configuration.Exchange,
	entgoClient repositories.EntgoClient,
	timeHelper timeex.TimeHelper,
	exchangeRepository coingekorepositories.ExchangeRepository) (CoingekoExchangeSubscriber, error) {
	instance := &coingekoExchangeSubscriber{
		ctx:                ctx,
		logger:             logger,
		coingekoConfig:     coingekoConfig,
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

func (ces *coingekoExchangeSubscriber) Run() {
	exchangesWithManualFeesOnlyMap := maps.Map(ces.exchanges, func(id string, exchange configuration.Exchange) (string, models.Exchange) {
		return id, mappers.FromConfigurationExchangeToExchange(exchange)
	})
	exchangesWithManualFeesOnly := maps.Values(exchangesWithManualFeesOnlyMap)

	if err := ces.exchangeRepository.CreateExchanges(ces.ctx, exchangesWithManualFeesOnly); err != nil {
		ces.logger.Errorf("Failed to create exchanges. Error: %v", err)

		return
	}

	coingekoClient, err := coingekov3.NewClientWithResponses(ces.coingekoConfig.BaseUrl)
	if err != nil {
		ces.logger.Errorf("Failed to create coingeko client. Error: %v", err)
		return
	}

	perPage := 250
	exchanges := make([]coingekov3.Exchange, 0)

	for page := 1; ; page++ {
		exchangesWithResponse, err := coingekoClient.GetExchangesWithResponse(ces.ctx, &coingekov3.GetExchangesParams{
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
		slices.Map(exchanges, func(exchange coingekov3.Exchange) models.Exchange {
			if extraDetails, ok := ces.exchanges[exchange.Id]; ok {
				return mappers.FromCoingekoExchangeToExchange(exchange, &extraDetails)
			}

			return mappers.FromCoingekoExchangeToExchange(exchange, nil)
		})); err != nil {
		ces.logger.Errorf("Failed to create exchanges. Error: %v", err)

		return
	}

	for _, exchange := range exchanges {
		exchangeId := exchange.Id

		// This is to avoid coingeko rate limiter blocking us from querying exchanges details
		ces.timeHelper.SleepOrWaitForContextGetCancelled(ces.ctx, 2*time.Second)

		if ces.ctx.Err() == context.Canceled {
			break
		}

		exchangeIdResponse, err := coingekoClient.GetExchangeWithResponse(
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
			mappedExchange = mappers.FromCoingekoExchangeDetailsToExchange(
				exchangeId,
				*exchangeIdResponse.JSON200,
				&extraDetails)
		} else {
			mappedExchange = mappers.FromCoingekoExchangeDetailsToExchange(
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
