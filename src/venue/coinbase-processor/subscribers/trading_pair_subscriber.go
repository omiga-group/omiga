package subscribers

import (
	"context"

	"github.com/go-co-op/gocron"
	"github.com/omiga-group/omiga/src/venue/coinbase-processor/configuration"
	"github.com/omiga-group/omiga/src/venue/coinbase-processor/mappers"
	"github.com/omiga-group/omiga/src/venue/shared/repositories"
	"github.com/preichenberger/go-coinbasepro/v2"
	"go.uber.org/zap"
)

type CoinbaseTradingPairSubscriber interface {
}

type coinbaseTradingPairSubscriber struct {
	ctx                   context.Context
	logger                *zap.SugaredLogger
	venueConfig           configuration.CoinbaseConfig
	tradingPairRepository repositories.TradingPairRepository
}

func NewCoinbaseTradingPairSubscriber(
	ctx context.Context,
	logger *zap.SugaredLogger,
	venueConfig configuration.CoinbaseConfig,
	jobScheduler *gocron.Scheduler,
	tradingPairRepository repositories.TradingPairRepository) (CoinbaseTradingPairSubscriber, error) {

	instance := &coinbaseTradingPairSubscriber{
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

func (ctps *coinbaseTradingPairSubscriber) Run() {
	ctps.logger.Infof("Start trading pairs sync for Venue: %s ...", ctps.venueConfig.Id)

	client := coinbasepro.NewClient()

	client.UpdateConfig(&coinbasepro.ClientConfig{
		BaseURL:    ctps.venueConfig.BaseUrl,
		Key:        ctps.venueConfig.ApiKey,
		Passphrase: ctps.venueConfig.Passphrase,
		Secret:     ctps.venueConfig.SecretKey,
	})

	products, err := client.GetProducts()
	if err != nil {
		ctps.logger.Errorf("Failed to call getProducts endpoint. Error: %v", err)

		return
	}

	if err = ctps.tradingPairRepository.CreateTradingPairs(
		ctps.ctx,
		ctps.venueConfig.Id,
		mappers.CoinbaseProductsToTradingPairs(products)); err != nil {
		ctps.logger.Errorf("Failed to create trading pairs. Error: %v", err)

		return
	}

	ctps.logger.Infof("Finished syncing trading pairs for Venue: %s", ctps.venueConfig.Id)
}
