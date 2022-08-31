package subscribers

import (
	"context"
	"time"

	"github.com/life4/genesis/maps"
	"github.com/life4/genesis/slices"
	"github.com/omiga-group/omiga/src/exchange/coingeko/configuration"
	"github.com/omiga-group/omiga/src/exchange/coingeko/mappers"
	"github.com/omiga-group/omiga/src/exchange/coingeko/models"
	coingekorepositories "github.com/omiga-group/omiga/src/exchange/coingeko/repositories"
	"github.com/omiga-group/omiga/src/exchange/shared/repositories"
	coingekov3 "github.com/omiga-group/omiga/src/shared/clients/openapi/coingeko/v3"
	"github.com/omiga-group/omiga/src/shared/enterprise/cron"
	timeex "github.com/omiga-group/omiga/src/shared/enterprise/time"
	"go.uber.org/zap"
)

type CoingekoSubscriber interface {
}

type coingekoSubscriber struct {
	ctx                context.Context
	logger             *zap.SugaredLogger
	coingekoConfig     configuration.CoingekoConfig
	exchanges          map[string]configuration.Exchange
	entgoClient        repositories.EntgoClient
	timeHelper         timeex.TimeHelper
	exchangeRepository coingekorepositories.ExchangeRepository
}

func NewCoingekoSubscriber(
	ctx context.Context,
	logger *zap.SugaredLogger,
	cronService cron.CronService,
	coingekoConfig configuration.CoingekoConfig,
	exchanges map[string]configuration.Exchange,
	entgoClient repositories.EntgoClient,
	timeHelper timeex.TimeHelper,
	exchangeRepository coingekorepositories.ExchangeRepository) (CoingekoSubscriber, error) {
	instance := &coingekoSubscriber{
		ctx:                ctx,
		logger:             logger,
		coingekoConfig:     coingekoConfig,
		exchanges:          exchanges,
		entgoClient:        entgoClient,
		timeHelper:         timeHelper,
		exchangeRepository: exchangeRepository,
	}

	if _, err := cronService.GetCron().AddJob("0/1 * * * * *", instance); err != nil {
		return nil, err
	}

	return instance, nil
}

func (cs *coingekoSubscriber) Run() {
	exchangesWithManualFeesOnlyMap := maps.Map(cs.exchanges, func(id string, exchange configuration.Exchange) (string, models.Exchange) {
		return id, mappers.FromConfigurationExchangeToExchange(exchange)
	})
	exchangesWithManualFeesOnly := maps.Values(exchangesWithManualFeesOnlyMap)

	if err := cs.exchangeRepository.CreateExchanges(cs.ctx, exchangesWithManualFeesOnly); err != nil {
		cs.logger.Errorf("Failed to create exchanges. Error: %v", err)

		return
	}

	coingekoClient, err := coingekov3.NewClientWithResponses(cs.coingekoConfig.BaseUrl)
	if err != nil {
		cs.logger.Errorf("Failed to create coingeko client. Error: %v", err)
		return
	}

	perPage := 250
	exchanges := make([]coingekov3.Exchange, 0)

	for page := 1; ; page++ {
		exchangesWithResponse, err := coingekoClient.GetExchangesWithResponse(cs.ctx, &coingekov3.GetExchangesParams{
			PerPage: &perPage,
			Page:    &page,
		})
		if err != nil {
			cs.logger.Errorf("Failed to get exchanges list. Error: %v", err)

			return
		}

		if exchangesWithResponse.HTTPResponse.StatusCode != 200 {
			cs.logger.Errorf(
				"Failed to get exchanges list. Return status code is %d",
				exchangesWithResponse.HTTPResponse.StatusCode)

			return
		}

		if exchangesWithResponse.JSON200 == nil || len(*exchangesWithResponse.JSON200) == 0 {
			break
		}

		exchanges = append(exchanges, *exchangesWithResponse.JSON200...)
	}

	if err := cs.exchangeRepository.CreateExchanges(
		cs.ctx,
		slices.Map(exchanges, func(exchange coingekov3.Exchange) models.Exchange {
			if extraDetails, ok := cs.exchanges[exchange.Id]; ok {
				return mappers.FromCoingekoExchangeToExchange(exchange, &extraDetails)
			}

			return mappers.FromCoingekoExchangeToExchange(exchange, nil)
		})); err != nil {
		cs.logger.Errorf("Failed to create exchanges. Error: %v", err)

		return
	}

	for _, exchange := range exchanges {
		exchangeId := exchange.Id

		// This is to avoid coingeko rate limiter blocking us from querying exchanges details
		cs.timeHelper.SleepOrWaitForContextGetCancelled(cs.ctx, 2*time.Second)

		if cs.ctx.Err() == context.Canceled {
			break
		}

		exchangeIdResponse, err := coingekoClient.GetExchangeWithResponse(
			cs.ctx,
			exchangeId)
		if err != nil {
			cs.logger.Errorf("Failed to get exchange details. Error: %v", err)

			continue
		}

		if exchangeIdResponse.HTTPResponse.StatusCode != 200 {
			cs.logger.Errorf(
				"Failed to get exchange details. Return status code is %d",
				exchangeIdResponse.HTTPResponse.StatusCode)

			continue
		}

		var mappedExchange models.Exchange

		if extraDetails, ok := cs.exchanges[mappedExchange.ExchangeId]; ok {
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

		if err := cs.exchangeRepository.CreateExchange(
			cs.ctx,
			mappedExchange); err != nil {

			cs.logger.Errorf("Failed to create exchange. Error: %v", err)

			return
		}
	}
}
