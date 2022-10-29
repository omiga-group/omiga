package subscribers

import (
	"context"

	"github.com/go-co-op/gocron"
	"github.com/omiga-group/omiga/src/venue/crypto-processor/configuration"
	cryptov2 "github.com/omiga-group/omiga/src/venue/crypto-processor/cryptoclient/v2"
	"github.com/omiga-group/omiga/src/venue/crypto-processor/mappers"
	"github.com/omiga-group/omiga/src/venue/shared/repositories"
	"go.uber.org/zap"
)

type CryptoTradingPairSubscriber interface {
}

type cryptoTradingPairSubscriber struct {
	ctx                   context.Context
	logger                *zap.SugaredLogger
	venueConfig           configuration.CryptoConfig
	tradingPairRepository repositories.TradingPairRepository
}

func NewCryptoTradingPairSubscriber(
	ctx context.Context,
	logger *zap.SugaredLogger,
	venueConfig configuration.CryptoConfig,
	jobScheduler *gocron.Scheduler,
	tradingPairRepository repositories.TradingPairRepository) (CryptoTradingPairSubscriber, error) {

	instance := &cryptoTradingPairSubscriber{
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

func (ctps *cryptoTradingPairSubscriber) Run() {
	ctps.logger.Infof("Start trading pairs sync for Venue: %s ...", ctps.venueConfig.Id)

	client, err := cryptov2.NewClientWithResponses(ctps.venueConfig.BaseUrl)
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
		ctps.venueConfig.Id,
		mappers.CryptoInstrumentsToTradingPairs(response.JSON200.Result.Instruments)); err != nil {
		ctps.logger.Errorf("Failed to create trading pairs. Error: %v", err)

		return
	}

	ctps.logger.Infof("Finished syncing trading pairs for Venue: %s", ctps.venueConfig.Id)
}
