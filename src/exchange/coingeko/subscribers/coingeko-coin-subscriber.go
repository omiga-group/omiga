package subscribers

import (
	"context"

	"github.com/life4/genesis/slices"
	"github.com/omiga-group/omiga/src/exchange/coingeko/configuration"
	"github.com/omiga-group/omiga/src/exchange/coingeko/mappers"
	coingekorepositories "github.com/omiga-group/omiga/src/exchange/coingeko/repositories"
	"github.com/omiga-group/omiga/src/exchange/shared/models"
	"github.com/omiga-group/omiga/src/exchange/shared/repositories"
	coingekov3 "github.com/omiga-group/omiga/src/shared/clients/openapi/coingeko/v3"
	"github.com/omiga-group/omiga/src/shared/enterprise/cron"
	timeex "github.com/omiga-group/omiga/src/shared/enterprise/time"
	"go.uber.org/zap"
)

type CoingekoCoinSubscriber interface {
}

type coingekoCoinSubscriber struct {
	ctx            context.Context
	logger         *zap.SugaredLogger
	coingekoConfig configuration.CoingekoConfig
	exchanges      map[string]configuration.Exchange
	entgoClient    repositories.EntgoClient
	timeHelper     timeex.TimeHelper
	coinRepository coingekorepositories.CoinRepository
}

func NewCoingekoCoinSubscriber(
	ctx context.Context,
	logger *zap.SugaredLogger,
	cronService cron.CronService,
	coingekoConfig configuration.CoingekoConfig,
	exchanges map[string]configuration.Exchange,
	entgoClient repositories.EntgoClient,
	timeHelper timeex.TimeHelper,
	coinRepository coingekorepositories.CoinRepository) (CoingekoCoinSubscriber, error) {
	instance := &coingekoCoinSubscriber{
		ctx:            ctx,
		logger:         logger,
		coingekoConfig: coingekoConfig,
		exchanges:      exchanges,
		entgoClient:    entgoClient,
		timeHelper:     timeHelper,
		coinRepository: coinRepository,
	}

	if _, err := cronService.GetCron().AddJob("0 * * * *", instance); err != nil {
		return nil, err
	}

	go instance.Run()

	return instance, nil
}

func (ces *coingekoCoinSubscriber) Run() {
	coingekoClient, err := coingekov3.NewClientWithResponses(ces.coingekoConfig.BaseUrl)
	if err != nil {
		ces.logger.Errorf("Failed to create coingeko client. Error: %v", err)
		return
	}

	coinListWithResponse, err := coingekoClient.GetCoinsListWithResponse(
		ces.ctx,
		&coingekov3.GetCoinsListParams{})
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

	coins := slices.Map(*coinListWithResponse.JSON200, func(coin coingekov3.Coin) models.Coin {
		return mappers.FromCoingekoCoinToCoin(coin)
	})

	err = ces.coinRepository.CreateCoins(ces.ctx, coins)
	if err != nil {
		ces.logger.Errorf("Failed to save coins. Error: %v", err)

		return
	}
}
