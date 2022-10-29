package subscribers

import (
	"context"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/life4/genesis/maps"
	"github.com/life4/genesis/slices"
	timeex "github.com/omiga-group/omiga/src/shared/enterprise/time"
	coingeckov3 "github.com/omiga-group/omiga/src/venue/coingecko-processor/coingeckoclient/v3"
	"github.com/omiga-group/omiga/src/venue/coingecko-processor/configuration"
	"github.com/omiga-group/omiga/src/venue/coingecko-processor/mappers"
	"github.com/omiga-group/omiga/src/venue/shared/entities"
	"github.com/omiga-group/omiga/src/venue/shared/models"
	"github.com/omiga-group/omiga/src/venue/shared/repositories"
	"go.uber.org/zap"
)

type CoingeckoExchangeSubscriber interface {
}

type coingeckoExchangeSubscriber struct {
	ctx             context.Context
	logger          *zap.SugaredLogger
	coingeckoConfig configuration.CoingeckoConfig
	exchanges       map[string]configuration.Exchange
	entgoClient     entities.EntgoClient
	timeHelper      timeex.TimeHelper
	venueRepository repositories.VenueRepository
}

func NewCoingeckoExchangeSubscriber(
	ctx context.Context,
	logger *zap.SugaredLogger,
	jobScheduler *gocron.Scheduler,
	coingeckoConfig configuration.CoingeckoConfig,
	exchanges map[string]configuration.Exchange,
	entgoClient entities.EntgoClient,
	timeHelper timeex.TimeHelper,
	venueRepository repositories.VenueRepository) (CoingeckoExchangeSubscriber, error) {
	instance := &coingeckoExchangeSubscriber{
		ctx:             ctx,
		logger:          logger,
		coingeckoConfig: coingeckoConfig,
		exchanges:       exchanges,
		entgoClient:     entgoClient,
		timeHelper:      timeHelper,
		venueRepository: venueRepository,
	}

	jobScheduler.Every(5).Minutes().Do(func() {
		instance.Run()
	})

	return instance, nil
}

func (ces *coingeckoExchangeSubscriber) Run() {
	venuesWithManualFeesOnlyMap := maps.Map(ces.exchanges, func(id string, exchange configuration.Exchange) (string, models.Venue) {
		return id, mappers.FromConfigurationExchangeToExchange(exchange)
	})
	venuesWithManualFeesOnly := maps.Values(venuesWithManualFeesOnlyMap)

	if _, err := ces.venueRepository.CreateVenues(ces.ctx, venuesWithManualFeesOnly); err != nil {
		ces.logger.Errorf("Failed to create venues. Error: %v", err)

		return
	}

	coingeckoClient, err := coingeckov3.NewClientWithResponses(ces.coingeckoConfig.BaseUrl)
	if err != nil {
		ces.logger.Errorf("Failed to create coingecko client. Error: %v", err)
		return
	}

	perPage := 250
	venues := make([]coingeckov3.Exchange, 0)

	for page := 1; ; page++ {
		ces.timeHelper.SleepOrWaitForContextGetCancelled(ces.ctx, 2*time.Second)

		exchangesWithResponse, err := coingeckoClient.GetExchangesWithResponse(ces.ctx, &coingeckov3.GetExchangesParams{
			PerPage: &perPage,
			Page:    &page,
		})
		if err != nil {
			ces.logger.Errorf("Failed to get venues list. Error: %v", err)

			return
		}

		if exchangesWithResponse.HTTPResponse.StatusCode != 200 {
			ces.logger.Errorf(
				"Failed to get venues list. Return status code is %d",
				exchangesWithResponse.HTTPResponse.StatusCode)

			return
		}

		if exchangesWithResponse.JSON200 == nil || len(*exchangesWithResponse.JSON200) == 0 {
			break
		}

		venues = append(venues, *exchangesWithResponse.JSON200...)
	}

	if _, err := ces.venueRepository.CreateVenues(
		ces.ctx,
		slices.Map(venues, func(exchange coingeckov3.Exchange) models.Venue {
			if extraDetails, ok := ces.exchanges[exchange.Id]; ok {
				return mappers.FromCoingeckoExchangeToExchange(exchange, &extraDetails)
			}

			return mappers.FromCoingeckoExchangeToExchange(exchange, nil)
		})); err != nil {
		ces.logger.Errorf("Failed to create venues. Error: %v", err)

		return
	}

	for _, venue := range venues {
		venueId := venue.Id

		// This is to avoid coingecko rate limiter blocking us from querying venues details
		ces.timeHelper.SleepOrWaitForContextGetCancelled(ces.ctx, 2*time.Second)

		if ces.ctx.Err() == context.Canceled {
			break
		}

		venueIdResponse, err := coingeckoClient.GetExchangeWithResponse(
			ces.ctx,
			venueId)
		if err != nil {
			ces.logger.Errorf("Failed to get venue details. Error: %v", err)

			continue
		}

		if venueIdResponse.HTTPResponse.StatusCode != 200 {
			ces.logger.Errorf(
				"Failed to get venue details. Return status code is %d",
				venueIdResponse.HTTPResponse.StatusCode)

			continue
		}

		var mappedVenue models.Venue

		if extraDetails, ok := ces.exchanges[mappedVenue.VenueId]; ok {
			mappedVenue = mappers.FromCoingeckoExchangeDetailsToExchange(
				venueId,
				*venueIdResponse.JSON200,
				&extraDetails)
		} else {
			mappedVenue = mappers.FromCoingeckoExchangeDetailsToExchange(
				venueId,
				*venueIdResponse.JSON200,
				nil)
		}

		if _, err := ces.venueRepository.CreateVenue(
			ces.ctx,
			mappedVenue); err != nil {

			ces.logger.Errorf("Failed to create venue. Error: %v", err)

			return
		}
	}
}
