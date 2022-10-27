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
	kuCoinConfig          configuration.KucoinConfig
	tradingPairRepository repositories.TradingPairRepository
}

func NewKucoinTradingPairSubscriber(
	ctx context.Context,
	logger *zap.SugaredLogger,
	kuCoinConfig configuration.KucoinConfig,
	cronService cron.CronService,
	tradingPairRepository repositories.TradingPairRepository) (KucoinTradingPairSubscriber, error) {

	instance := &kuCoinTradingPairSubscriber{
		ctx:                   ctx,
		logger:                logger,
		kuCoinConfig:          kuCoinConfig,
		tradingPairRepository: tradingPairRepository,
	}

	// Run at every minute from 0 through 59.
	if _, err := cronService.GetCron().AddJob("* 0/1 * * * *", instance); err != nil {
		return nil, err
	}

	return instance, nil
}

func (ktps *kuCoinTradingPairSubscriber) Run() {
	apiService := kucoin.NewApiService(
		kucoin.ApiKeyOption(ktps.kuCoinConfig.ApiKey),
		kucoin.ApiPassPhraseOption(ktps.kuCoinConfig.Passphrase),
		kucoin.ApiSecretOption(ktps.kuCoinConfig.SecretKey),
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
		ktps.kuCoinConfig.Id,
		mappers.KucoinSymbolModelToTradingPairs(symbolModel)); err != nil {
		ktps.logger.Errorf("Failed to create trading pairs. Error: %v", err)

		return
	}
}
