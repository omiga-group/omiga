package subscribers

import (
	"context"

	"github.com/omiga-group/omiga/src/shared/enterprise/cron"
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
	coinbaseConfig        configuration.CoinbaseConfig
	tradingPairRepository repositories.TradingPairRepository
}

func NewCoinbaseTradingPairSubscriber(
	ctx context.Context,
	logger *zap.SugaredLogger,
	coinbaseConfig configuration.CoinbaseConfig,
	cronService cron.CronService,
	tradingPairRepository repositories.TradingPairRepository) (CoinbaseTradingPairSubscriber, error) {

	instance := &coinbaseTradingPairSubscriber{
		ctx:                   ctx,
		logger:                logger,
		coinbaseConfig:        coinbaseConfig,
		tradingPairRepository: tradingPairRepository,
	}

	// Run at every minute from 0 through 59.
	if _, err := cronService.GetCron().AddJob("* 0/1 * * * *", instance); err != nil {
		return nil, err
	}

	go instance.Run()

	return instance, nil
}

func (ctps *coinbaseTradingPairSubscriber) Run() {
	client := coinbasepro.NewClient()

	client.UpdateConfig(&coinbasepro.ClientConfig{
		BaseURL:    ctps.coinbaseConfig.BaseUrl,
		Key:        ctps.coinbaseConfig.ApiKey,
		Passphrase: ctps.coinbaseConfig.Passphrase,
		Secret:     ctps.coinbaseConfig.SecretKey,
	})

	products, err := client.GetProducts()
	if err != nil {
		ctps.logger.Errorf("Failed to call getProducts endpoint. Error: %v", err)

		return
	}

	if err = ctps.tradingPairRepository.CreateTradingPairs(
		ctps.ctx,
		ctps.coinbaseConfig.Id,
		mappers.CoinbaseProductsToTradingPairs(products)); err != nil {
		ctps.logger.Errorf("Failed to create trading pairs. Error: %v", err)

		return
	}
}
