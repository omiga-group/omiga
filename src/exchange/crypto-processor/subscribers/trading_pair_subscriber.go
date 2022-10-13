package subscribers

import (
	"context"

	"github.com/omiga-group/omiga/src/exchange/crypto-processor/configuration"
	"github.com/omiga-group/omiga/src/exchange/crypto-processor/mappers"
	exchangeConfiguration "github.com/omiga-group/omiga/src/exchange/shared/configuration"
	"github.com/omiga-group/omiga/src/exchange/shared/repositories"
	cryptov2 "github.com/omiga-group/omiga/src/shared/clients/openapi/crypto/v2"
	"github.com/omiga-group/omiga/src/shared/enterprise/cron"
	"go.uber.org/zap"
)

type CryptoTradingPairSubscriber interface {
}

type cryptoTradingPairSubscriber struct {
	ctx                   context.Context
	logger                *zap.SugaredLogger
	cryptoConfig          configuration.CryptoConfig
	exchangeConfig        exchangeConfiguration.ExchangeConfig
	tradingPairRepository repositories.TradingPairRepository
}

func NewCryptoTradingPairSubscriber(
	ctx context.Context,
	logger *zap.SugaredLogger,
	cryptoConfig configuration.CryptoConfig,
	exchangeConfig exchangeConfiguration.ExchangeConfig,
	cronService cron.CronService,
	tradingPairRepository repositories.TradingPairRepository) (CryptoTradingPairSubscriber, error) {

	instance := &cryptoTradingPairSubscriber{
		ctx:                   ctx,
		logger:                logger,
		cryptoConfig:          cryptoConfig,
		exchangeConfig:        exchangeConfig,
		tradingPairRepository: tradingPairRepository,
	}

	// Run at every minute from 0 through 59.
	if _, err := cronService.GetCron().AddJob("* 0/1 * * * *", instance); err != nil {
		return nil, err
	}

	go instance.Run()

	return instance, nil
}

func (ctps *cryptoTradingPairSubscriber) Run() {
	client, err := cryptov2.NewClientWithResponses(ctps.cryptoConfig.BaseUrl)
	if err != nil {
		ctps.logger.Errorf("Failed to create client with response. Error: %v", err)

		return
	}

	response, err := client.GetAllInstrumentsWithResponse(ctps.ctx)
	if err != nil {
		ctps.logger.Errorf("Failed to call getAllInstruments endpoint. Error: %v", err)

		return
	}

	if response.HTTPResponse.StatusCode != 200 {
		ctps.logger.Errorf("Failed to call getAllInstruments endpoint. Return status code is %d", response.HTTPResponse.StatusCode)

		return
	}

	if response.JSON200 == nil {
		ctps.logger.Errorf("Returned JSON object is nil")

		return
	}

	if err = ctps.tradingPairRepository.CreateTradingPairs(
		ctps.ctx,
		ctps.exchangeConfig.Id,
		mappers.CryptoInstrumentsToTradingPairs(response.JSON200.Result.Instruments)); err != nil {
		ctps.logger.Errorf("Failed to create trading pairs. Error: %v", err)

		return
	}
}
