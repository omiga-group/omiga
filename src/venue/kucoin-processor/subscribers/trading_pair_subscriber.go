package subscribers

import (
	"context"

	"github.com/Kucoin/kucoin-go-sdk"
	"github.com/omiga-group/omiga/src/shared/enterprise/cron"
	"github.com/omiga-group/omiga/src/venue/kucoin-processor/configuration"
	"github.com/omiga-group/omiga/src/venue/kucoin-processor/mappers"
	"github.com/omiga-group/omiga/src/venue/shared/repositories"
	"go.uber.org/zap"
)

type KucoinTradingPairSubscriber interface {
}

type kuCoinTradingPairSubscriber struct {
	ctx                   context.Context
	logger                *zap.SugaredLogger
	venueConfig           configuration.KucoinConfig
	tradingPairRepository repositories.TradingPairRepository
}

func NewKucoinTradingPairSubscriber(
	ctx context.Context,
	logger *zap.SugaredLogger,
	venueConfig configuration.KucoinConfig,
	cronService cron.CronService,
	tradingPairRepository repositories.TradingPairRepository) (KucoinTradingPairSubscriber, error) {

	instance := &kuCoinTradingPairSubscriber{
		ctx:                   ctx,
		logger:                logger,
		venueConfig:           venueConfig,
		tradingPairRepository: tradingPairRepository,
	}

	// Run at every 5th minute from 0 through 59..
	if _, err := cronService.GetCron().AddJob("* 0/5 * * * *", instance); err != nil {
		return nil, err
	}

	return instance, nil
}

func (ktps *kuCoinTradingPairSubscriber) Run() {
	ktps.logger.Errorf("Start trading pairs sync for Venue: %s ...", ktps.venueConfig.Id)

	apiService := kucoin.NewApiService(
		kucoin.ApiKeyOption(ktps.venueConfig.ApiKey),
		kucoin.ApiPassPhraseOption(ktps.venueConfig.Passphrase),
		kucoin.ApiSecretOption(ktps.venueConfig.SecretKey),
	)

	apiResponse, err := apiService.Symbols("")
	if err != nil {
		ktps.logger.Errorf("Failed to call symbols endpoint. Error: %v", err)

		return
	}

	symbolModel := kucoin.SymbolsModel{}
	if err := apiResponse.ReadData(&symbolModel); err != nil {
		ktps.logger.Errorf("Failed to call de-serailize symbols response. Error: %v", err)

		return
	}

	if err = ktps.tradingPairRepository.CreateTradingPairs(
		ktps.ctx,
		ktps.venueConfig.Id,
		mappers.KucoinSymbolModelToTradingPairs(symbolModel)); err != nil {
		ktps.logger.Errorf("Failed to create trading pairs. Error: %v", err)

		return
	}

	ktps.logger.Errorf("Finished syncing trading pairs for Venue: %s", ktps.venueConfig.Id)
}
