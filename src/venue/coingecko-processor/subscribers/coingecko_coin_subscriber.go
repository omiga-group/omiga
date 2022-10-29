package subscribers

import (
	"context"
	"strings"

	"github.com/go-co-op/gocron"
	"github.com/life4/genesis/slices"
	coingeckov3 "github.com/omiga-group/omiga/src/venue/coingecko-processor/coingeckoclient/v3"
	"github.com/omiga-group/omiga/src/venue/coingecko-processor/configuration"
	"github.com/omiga-group/omiga/src/venue/coingecko-processor/mappers"
	"github.com/omiga-group/omiga/src/venue/shared/entities"
	"github.com/omiga-group/omiga/src/venue/shared/models"
	"github.com/omiga-group/omiga/src/venue/shared/repositories"
	"go.uber.org/zap"
)

type CoingeckoCoinSubscriber interface {
}

type coingeckoCoinSubscriber struct {
	ctx                context.Context
	logger             *zap.SugaredLogger
	coingeckoConfig    configuration.CoingeckoConfig
	exchanges          map[string]configuration.Exchange
	entgoClient        entities.EntgoClient
	currencyRepository repositories.CurrencyRepository
}

func NewCoingeckoCoinSubscriber(
	ctx context.Context,
	logger *zap.SugaredLogger,
	jobScheduler *gocron.Scheduler,
	coingeckoConfig configuration.CoingeckoConfig,
	exchanges map[string]configuration.Exchange,
	entgoClient entities.EntgoClient,
	currencyRepository repositories.CurrencyRepository) (CoingeckoCoinSubscriber, error) {
	instance := &coingeckoCoinSubscriber{
		ctx:                ctx,
		logger:             logger,
		coingeckoConfig:    coingeckoConfig,
		exchanges:          exchanges,
		entgoClient:        entgoClient,
		currencyRepository: currencyRepository,
	}

	jobScheduler.Every(5).Minutes().Do(func() {
		instance.Run()
	})

	return instance, nil
}

func (ces *coingeckoCoinSubscriber) Run() {
	coingeckoClient, err := coingeckov3.NewClientWithResponses(ces.coingeckoConfig.BaseUrl)
	if err != nil {
		ces.logger.Errorf("Failed to create coingecko client. Error: %v", err)

		return
	}

	coinListWithResponse, err := coingeckoClient.GetCoinsListWithResponse(
		ces.ctx,
		&coingeckov3.GetCoinsListParams{})
	if err != nil {
		ces.logger.Errorf("Failed to get coins list. Error: %v", err)

		return
	}

	if coinListWithResponse.HTTPResponse.StatusCode != 200 {
		ces.logger.Errorf(
			"Failed to get coin list. Return status code is %d",
			coinListWithResponse.HTTPResponse.StatusCode)

		return
	}

	if coinListWithResponse.JSON200 == nil || len(*coinListWithResponse.JSON200) == 0 {
		return
	}

	currencies := slices.Map(*coinListWithResponse.JSON200, func(coin coingeckov3.Coin) models.Currency {
		coin.Symbol = strings.ToLower(coin.Symbol)

		return mappers.FromCoingeckoCoinToCurrency(coin)
	})

	_, err = ces.currencyRepository.CreateCurrencies(ces.ctx, currencies)
	if err != nil {
		ces.logger.Errorf("Failed to save coins. Error: %v", err)

		return
	}
}
