package subscribers

import (
	"context"
	"strings"

	"github.com/life4/genesis/slices"
	"github.com/omiga-group/omiga/src/exchange/coingecko-processor/configuration"
	"github.com/omiga-group/omiga/src/exchange/coingecko-processor/mappers"
	"github.com/omiga-group/omiga/src/exchange/shared/entities"
	"github.com/omiga-group/omiga/src/exchange/shared/models"
	"github.com/omiga-group/omiga/src/exchange/shared/repositories"
	coingeckov3 "github.com/omiga-group/omiga/src/shared/clients/openapi/coingecko/v3"
	"github.com/omiga-group/omiga/src/shared/enterprise/cron"
	"go.uber.org/zap"
)

type CoingeckoCoinSubscriber interface {
}

type coingeckoCoinSubscriber struct {
	ctx             context.Context
	logger          *zap.SugaredLogger
	coingeckoConfig configuration.CoingeckoConfig
	exchanges       map[string]configuration.Exchange
	entgoClient     entities.EntgoClient
	coinRepository  repositories.CoinRepository
}

func NewCoingeckoCoinSubscriber(
	ctx context.Context,
	logger *zap.SugaredLogger,
	cronService cron.CronService,
	coingeckoConfig configuration.CoingeckoConfig,
	exchanges map[string]configuration.Exchange,
	entgoClient entities.EntgoClient,
	coinRepository repositories.CoinRepository) (CoingeckoCoinSubscriber, error) {
	instance := &coingeckoCoinSubscriber{
		ctx:             ctx,
		logger:          logger,
		coingeckoConfig: coingeckoConfig,
		exchanges:       exchanges,
		entgoClient:     entgoClient,
		coinRepository:  coinRepository,
	}

	// Run at minute 0
	if _, err := cronService.GetCron().AddJob("* 0 * * * *", instance); err != nil {
		return nil, err
	}

	go instance.Run()

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
		ces.logger.Errorf("Failed to get coins list list. Error: %v", err)

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

	coins := slices.Map(*coinListWithResponse.JSON200, func(coin coingeckov3.Coin) models.Coin {
		coin.Symbol = strings.ToLower(coin.Symbol)

		return mappers.FromCoingeckoCoinToCoin(coin)
	})

	err = ces.coinRepository.CreateCoins(ces.ctx, coins)
	if err != nil {
		ces.logger.Errorf("Failed to save coins. Error: %v", err)

		return
	}
}
