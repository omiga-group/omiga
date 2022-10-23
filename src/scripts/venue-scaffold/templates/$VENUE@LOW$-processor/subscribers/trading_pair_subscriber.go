package subscribers

import (
	"context"

	"github.com/omiga-group/omiga/src/shared/enterprise/cron"
	"github.com/omiga-group/omiga/src/venue/$VENUE@LOW$-processor/configuration"
	"github.com/omiga-group/omiga/src/venue/$VENUE@LOW$-processor/mappers"
	"github.com/omiga-group/omiga/src/venue/shared/repositories"
	"go.uber.org/zap"
)

type $VENUE@PAS$TradingPairSubscriber interface {
}

type $VENUE@LOW$TradingPairSubscriber struct {
	ctx                   context.Context
	logger                *zap.SugaredLogger
	$VENUE@LOW$Config         configuration.$VENUE@PAS$Config
	tradingPairRepository repositories.TradingPairRepository
}

func New$VENUE@PAS$TradingPairSubscriber(
	ctx context.Context,
	logger *zap.SugaredLogger,
	$VENUE@LOW$Config configuration.$VENUE@PAS$Config,
	cronService cron.CronService,
	tradingPairRepository repositories.TradingPairRepository) ($VENUE@PAS$TradingPairSubscriber, error) {

	instance := &$VENUE@LOW$TradingPairSubscriber{
		ctx:                   ctx,
		logger:                logger,
		$VENUE@LOW$Config:         $VENUE@LOW$Config,
		tradingPairRepository: tradingPairRepository,
	}

	// Run at every minute from 0 through 59.
	if _, err := cronService.GetCron().AddJob("* 0/1 * * * *", instance); err != nil {
		return nil, err
	}

	go instance.Run()

	return instance, nil
}

func (btps *$VENUE@LOW$TradingPairSubscriber) Run() {
	exchangeInfo, err := $VENUE@LOW$.
		NewClient(btps.$VENUE@LOW$Config.ApiKey, btps.$VENUE@LOW$Config.SecretKey).
		NewExchangeInfoService().
		Do(btps.ctx)
	if err != nil {
		btps.logger.Errorf("Failed to call exchangeInfo endpoint. Error: %v", err)

		return
	}

	if err = btps.tradingPairRepository.CreateTradingPairs(
		btps.ctx,
		btps.$VENUE@LOW$Config.Id,
		mappers.$VENUE@PAS$SymbolsToTradingPairs(exchangeInfo.Symbols)); err != nil {
		btps.logger.Errorf("Failed to create trading pairs. Error: %v", err)

		return
	}
}
