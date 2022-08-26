package subscribers

import (
	"context"

	"github.com/omiga-group/omiga/src/exchange/coingeko/configuration"
	"github.com/omiga-group/omiga/src/shared/enterprise/cron"
	"go.uber.org/zap"
)

type CoingekoSubscriber interface {
}

type coingekoSubscriber struct {
	ctx              context.Context
	logger           *zap.SugaredLogger
	coingekoSettings configuration.CoingekoSettings
}

func NewCoingekoSubscriber(
	ctx context.Context,
	logger *zap.SugaredLogger,
	cronService cron.CronService,
	coingekoSettings configuration.CoingekoSettings) (CoingekoSubscriber, error) {
	instance := &coingekoSubscriber{
		ctx:              ctx,
		logger:           logger,
		coingekoSettings: coingekoSettings,
	}

	if _, err := cronService.GetCron().AddJob("0/10 * * * *", instance); err != nil {
		return nil, err
	}

	return instance, nil
}

func (obs *coingekoSubscriber) Run() {

}
